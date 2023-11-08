package textiletorg

import (
	"fmt"

	"github.com/gocolly/colly"
)

const PageUrl string = "https://www.textiletorg.ru/catalog/vse-dlya-shitya/shveynye-mashiny/?PAGEN_1=%d"

func URLs_Pages(Url string) []string {

	productsURLs := make([]string, 0)

	for i := 0; i < 51; i++ {
		fmt.Println("Pages:", i, "/", 51)
		URLs := URLs_Page(fmt.Sprintf(Url, i))
		productsURLs = append(productsURLs, URLs...)
		// break
	}

	return productsURLs
}

func URLs_Page(Url string) []string {
	strs := make([]string, 0)

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML(`div[class="n_list n_catalog_name name name-catalog-category-3000"]>a`, func(e *colly.HTMLElement) {
		strs = append(strs, "https://www.textiletorg.ru"+e.Attr("href"))
	})

	c.Visit(Url)

	return strs
}
