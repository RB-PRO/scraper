package parser

import (
	"log"
	"strings"

	"github.com/RB-PRO/PhotoTemaParser/pkg/webmg"
	"github.com/cheggaaa/pb"
)

func Parsing_webmg() {
	Categorys := []webmg.Category{
		{Name: "Рисунки", Slug: "risunki", URL: "https://webmg.ru/risunki/"},
		{Name: "Открытки", Slug: "otkrytki", URL: "https://webmg.ru/otkrytki/"},
		{Name: "Картинки и фото", Slug: "kartinki", URL: "https://webmg.ru/kartinki/"},
		{Name: "Обои", Slug: "oboi", URL: "https://webmg.ru/oboi/"},
		{Name: "Знаменитости", Slug: "znamenitosti", URL: "https://webmg.ru/kartinki/znamenitosti/"},
	}

	dr := webmg.NewDir("webmg/")
	xlsx := webmg.NewXLSX("webmg/webmg.xlsx", Categorys)

	// Цикл по категориям
	for _, Category := range Categorys {

		dr.MakeDir(Category.Slug + "/")

		// Получить ссылки на все статьи
		ArticlesLinks, ErrCategory := webmg.ParseCategorys(Category.URL)
		if ErrCategory != nil {
			log.Fatalln(Category.Name, ErrCategory)
		}

		// Цикл по всем статьям
		Bar := pb.StartNew(len(ArticlesLinks))
		Bar.Prefix(Category.Name)
		for _, ArticleLink := range ArticlesLinks {
			// if i == 3 {
			// 	break
			// }
			photos, info, ErrPhotos := webmg.ParsePhoto(ArticleLink)
			if ErrPhotos != nil {
				log.Fatalln(Category.Name, ErrPhotos)
			}

			DirectionForPhotos, _ := dr.MakeDir(Category.Slug + "/" + info.Slug + "/")
			for iPhoto := range photos {
				// https://webmg.ru/wp-content/uploads/2023/09/10097-8-jpg.webp
				FileName := EditFileName(photos[iPhoto].URL)

				photos[iPhoto].Path = DirectionForPhotos + "\\" + FileName

				dr.SavePhoto(photos[iPhoto].URL, photos[iPhoto].Path)
			}
			xlsx.WriteXLSX(Category, info, photos)

			Bar.Increment()
		}
		Bar.Finish()
	}
	xlsx.CloceAndSaveXLSX()
}

// Получить название файла по ссылке на файл
func EditFileName(str string) string {
	strs := strings.Split(str, "/")
	if len(strs) > 0 {
		return strs[len(strs)-1]
	} else {
		return str
	}
}
