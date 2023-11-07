package mirkubikov

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	playwright "github.com/playwright-community/playwright-go"
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
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "BITRIX_SM_SALE_UID=1342262728; rerf=AAAAAGUe7mQ0lh9EAwhyAg==; ipp_uid=1696525923955/erqIHOCR0gOela7Z/uM4ghwN8zRrnosYOoY7JgQ==; _ga=GA1.1.2087411697.1696525926; tmr_lvid=a713f136e8c395fa5febb356d7571f8e; tmr_lvidTS=1696525925788; _ym_uid=1696525926329875913; _ym_d=1696525926; _gpVisits={\"isFirstVisitDomain\":true,\"idContainer\":\"10002548\"}; popmechanic_sbjs_migrations=popmechanic_1418474375998%3D1%7C%7C%7C1471519752600%3D1%7C%7C%7C1471519752605%3D1; adspire_uid=AS.1308471874.1696525927; ads_adware=true; BX_USER_ID=15418d1c03b9e25cce8e7a595d8510e2; flocktory-uuid=f61d44d3-9d10-4273-b477-ad0faa4edab4-7; first-visit=no; user_city=%D0%A5%D0%B8%D0%BC%D0%BA%D0%B8; _ym_isad=1; PHPSESSID=GYGNlgVTNcY21AyU9Eg9Rv1XWr53zuUO; user_usee=a%3A109%3A%7Bi%3A0%3Bs%3A7%3A%223875508%22%3Bi%3A1%3Bs%3A7%3A%223875510%22%3Bi%3A2%3Bs%3A7%3A%223875514%22%3Bi%3A3%3Bs%3A7%3A%223875516%22%3Bi%3A4%3Bs%3A7%3A%223875518%22%3Bi%3A5%3Bs%3A7%3A%223875524%22%3Bi%3A6%3Bs%3A7%3A%223875526%22%3Bi%3A7%3Bs%3A7%3A%223875536%22%3Bi%3A8%3Bs%3A7%3A%223875540%22%3Bi%3A9%3Bs%3A7%3A%223875542%22%3Bi%3A10%3Bs%3A7%3A%223875544%22%3Bi%3A11%3Bs%3A7%3A%223875548%22%3Bi%3A12%3Bs%3A7%3A%223875558%22%3Bi%3A13%3Bs%3A7%3A%223875564%22%3Bi%3A14%3Bs%3A7%3A%223875566%22%3Bi%3A15%3Bs%3A7%3A%223876068%22%3Bi%3A16%3Bs%3A7%3A%223876078%22%3Bi%3A17%3Bs%3A7%3A%223878932%22%3Bi%3A18%3Bs%3A7%3A%223879432%22%3Bi%3A19%3Bs%3A7%3A%223880186%22%3Bi%3A20%3Bs%3A7%3A%223880420%22%3Bi%3A21%3Bs%3A7%3A%223880700%22%3Bi%3A22%3Bs%3A7%3A%223880706%22%3Bi%3A23%3Bs%3A7%3A%223880708%22%3Bi%3A24%3Bs%3A7%3A%223880710%22%3Bi%3A25%3Bs%3A7%3A%223880748%22%3Bi%3A26%3Bs%3A7%3A%223880750%22%3Bi%3A27%3Bs%3A7%3A%223883006%22%3Bi%3A28%3Bs%3A7%3A%221085120%22%3Bi%3A29%3Bs%3A7%3A%221085129%22%3Bi%3A30%3Bs%3A7%3A%221372233%22%3Bi%3A31%3Bs%3A7%3A%221388120%22%3Bi%3A32%3Bs%3A7%3A%221702388%22%3Bi%3A33%3Bs%3A7%3A%221736859%22%3Bi%3A34%3Bs%3A7%3A%221941029%22%3Bi%3A35%3Bs%3A7%3A%222760644%22%3Bi%3A36%3Bs%3A7%3A%222951844%22%3Bi%3A37%3Bs%3A7%3A%223000863%22%3Bi%3A38%3Bs%3A7%3A%223000871%22%3Bi%3A39%3Bs%3A7%3A%223083000%22%3Bi%3A40%3Bs%3A7%3A%223212215%22%3Bi%3A41%3Bs%3A7%3A%223375342%22%3Bi%3A42%3Bs%3A7%3A%223423087%22%3Bi%3A43%3Bs%3A7%3A%223424453%22%3Bi%3A44%3Bs%3A7%3A%223424476%22%3Bi%3A45%3Bs%3A7%3A%223470147%22%3Bi%3A46%3Bs%3A7%3A%223551407%22%3Bi%3A47%3Bs%3A7%3A%223551418%22%3Bi%3A48%3Bs%3A7%3A%223551421%22%3Bi%3A49%3Bs%3A7%3A%223551445%22%3Bi%3A50%3Bs%3A7%3A%223551451%22%3Bi%3A51%3Bs%3A7%3A%223551458%22%3Bi%3A52%3Bs%3A7%3A%223657874%22%3Bi%3A53%3Bs%3A7%3A%223735376%22%3Bi%3A54%3Bs%3A7%3A%223735381%22%3Bi%3A55%3Bs%3A7%3A%223760767%22%3Bi%3A56%3Bs%3A7%3A%223766968%22%3Bi%3A57%3Bs%3A7%3A%223766994%22%3Bi%3A58%3Bs%3A7%3A%223877294%22%3Bi%3A59%3Bs%3A7%3A%223877310%22%3Bi%3A60%3Bs%3A7%3A%223877314%22%3Bi%3A61%3Bs%3A7%3A%223877322%22%3Bi%3A62%3Bs%3A7%3A%223877330%22%3Bi%3A63%3Bs%3A7%3A%223877332%22%3Bi%3A64%3Bs%3A7%3A%223877334%22%3Bi%3A65%3Bs%3A7%3A%223877336%22%3Bi%3A66%3Bs%3A7%3A%223880336%22%3Bi%3A67%3Bs%3A7%3A%223880410%22%3Bi%3A68%3Bs%3A7%3A%223880412%22%3Bi%3A69%3Bs%3A7%3A%223875288%22%3Bi%3A70%3Bs%3A7%3A%223875444%22%3Bi%3A71%3Bs%3A7%3A%223875472%22%3Bi%3A72%3Bs%3A7%3A%223875562%22%3Bi%3A73%3Bs%3A7%3A%223876044%22%3Bi%3A74%3Bs%3A7%3A%223876052%22%3Bi%3A75%3Bs%3A7%3A%223876064%22%3Bi%3A76%3Bs%3A7%3A%223880696%22%3Bi%3A77%3Bs%3A7%3A%223829171%22%3Bi%3A78%3Bs%3A7%3A%223880704%22%3Bi%3A79%3Bs%3A7%3A%223870938%22%3Bi%3A80%3Bs%3A7%3A%223856363%22%3Bi%3A81%3Bs%3A7%3A%223875450%22%3Bi%3A82%3Bs%3A7%3A%223875456%22%3Bi%3A83%3Bs%3A7%3A%223880740%22%3Bi%3A84%3Bs%3A7%3A%223821833%22%3Bi%3A85%3Bs%3A7%3A%223875458%22%3Bi%3A86%3Bs%3A7%3A%223871150%22%3Bi%3A87%3Bs%3A7%3A%223876072%22%3Bi%3A88%3Bs%3A7%3A%222936420%22%3Bi%3A89%3Bs%3A7%3A%223794905%22%3Bi%3A90%3Bs%3A7%3A%223795011%22%3Bi%3A91%3Bs%3A7%3A%223875552%22%3Bi%3A92%3Bs%3A7%3A%223876046%22%3Bi%3A93%3Bs%3A7%3A%223880690%22%3Bi%3A94%3Bs%3A7%3A%223880722%22%3Bi%3A95%3Bs%3A7%3A%223876034%22%3Bi%3A96%3Bs%3A7%3A%223880712%22%3Bi%3A97%3Bs%3A7%3A%223880720%22%3Bi%3A98%3Bs%3A7%3A%223871156%22%3Bi%3A99%3Bs%3A7%3A%223873734%22%3Bi%3A100%3Bs%3A7%3A%223847816%22%3Bi%3A101%3Bs%3A7%3A%223873754%22%3Bi%3A102%3Bs%3A7%3A%223875426%22%3Bi%3A103%3Bs%3A7%3A%223822977%22%3Bi%3A104%3Bs%3A7%3A%223876076%22%3Bi%3A105%3Bs%3A7%3A%223850368%22%3Bi%3A106%3Bs%3A7%3A%223212210%22%3Bi%3A107%3Bs%3A7%3A%223816152%22%3Bi%3A108%3Bs%3A7%3A%223870918%22%3B%7D; _ym_visorc=b; ipp_key=v1696629213873/v33947245ba5adc7a72e273/MbWAqdkZMi1v/IJ6MTf8Pw==; _gp10002548={\"hits\":98,\"vc\":1,\"ac\":1,\"a6\":1}; tmr_detect=1%7C1696629251712; _ga_61W59B3053=GS1.1.1696623419.6.1.1696629252.0.0.0; mindboxDeviceUUID=277600e1-2469-4aea-8214-06dfbed47d1f; directCrm-session=%7B%22deviceGuid%22%3A%22277600e1-2469-4aea-8214-06dfbed47d1f%22%7D; BITRIX_SM_SALE_UID=1342741216; PHPSESSID=tdnVVWoT0P2kYZGR1yOt9H2eiAaToaSl; ipp_key=v1696629042223/v33947245b85ad87a72e273/K4QONnANpmUAutS/y4cSTA==; ipp_uid=1696629042223/HUGpho0bEt0CxGYM/PoxcB61ZnNSzWThwd0l0UA==; rerf=AAAAAGUggTMG0UmMAwTbAg==; user_usee=a%3A133%3A%7Bi%3A0%3Bs%3A7%3A%223858633%22%3Bi%3A1%3Bs%3A7%3A%223858634%22%3Bi%3A2%3Bs%3A7%3A%223858635%22%3Bi%3A3%3Bs%3A7%3A%223858636%22%3Bi%3A4%3Bs%3A7%3A%223858638%22%3Bi%3A5%3Bs%3A7%3A%223858640%22%3Bi%3A6%3Bs%3A7%3A%223858641%22%3Bi%3A7%3Bs%3A7%3A%223858642%22%3Bi%3A8%3Bs%3A7%3A%223859230%22%3Bi%3A9%3Bs%3A7%3A%223859231%22%3Bi%3A10%3Bs%3A7%3A%223859232%22%3Bi%3A11%3Bs%3A7%3A%223859233%22%3Bi%3A12%3Bs%3A7%3A%223859234%22%3Bi%3A13%3Bs%3A7%3A%223871096%22%3Bi%3A14%3Bs%3A7%3A%223871132%22%3Bi%3A15%3Bs%3A7%3A%223871134%22%3Bi%3A16%3Bs%3A7%3A%223871136%22%3Bi%3A17%3Bs%3A7%3A%223871138%22%3Bi%3A18%3Bs%3A7%3A%223871140%22%3Bi%3A19%3Bs%3A7%3A%223871142%22%3Bi%3A20%3Bs%3A7%3A%223871144%22%3Bi%3A21%3Bs%3A7%3A%223871146%22%3Bi%3A22%3Bs%3A7%3A%223871148%22%3Bi%3A23%3Bs%3A7%3A%223871152%22%3Bi%3A24%3Bs%3A7%3A%223871154%22%3Bi%3A25%3Bs%3A7%3A%223871158%22%3Bi%3A26%3Bs%3A7%3A%223871160%22%3Bi%3A27%3Bs%3A7%3A%223871162%22%3Bi%3A28%3Bs%3A7%3A%223871164%22%3Bi%3A29%3Bs%3A7%3A%223871166%22%3Bi%3A30%3Bs%3A7%3A%223871168%22%3Bi%3A31%3Bs%3A7%3A%223871170%22%3Bi%3A32%3Bs%3A7%3A%223871172%22%3Bi%3A33%3Bs%3A7%3A%223871174%22%3Bi%3A34%3Bs%3A7%3A%223871176%22%3Bi%3A35%3Bs%3A7%3A%223871178%22%3Bi%3A36%3Bs%3A7%3A%223871180%22%3Bi%3A37%3Bs%3A7%3A%223871182%22%3Bi%3A38%3Bs%3A7%3A%223871184%22%3Bi%3A39%3Bs%3A7%3A%223871186%22%3Bi%3A40%3Bs%3A7%3A%223871188%22%3Bi%3A41%3Bs%3A7%3A%223871190%22%3Bi%3A42%3Bs%3A7%3A%223871192%22%3Bi%3A43%3Bs%3A7%3A%223871194%22%3Bi%3A44%3Bs%3A7%3A%223871196%22%3Bi%3A45%3Bs%3A7%3A%223871198%22%3Bi%3A46%3Bs%3A7%3A%223871200%22%3Bi%3A47%3Bs%3A7%3A%223871202%22%3Bi%3A48%3Bs%3A7%3A%223871204%22%3Bi%3A49%3Bs%3A7%3A%223871206%22%3Bi%3A50%3Bs%3A7%3A%223871280%22%3Bi%3A51%3Bs%3A7%3A%223871282%22%3Bi%3A52%3Bs%3A7%3A%223871284%22%3Bi%3A53%3Bs%3A7%3A%223871286%22%3Bi%3A54%3Bs%3A7%3A%223871886%22%3Bi%3A55%3Bs%3A7%3A%223872392%22%3Bi%3A56%3Bs%3A7%3A%223872394%22%3Bi%3A57%3Bs%3A7%3A%223872872%22%3Bi%3A58%3Bs%3A7%3A%223872880%22%3Bi%3A59%3Bs%3A7%3A%223872890%22%3Bi%3A60%3Bs%3A7%3A%223873322%22%3Bi%3A61%3Bs%3A7%3A%223873472%22%3Bi%3A62%3Bs%3A7%3A%223873758%22%3Bi%3A63%3Bs%3A7%3A%223875280%22%3Bi%3A64%3Bs%3A7%3A%223875286%22%3Bi%3A65%3Bs%3A7%3A%223875290%22%3Bi%3A66%3Bs%3A7%3A%223875296%22%3Bi%3A67%3Bs%3A7%3A%223875298%22%3Bi%3A68%3Bs%3A7%3A%223875302%22%3Bi%3A69%3Bs%3A7%3A%223875312%22%3Bi%3A70%3Bs%3A7%3A%223875314%22%3Bi%3A71%3Bs%3A7%3A%223875316%22%3Bi%3A72%3Bs%3A7%3A%223875318%22%3Bi%3A73%3Bs%3A7%3A%223875460%22%3Bi%3A74%3Bs%3A7%3A%223875464%22%3Bi%3A75%3Bs%3A7%3A%223875468%22%3Bi%3A76%3Bs%3A7%3A%223875474%22%3Bi%3A77%3Bs%3A7%3A%223875478%22%3Bi%3A78%3Bs%3A7%3A%223875480%22%3Bi%3A79%3Bs%3A7%3A%223875490%22%3Bi%3A80%3Bs%3A7%3A%223875496%22%3Bi%3A81%3Bs%3A7%3A%223875498%22%3Bi%3A82%3Bs%3A7%3A%223875508%22%3Bi%3A83%3Bs%3A7%3A%223875510%22%3Bi%3A84%3Bs%3A7%3A%223875514%22%3Bi%3A85%3Bs%3A7%3A%223875516%22%3Bi%3A86%3Bs%3A7%3A%223875518%22%3Bi%3A87%3Bs%3A7%3A%223875524%22%3Bi%3A88%3Bs%3A7%3A%223875526%22%3Bi%3A89%3Bs%3A7%3A%223875536%22%3Bi%3A90%3Bs%3A7%3A%223875540%22%3Bi%3A91%3Bs%3A7%3A%223875542%22%3Bi%3A92%3Bs%3A7%3A%223875544%22%3Bi%3A93%3Bs%3A7%3A%223875548%22%3Bi%3A94%3Bs%3A7%3A%223875558%22%3Bi%3A95%3Bs%3A7%3A%223875564%22%3Bi%3A96%3Bs%3A7%3A%223875566%22%3Bi%3A97%3Bs%3A7%3A%223876068%22%3Bi%3A98%3Bs%3A7%3A%223876078%22%3Bi%3A99%3Bs%3A7%3A%223878932%22%3Bi%3A100%3Bs%3A7%3A%223879432%22%3Bi%3A101%3Bs%3A7%3A%223880186%22%3Bi%3A102%3Bs%3A7%3A%223880420%22%3Bi%3A103%3Bs%3A7%3A%223880700%22%3Bi%3A104%3Bs%3A7%3A%223880706%22%3Bi%3A105%3Bs%3A7%3A%223880708%22%3Bi%3A106%3Bs%3A7%3A%223880710%22%3Bi%3A107%3Bs%3A7%3A%223880748%22%3Bi%3A108%3Bs%3A7%3A%223880750%22%3Bi%3A109%3Bs%3A7%3A%223883006%22%3Bi%3A110%3Bs%3A7%3A%221085120%22%3Bi%3A111%3Bs%3A7%3A%223871150%22%3Bi%3A112%3Bs%3A7%3A%223876072%22%3Bi%3A113%3Bs%3A7%3A%223794905%22%3Bi%3A114%3Bs%3A7%3A%223795011%22%3Bi%3A115%3Bs%3A7%3A%223875552%22%3Bi%3A116%3Bs%3A7%3A%223876046%22%3Bi%3A117%3Bs%3A7%3A%223880690%22%3Bi%3A118%3Bs%3A7%3A%223880722%22%3Bi%3A119%3Bs%3A7%3A%223876034%22%3Bi%3A120%3Bs%3A7%3A%223880712%22%3Bi%3A121%3Bs%3A7%3A%223880720%22%3Bi%3A122%3Bs%3A7%3A%223871156%22%3Bi%3A123%3Bs%3A7%3A%223873734%22%3Bi%3A124%3Bs%3A7%3A%223847816%22%3Bi%3A125%3Bs%3A7%3A%223873754%22%3Bi%3A126%3Bs%3A7%3A%223875426%22%3Bi%3A127%3Bs%3A7%3A%223822977%22%3Bi%3A128%3Bs%3A7%3A%223876076%22%3Bi%3A129%3Bs%3A7%3A%223850368%22%3Bi%3A130%3Bs%3A7%3A%223212210%22%3Bi%3A131%3Bs%3A7%3A%223816152%22%3Bi%3A132%3Bs%3A7%3A%223870918%22%3B%7D")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
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

