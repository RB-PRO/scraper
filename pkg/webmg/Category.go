package webmg

import (
	"fmt"

	"github.com/gocolly/colly"
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

	c.OnHTML(`div[class=post-image]>a`, func(e *colly.HTMLElement) {
		link, _ := e.DOM.Attr("href")
		links = append(links, link)
	})

	c.OnError(func(r *colly.Response, err error) {
		Err = err
	})

	c.Visit(fmt.Sprintf("%spage/%d/", URL, page))
	// page/2/
	return links, nil
}
