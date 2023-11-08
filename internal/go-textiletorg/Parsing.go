package gotextiletorg

import (
	"fmt"

	"github.com/RB-PRO/PhotoTemaParser/pkg/textiletorg"
)

func Parsing() {

	urls := textiletorg.URLs_Pages(textiletorg.PageUrl)
	prods := make([]textiletorg.Product, len(urls))
	for i, url := range urls {
		fmt.Println(i, "/", len(urls))
		prod := textiletorg.Prod(url)
		prods[i] = prod

		// // фото
		// Folder := "textiletorg" + "/" + prod.SKU + "/"
		// os.MkdirAll(Folder, 0777)
		// for j := range prod.PhotoDir {
		// 	// fmt.Println(prod.PhotoURL[j], prod.PhotoDir[j])
		// 	textiletorg.SavePhoto(prod.PhotoURL[j], prod.PhotoDir[j])
		// }
	}
	textiletorg.SaveJson("textiletorg.json", prods)

	textiletorg.SaveXlsx("textiletorg.xlsx", prods)
}