type Core struct {
	pw      *playwright.Playwright
	browser playwright.Browser
	pagea   playwright.Page
}

func NewCore() (*Core, error) {
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
	return &Core{
		pw:      pw,
		browser: browser,
		pagea:   pagea,
	}, nil
}
func (c *Core) DeleteCore() (Err error) {
	if Err = c.browser.Close(); Err != nil {
		return Err
	}
	if Err = c.pw.Stop(); Err != nil {
		return Err
	}
	return nil
}
func (c *Core) Parsing(url string) (prod Product, Err error) {

	if _, Err := c.pagea.Goto(url, playwright.PageGotoOptions{WaitUntil: (*playwright.WaitUntilState)(playwright.String("commit"))}); Err != nil {
		log.Fatalf("could not goto: %v", Err)
	}

	c.pagea.WaitForSelector("span[class=product__header__name]", playwright.PageWaitForSelectorOptions{State: playwright.WaitForSelectorStateVisible})

	// Артикул
	ElementHandleSKU, ErrElementHandleSKU := c.pagea.QuerySelector("div[class=product__header]>div[class=security]")
	if ErrElementHandleSKU == nil {
		sku, _ := ElementHandleSKU.GetAttribute("data-id")
		prod.SKU = sku
	}
	// Название
	ElementHandleName, ErrElementHandleName := c.pagea.QuerySelector("span[class=product__header__name]")
	if ErrElementHandleName == nil {
		Name, _ := ElementHandleName.TextContent()
		Name = strings.TrimSpace(Name)
		prod.Name = Name
	}
	// Описание товара
	ElementHandleDesc, ErrElementHandleDesc := c.pagea.QuerySelector(`div[class="product__additional__text g-content g-relative js-product-text"]`)
	if ErrElementHandleDesc == nil {
		if ElementHandleDesc != nil {
			Description, err := ElementHandleDesc.TextContent()
			if err == nil {
				Description = strings.ReplaceAll(Description, "Читать далее", "")
				Description = strings.ReplaceAll(Description, "...", "")
				Description = strings.TrimSpace(Description)
				prod.Description = Description
			}
		}
	}
	// Цена
	ElementHandlePrice, ErrElementHandlePrice := c.pagea.QuerySelector("span[class=js-datalayer-data]")
	if ErrElementHandlePrice == nil {
		if ElementHandlePrice != nil {
			PriceStr, err := ElementHandlePrice.GetAttribute("data-price")
			if err == nil {
				if Price, err := strconv.ParseFloat(PriceStr, 64); err == nil {
					prod.Price = Price
				}
			}
		}
	}

	// Фото
	handles, _ := c.pagea.QuerySelectorAll(`div[class=fastview__main-container]>div>div[class="slick-list draggable"]>div[class=slick-track]>div[aria-hidden=true]`)
	for _, handle := range handles {
		LinkSrc, err := handle.GetAttribute("data-hires-src")
		if err == nil {
			if LinkSrc != "" {
				prod.PhotoLinks = append(prod.PhotoLinks, LinkSrc)
			}
		}
	}
	prod.PhotoLinks = RemoveDuplicateStr(prod.PhotoLinks)

	// Ссылка на товар
	prod.URL = fmt.Sprintf("https://mir-kubikov.ru/catalog/%s/", prod.SKU)

	// for page := 1; ; {
	// 	var LinksPages []string
	// 	pagea.Goto(fmt.Sprintf("%spage/%d/", URL, page))
	// 	time.Sleep(3 * time.Second)

	// 	handles, _ := pagea.QuerySelectorAll(`article>div>div>a`)
	// 	fmt.Println("handles", len(handles))
	// 	for _, handle := range handles {
	// 		link, _ := handle.GetAttribute("href")
	// 		LinksPages = append(LinksPages, link)
	// 	}

	// 	fmt.Println("len(LinksPages)", len(LinksPages))
	// 	if len(LinksPages) == 0 {
	// 		break
	// 	}
	// 	links = append(links, LinksPages...)
	// 	page++
	// }

	return prod, nil
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
