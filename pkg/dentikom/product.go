package dentikom

import (
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Product struct {
	Name        string
	URL         string
	SKU         string
	Category    string
	Price       float64
	Images      []string
	Deskription string
}

func ParserProduct(link string) (prod Product) {
	c := colly.NewCollector()

	// Название товара
	c.OnHTML(`div[class="detail-product-page in-basket"]>h1`, func(e *colly.HTMLElement) {
		prod.Name = e.Text
	})

	// Цена
	c.OnHTML(`div[class="dpp-price_data__price"] span[class=current-price]`, func(e *colly.HTMLElement) {
		text := e.Text
		text = strings.ReplaceAll(text, " ", "")
		text = strings.TrimSpace(text)
		prod.Price, _ = strconv.ParseFloat(text, 64)
	})

	// Картинки
	//        body > div.site > div.wrap > div.detail-product-page.in-basket > div.dpp-top > div.dpp-left > div.dpp-thumbs.slider-nav.slick-initialized.slick-slider.slick-vertical > div > div > div > img
	c.OnHTML(`a[data-fancybox=gallery]`, func(e *colly.HTMLElement) {
		//fmt.Println(e.Text)
		txt := URL + e.Attr("href")
		prod.Images = append(prod.Images, txt)
	})

	// описание
	c.OnHTML(`body > div.site > div.wrap > div.detail-product-page.in-basket > div.dpp-top > div.dpp-right > div:nth-child(2) > div.dpp-right__block_content`, func(e *colly.HTMLElement) {
		txt := e.Text
		txt = strings.TrimSpace(txt)
		prod.Deskription = txt
	})

	// Артикул
	c.OnHTML("body > div.site > div.wrap > div.detail-product-page.in-basket > div:nth-child(3) > div.art > span", func(e *colly.HTMLElement) {
		txt := e.Text
		txt = strings.TrimSpace(txt)
		txt = strings.ReplaceAll(txt, "\n\n", "\n")
		prod.SKU = txt
	})

	c.OnHTML("body > div.site > div.wrap > div.detail-product-page.in-basket > div.dpp-middle > div.dpp-middle_2 > div > div", func(e *colly.HTMLElement) {
		txt := e.Text
		txt = strings.TrimSpace(txt)
		prod.Deskription += "\n" + txt
	})

	// // Артикул
	// c.OnHTML("body", func(e *colly.HTMLElement) {
	// 	txt, _ := e.DOM.Html()
	// 	Save("body.html", txt)
	// })

	prod.URL = link
	c.Visit(link)
	return prod
}

// Простое сохранение результата в файл
func Save(FileName string, str string) error {
	return os.WriteFile(FileName, []byte(str), 0666)
}
