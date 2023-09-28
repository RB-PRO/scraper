package webmg

import (
	"strings"

	"github.com/gocolly/colly"
)

// Структура базовой единицы парсинга - фотографии
type Photo struct {
	URL  string // Ссылка на картинку в источнике
	Path string // Путь к картинке в папке
}

// Информация по статье
type Info struct {
	Title       string // Заголовок
	Slug        string // Статья slug на русском, откуда взята картинка
	Description string // Описание
	URL         string // Ссылка на категорию
}

// Спарсить фото со страницы
func ParsePhoto(URL string) (photos []Photo, info Info, Err error) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.2271 YaBrowser/23.9.0.2271 Yowser/2.5 Safari/537.36"

	// Информация о разделе
	c.OnHTML(`div[class^=inside-page-hero]>h1`, func(e *colly.HTMLElement) {
		info.Title = e.DOM.Text()
	})
	c.OnHTML("div[class=entry-content]>p:first-of-type", func(e *colly.HTMLElement) {
		info.Description = e.DOM.Text()
	})
	c.OnHTML(`link[rel=canonical]`, func(e *colly.HTMLElement) {
		info.URL, _ = e.DOM.Attr("href")
		info.Slug = strings.ReplaceAll(info.URL, "https://webmg.ru/", "")
		info.Slug = strings.ReplaceAll(info.Slug, "/", "")
	})
	info.Title = EditInfo(info.Title)
	info.Slug = EditInfo(info.Slug)
	info.Description = EditInfo(info.Description)
	info.URL = EditInfo(info.URL)

	// Картинки :nth-of-type(2)
	c.OnHTML("div[class=entry-content]>p>a>img", func(e *colly.HTMLElement) {
		url, _ := e.DOM.Attr("data-lazy-src")
		if url != "" {
			photos = append(photos, Photo{
				URL: url,
			})
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		Err = err
	})

	c.Visit(URL)
	return photos, info, nil
}

// Редактировать текста по этой статье
func EditInfo(str string) string {
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.TrimSpace(str)
	return str
}
