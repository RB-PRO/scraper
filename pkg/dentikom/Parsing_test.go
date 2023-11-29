package dentikom

import (
	"fmt"
	"testing"
)

func TestProd(t *testing.T) {
	cat := Categs()
	fmt.Println(len(cat))
	fmt.Println(cat[2])
	pagelink, _ := page(cat[2].Link, 1)
	// fmt.Println(len(Pages(cat[2].Link)))
	fmt.Println(pagelink[0])
	fmt.Printf("\n\n%+v\n\n", ParserProduct(pagelink[0]))
}
