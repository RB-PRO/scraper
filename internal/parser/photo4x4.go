package parser

import (
	"fmt"
	"log"

	"github.com/RB-PRO/PhotoTemaParser/pkg/photo4x4"
)

func Parsing_4x4photo() {
	Categorys := []photo4x4.Category{
		{Name: "Фото-картинки", Slug: "foto-kartinki", URL: "https://4x4photo.ru/foto-kartinki/"},
		{Name: "Открытки", Slug: "otkrytki", URL: "https://4x4photo.ru/otkrytki/"},
		{Name: "Рисунки", Slug: "risunki", URL: "https://4x4photo.ru/risunki/"},
	}

	// dr := photo4x4.NewDir("photo4x4/")
	// xlsx := photo4x4.NewXLSX("photo4x4/photo4x4.xlsx", Categorys)

	// Цикл по категориям
	for _, Category := range Categorys {

		// Получить ссылки на все статьи
		ArticlesLinks, ErrCategory := photo4x4.ParseCategorys(Category.URL)
		if ErrCategory != nil {
			log.Fatalln(Category.Name, ErrCategory)
		}

		fmt.Println(Category.Name, len(ArticlesLinks))

	}

}
