package padentikom

import (
	"fmt"
	"slices"
	"time"

	"github.com/RB-PRO/PhotoTemaParser/pkg/dentikom"
)

func Parsing() {
	cats := dentikom.Categs()
	var prods []dentikom.Product
	// for _, cat := range cats {
	fmt.Println(len(cats))
	urls := []string{}
	for i := len(cats) - 1; i >= 0; i-- {
		cat := cats[i]
		fmt.Println(cat.Name)
		links := dentikom.Pages(cat.Link)
		for _, link := range links {
			if !slices.Contains(urls, link) {
				prod := dentikom.ParserProduct(link)
				prod.Category = cat.Name
				urls = append(urls, link)
				prods = append(prods, prod)
				time.Sleep(100 * time.Millisecond)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	dentikom.SaveXlsx("dentikom.xlsx", prods)
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
