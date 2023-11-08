package textiletorg

import (
	"fmt"
	"testing"
)

func TestPages(t *testing.T) {
	urls := URLs_Pages(PageUrl)
	fmt.Println(len(urls))
	fmt.Println(urls[0])
}
func TestProd(t *testing.T) {
	Prod := Prod("https://www.textiletorg.ru/catalog/vse-dlya-shitya/shveynye-mashiny/shveynaya-mashina-comfort-14.html")
	fmt.Printf("%+v\n", Prod)
	fmt.Println(len(Prod.PhotoURL))
	fmt.Println(len(Prod.PhotoDir))
	SaveJson("textiletorg_test.json", []Product{Prod})
	SaveXlsx("textiletorg_test.xlsx", []Product{Prod})
}
