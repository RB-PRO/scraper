package dentikom

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const URL string = "https://dentikom.ru"

type Category struct {
	Link string
	Name string
}

func Categs() (cat []Category) {
	c := colly.NewCollector()
	c.OnHTML("div[class=index-catalog] ul>li>a", func(e *colly.HTMLElement) {
		href := URL + e.Attr("href")
		text := e.Text
		text = strings.ReplaceAll(text, "+", "")
		text = strings.TrimSpace(text)
		cat = append(cat, Category{Link: href, Name: text})
	})
	c.Visit(URL)
	return cat
}

func Pages(link string) (links []string) {
	next := true
	for pg := 1; next; pg++ {
		var lim []string
		lim, next = page(link, pg)
		fmt.Println(pg, len(lim))
		links = append(links, lim...)
	}
	return links
}

func page(link string, page int) (links []string, nexts bool) {
	c := colly.NewCollector()
	c.OnHTML("div[id=catalog-products]>div[class=flex]>div[class=item]>a[class=name]", func(e *colly.HTMLElement) {
		href := URL + e.Attr("href")
		links = append(links, href)
	})
	c.OnHTML("a[class=pnext]", func(e *colly.HTMLElement) {
		nexts = true
	})
	c.Visit(link + "?PAGEN_2=" + strconv.Itoa(page))
	return links, nexts
}
