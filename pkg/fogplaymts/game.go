package fogplay

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const URL string = "https://fogplay.mts.ru/"

type Game struct {
	Name  string
	Coats []int
}

func Servers() (Games []Game) {
	c := colly.NewCollector()
	var GameServer Game

	c.OnHTML(`div[class=card__price]`, func(e *colly.HTMLElement) {
		text := e.DOM.Text()
		text = strings.ReplaceAll(text, "₽/час", "")
		text = strings.TrimSpace(text)
		price, _ := strconv.Atoi(text)
		GameServer.Coats = append(GameServer.Coats, price)
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
