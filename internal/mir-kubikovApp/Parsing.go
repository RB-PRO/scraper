package mirkubikovapp

import (
	"fmt"
	"strings"
	"time"

	mirkubikov "github.com/RB-PRO/PhotoTemaParser/pkg/mir-kubikov"
	"github.com/cheggaaa/pb"
)

func Parsing() {
	URL := "https://mir-kubikov.ru/catalog/constructors/lego/?page=%d&show_more=Y"

	dr := mirkubikov.NewDir("mir-kubikov/")
	xlsx := mirkubikov.NewXLSX("mir-kubikov/mir-kubikov.xlsx")

	ProductsUrlsAll := make([]string, 0, 1000)
	Bar := pb.StartNew(100)
	Bar.Increment()
	for page := 1; ; page++ {
		url := fmt.Sprintf(URL, page)
		ProductsUrls, NextPageIsExit, ErrParseProductsUrl := mirkubikov.ParseProductsUrl(url)
		if ErrParseProductsUrl != nil {
			panic(ErrParseProductsUrl)
		}
		ProductsUrlsAll = append(ProductsUrlsAll, ProductsUrls...)
		if !NextPageIsExit {
			break
		}
		Bar.Increment()
		// break
		time.Sleep(time.Second)
	}
	Bar.Finish()

	//
	// Парсинг каждого товара
	products := make([]mirkubikov.Product, len(ProductsUrlsAll))
	BarProd := pb.StartNew(len(ProductsUrlsAll))
	for iProductUrl, ProductUrl := range ProductsUrlsAll {
		product, ErrParse := mirkubikov.ParseProduct(ProductUrl)
		if ErrParse != nil {
			panic(ErrParse)
		}

		// Скачивание фото
		dr.MakeDir(product.SKU)
		for ilink, link := range product.PhotoLinks {
			BarProd.Prefix(fmt.Sprintf("[%d/%d]", ilink+1, len(product.PhotoLinks)))
			dr.SavePhoto(link, fmt.Sprintf("mir-kubikov/%s/%d.jpeg", product.SKU, ilink))
			product.PhotoPaths = append(product.PhotoPaths, fmt.Sprintf("mir-kubikov/%s/%d.jpeg", product.SKU, ilink))
			time.Sleep(200 * time.Millisecond)
		}
		products[iProductUrl] = product

		xlsx.WriteXLSX(product)

		BarProd.Increment()
		// break
		time.Sleep(time.Second)
	}
	BarProd.Finish()
	xlsx.CloceAndSaveXLSX()

	// fmt.Println(products[0])

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
