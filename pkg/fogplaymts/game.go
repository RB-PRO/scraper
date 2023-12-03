package fogplay

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const URL string = "https://fogplay.mts.ru"

type Game struct {
	Name  string
	Coats []int
}

func Servers() (Games []Game) {
	c := colly.NewCollector()
	var GameServer Game

	// GG := make(map[string][]int)

	c.OnHTML(`div[class=card__price]`, func(e *colly.HTMLElement) {
		text := e.DOM.Text()
		text = strings.ReplaceAll(text, "₽/час", "")
		text = strings.TrimSpace(text)
		price, _ := strconv.Atoi(text)
		GameServer.Coats = append(GameServer.Coats, price)
	})
	c.OnHTML(`a[class=pagination__next]`, func(f *colly.HTMLElement) {
		if linkF := f.Attr("href"); linkF != "" && strings.Contains(linkF, "game") {
			// fmt.Println(linkF)
			time.Sleep(time.Millisecond * 200)
			f.Request.Visit(linkF)
		}
	})
	c.OnHTML(`div[class="card card--version-3"]`, func(e *colly.HTMLElement) {
		GameServer = Game{}
		GameServer.Name = e.DOM.Find("h2[class=card__title]").Text()
		link, _ := e.DOM.Find(`a[class="stretched-link"]`).Attr("href")
		fmt.Println(GameServer.Name)

		e.Request.Visit(link)
		Games = append(Games, GameServer)

	})

	c.Visit(URL)
	return Games
}

type Server struct {
	GameName string
	URL      string
}

func Serverss() (Servers []Server) {
	c := colly.NewCollector()

	c.OnHTML(`div[class="card card--version-3"]`, func(e *colly.HTMLElement) {
		s := Server{}
		s.GameName = e.DOM.Find("h2[class=card__title]").Text()
		s.URL, _ = e.DOM.Find(`a[class="stretched-link"]`).Attr("href")
		// fmt.Println(GameServer.Name)

		Servers = append(Servers, s)
		// Games = append(Games, GameServer)

	})

	c.OnHTML(`a[class=pagination__next]`, func(f *colly.HTMLElement) {
		if linkF := f.Attr("href"); linkF != "" {
			fmt.Println(linkF)
			time.Sleep(time.Millisecond * 100)
			f.Request.Visit(linkF)
		}
	})

	c.Visit(URL)
	return Servers
}
func Gamess(Servers Server) (Games Game) {
	Games.Name = Servers.GameName
	c := colly.NewCollector()

	c.OnHTML(`div[class=card__price]`, func(e *colly.HTMLElement) {
		text := e.DOM.Text()
		text = strings.ReplaceAll(text, "₽/час", "")
		text = strings.TrimSpace(text)
		price, _ := strconv.Atoi(text)
		Games.Coats = append(Games.Coats, price)
	})
	c.OnHTML(`a[class=pagination__next]`, func(f *colly.HTMLElement) {
		if linkF := f.Attr("href"); linkF != "" && strings.Contains(linkF, "game") {
			fmt.Println(Games.Name, linkF)
			time.Sleep(time.Millisecond * 150)
			f.Request.Visit(linkF)
		}
	})

	fmt.Println(URL + Servers.URL)
	c.Visit(URL + Servers.URL)
	return Games
}

// #fog-banner > div > div.fog-banner__content > h2 > span:nth-child(1)
// #servers_grid > div.card.card-outside.computer-4968
