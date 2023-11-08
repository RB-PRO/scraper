package textiletorg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Product struct {
	Name     string
	SKU      string
	URL      string
	Price    float64
	PhotoURL []string
	PhotoDir []string

	// Technical specifications
	TS map[string]string

	Deskription string
}

func Prod(url string) (prod Product) {
	prod.TS = make(map[string]string)
	prod.URL = url
	c := colly.NewCollector()

	c.OnHTML(`div[class="detail-page"]>h1`, func(e *colly.HTMLElement) {
		Name := e.DOM.Text()
		prod.Name = Name
	})
	c.OnHTML(`div[class="item-outer-block"]>span[class="article"]`, func(e *colly.HTMLElement) {
		SKU := e.DOM.Text()
		SKU = strings.ReplaceAll(SKU, "Артикул", "")
		SKU = strings.TrimSpace(SKU)
		prod.SKU = SKU
	})
	c.OnHTML(`span[class^="price-price"]`, func(e *colly.HTMLElement) {
		PriceStr := e.DOM.Text()
		PriceStr = strings.ReplaceAll(PriceStr, " ", "")
		Price, _ := strconv.ParseFloat(PriceStr, 64)
		prod.Price = Price
	})
	c.OnHTML(`div[id="multiple-items"] img`, func(e *colly.HTMLElement) {
		src, _ := e.DOM.Attr("src")
		if src != "" {
			prod.PhotoURL = append(prod.PhotoURL, "https://www.textiletorg.ru"+src)
			// prod.PhotoDir = append(prod.PhotoDir, fmt.Sprintf("textiletorg/%s/%d.jpg", prod.SKU, len(prod.PhotoURL)))
		}
	})
	c.OnHTML(`div[class="descriptionp"]>p`, func(e *colly.HTMLElement) {
		Desk := e.DOM.Text()
		Desk = strings.TrimSpace(Desk)
		prod.Deskription += Desk + "\n"
	})

	c.OnHTML(`div[id="item_har"]>ul[class="eshop-item-param"]>li`, func(e *colly.HTMLElement) {
		key := e.DOM.Find(`p[class="eshop-item-param-name"]`)
		val := e.DOM.Find(`p[class="eshop-item-param-value"]`)
		keyStr := key.Text()
		valStr := val.Text()
		keyStr = strings.TrimSpace(keyStr)
		valStr = strings.TrimSpace(valStr)
		prod.TS[keyStr] = valStr
	})

	// доп описание
	c.OnHTML(`div[class="descriptionp"] tr`, func(e *colly.HTMLElement) {
		key := e.DOM.Find(`td:first-of-type`)
		val := e.DOM.Find(`td:last-of-type`)
		keyStr := key.Text()
		valStr := val.Text()
		keyStr = strings.TrimSpace(keyStr)
		valStr = strings.TrimSpace(valStr)
		keyStr = strings.ReplaceAll(keyStr, "\n", "")
		valStr = strings.ReplaceAll(valStr, "\n", "")
		// fmt.Println(keyStr, valStr)
		if keyStr != "" && valStr != "" {
			prod.TS[keyStr] = valStr
		}
	})

	c.Visit(url)

	prod.PhotoURL = RemoveDuplicateStr(prod.PhotoURL)
	for i := range prod.PhotoURL {
		prod.PhotoDir = append(prod.PhotoDir, fmt.Sprintf("textiletorg/%s/%d.jpg", prod.SKU, i+1))
	}
	return prod
}

// Удалить дубликаты в слайсе
func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
