package photo4x4

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
	"github.com/playwright-community/playwright-go"
)

type Category struct {
	Name string // Название категории
	Slug string // Ярлык категории
	URL  string // Ссылка на категорию
}

// Спарсить ссылки на все статьи на одной странице
func ParseCategorys(URL string) (links []string, Err error) {

	// Цикл по всем страницам
	for page := 1; ; {
		LinksPages, Err := ParseCategory(URL, page)
		if Err != nil {
			return nil, Err
		}
		if len(LinksPages) == 0 {
			break
		}
		links = append(links, LinksPages...)
		page++
	}

	return links, nil
}

// Спарсить ссылки на все статьи на одной странице
func ParseCategory(URL string, page int) (links []string, Err error) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.2271 YaBrowser/23.9.0.2271 Yowser/2.5 Safari/537.36"

	c.OnHTML(`article>div>div>a`, func(e *colly.HTMLElement) {
		link, _ := e.DOM.Attr("href")
		links = append(links, link)
	})

	c.OnError(func(r *colly.Response, err error) {
		Err = err
	})

	fmt.Printf("%spage/%d/", URL, page)
	c.Visit(fmt.Sprintf("%spage/%d/", URL, page))

	//////////////////////////////////////////////////////
	//////////////////////////////////////////////////////

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	pagea, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = pagea.Goto(fmt.Sprintf("%spage/%d/", URL, page)); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	pagea.WaitForSelector("div[id=primary-menu]")

	for page := 1; ; {
		var LinksPages []string
		pagea.Goto(fmt.Sprintf("%spage/%d/", URL, page))
		time.Sleep(3 * time.Second)

		handles, _ := pagea.QuerySelectorAll(`article>div>div>a`)
		fmt.Println("handles", len(handles))
		for _, handle := range handles {
			link, _ := handle.GetAttribute("href")
			LinksPages = append(LinksPages, link)
		}

		fmt.Println("len(LinksPages)", len(LinksPages))
		if len(LinksPages) == 0 {
			break
		}
		links = append(links, LinksPages...)
		page++
	}

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
	return links, nil
}
