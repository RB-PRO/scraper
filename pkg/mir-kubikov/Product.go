package mirkubikov

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	Name        string   // Название товара
	SKU         string   // Артикул
	Price       float64  // Цена
	Description string   // Описание товара
	PhotoLinks  []string // Ссылки на фото источника
	PhotoPaths  []string // Ссылки на локальные файлы
	URL         string   // Ссылка на товар
}

// Спарсить страницу товара
func ParseProduct(url string) (prod Product, Err error) {
	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return Product{}, ErrNewRequest
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "BITRIX_SM_SALE_UID=1342262728; rerf=AAAAAGUe7mQ0lh9EAwhyAg==; ipp_uid=1696525923955/erqIHOCR0gOela7Z/uM4ghwN8zRrnosYOoY7JgQ==; _ga=GA1.1.2087411697.1696525926; tmr_lvid=a713f136e8c395fa5febb356d7571f8e; tmr_lvidTS=1696525925788; _ym_uid=1696525926329875913; _ym_d=1696525926; _gpVisits={\"isFirstVisitDomain\":true,\"idContainer\":\"10002548\"}; popmechanic_sbjs_migrations=popmechanic_1418474375998%3D1%7C%7C%7C1471519752600%3D1%7C%7C%7C1471519752605%3D1; adspire_uid=AS.1308471874.1696525927; ads_adware=true; _ym_isad=1; BX_USER_ID=15418d1c03b9e25cce8e7a595d8510e2; flocktory-uuid=f61d44d3-9d10-4273-b477-ad0faa4edab4-7; first-visit=no; user_city=%D0%A5%D0%B8%D0%BC%D0%BA%D0%B8; PHPSESSID=V5kg5dd1GevSQmkcXcdvpfzYDkT6sL5R; _ym_visorc=w; ipp_key=v1696594270546/v33947245ba5adc7a72e273/4Fs1X6oM3GtTTi+DhB08NA==; _gp10002548={\"hits\":70,\"vc\":1,\"ac\":1,\"a6\":1}; tmr_detect=1%7C1696594386759; mindboxDeviceUUID=277600e1-2469-4aea-8214-06dfbed47d1f; directCrm-session=%7B%22deviceGuid%22%3A%22277600e1-2469-4aea-8214-06dfbed47d1f%22%7D; user_usee=a%3A19%3A%7Bi%3A0%3Bs%3A7%3A%223795011%22%3Bi%3A1%3Bs%3A7%3A%223875552%22%3Bi%3A2%3Bs%3A7%3A%223876046%22%3Bi%3A3%3Bs%3A7%3A%223880690%22%3Bi%3A4%3Bs%3A7%3A%223880722%22%3Bi%3A5%3Bs%3A7%3A%223876034%22%3Bi%3A6%3Bs%3A7%3A%223880712%22%3Bi%3A7%3Bs%3A7%3A%223880720%22%3Bi%3A8%3Bs%3A7%3A%223871156%22%3Bi%3A9%3Bs%3A7%3A%223873734%22%3Bi%3A10%3Bs%3A7%3A%223847816%22%3Bi%3A11%3Bs%3A7%3A%223873754%22%3Bi%3A12%3Bs%3A7%3A%223875426%22%3Bi%3A13%3Bs%3A7%3A%223822977%22%3Bi%3A14%3Bs%3A7%3A%223876076%22%3Bi%3A15%3Bs%3A7%3A%223850368%22%3Bi%3A16%3Bs%3A7%3A%223212210%22%3Bi%3A17%3Bs%3A7%3A%223816152%22%3Bi%3A18%3Bs%3A7%3A%223870918%22%3B%7D; _ga_61W59B3053=GS1.1.1696594270.3.1.1696594621.0.0.0; ipp_key=v1696582862834/v33947245b85ad87a72e273/hioacGMW2h0wY4Eo6HvCnQ==; user_usee=a%3A6%3A%7Bi%3A0%3Bs%3A7%3A%223822977%22%3Bi%3A1%3Bs%3A7%3A%223876076%22%3Bi%3A2%3Bs%3A7%3A%223850368%22%3Bi%3A3%3Bs%3A7%3A%223212210%22%3Bi%3A4%3Bs%3A7%3A%223816152%22%3Bi%3A5%3Bs%3A7%3A%223870918%22%3B%7D")
	req.Header.Add("Referer", "https://mir-kubikov.ru/catalog/?page=176")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.967 YaBrowser/23.9.1.967 Yowser/2.5 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return Product{}, ErrDo
	}
	defer res.Body.Close()

	// parse the HTML document
	doc, ErrNewDocumentFromReader := goquery.NewDocumentFromReader(res.Body)
	if ErrNewDocumentFromReader != nil {
		return Product{}, ErrNewDocumentFromReader
	}

	// Вычленяем инфомрацию
	// Артикул
	doc.Find("div[class=product__header]>div[class=security]").Each(func(i int, e *goquery.Selection) {
		sku, _ := e.Attr("data-id")
		prod.SKU = sku
	})
	// Название
	doc.Find("span[class=product__header__name]").Each(func(i int, e *goquery.Selection) {
		Name := e.Text()
		Name = strings.TrimSpace(Name)
		prod.Name = Name
	})
	// Описание товара
	doc.Find(`div[class="product__additional__text g-content g-relative js-product-text"]`).Each(func(i int, e *goquery.Selection) {
		Description := e.Text()
		Description = strings.ReplaceAll(Description, "Читать далее", "")
		Description = strings.ReplaceAll(Description, "...", "")
		Description = strings.TrimSpace(Description)
		prod.Description = Description
	})
	// Цена
	doc.Find("span[class=js-datalayer-data]").Each(func(i int, e *goquery.Selection) {
		PriceStr, _ := e.Attr("data-price")
		if Price, err := strconv.ParseFloat(PriceStr, 64); err == nil {
			prod.Price = Price
		}
	})
	// Фотографии
	doc.Find(`div[class=fastview__images-slider-item]`).Each(func(i int, e *goquery.Selection) {
		LinkSrc, _ := e.Attr("data-hires-src")
		prod.PhotoLinks = append(prod.PhotoLinks, LinkSrc)
	})
	// Ссылка на товар
	prod.URL = fmt.Sprintf("https://mir-kubikov.ru/catalog/%s/", prod.SKU)

	return prod, nil
}
