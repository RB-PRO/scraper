package mirkubikovapp

import (
	"fmt"
	"io/ioutil"
	"time"

	mirkubikov "github.com/RB-PRO/PhotoTemaParser/pkg/mir-kubikov"
)

func Categorys() {
	URL := "https://mir-kubikov.ru/catalog/constructors/lego/?page=%d&show_more=Y"

	xlsx := mirkubikov.NewXLSX("mir-kubikov.xlsx")
	defer xlsx.CloceAndSaveXLSX()

	var datas string

	for page := 1; ; page++ {
		url := fmt.Sprintf(URL, page)
		ProductsLine, NextPageIsExit, ErrParseProductsUrl := mirkubikov.ParseProductsUrlWithCategory(url)
		if ErrParseProductsUrl != nil {
			panic(ErrParseProductsUrl)
		}

		for _, prod := range ProductsLine {
			datas += fmt.Sprintf("%s;%s\n", prod.Name, prod.Category)
		}

		// Если дальше страниц не будет
		if !NextPageIsExit {
			break
		}
		// break
		time.Sleep(time.Second)
	}

	ioutil.WriteFile("testfile.txt", []byte(datas), 0644)
}
