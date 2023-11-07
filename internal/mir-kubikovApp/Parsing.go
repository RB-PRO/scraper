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
	core, _ := mirkubikov.NewCore()
	defer core.DeleteCore()
	defer xlsx.CloceAndSaveXLSX()

	for page := 1; ; page++ {
		if page < 72 {
			continue
		}
		url := fmt.Sprintf(URL, page)
		ProductsUrls, NextPageIsExit, ErrParseProductsUrl := mirkubikov.ParseProductsUrl(url)
		if ErrParseProductsUrl != nil {
			panic(ErrParseProductsUrl)
		}

		BarProd := pb.StartNew(len(ProductsUrls))
		for _, ProductUrl := range ProductsUrls {
			// product, ErrParse := mirkubikov.ParseProduct(ProductUrl)
			product, ErrParse := core.Parsing(ProductUrl)
			product.URL = ProductUrl
			if ErrParse != nil {
				//panic(ErrParse)
				continue
			}

			// Скачивание фото
			dr.MakeDir(product.SKU)
			if product.SKU == "" {
				fmt.Println("---SKU nill>", product.URL)
			}
			if len(product.PhotoLinks) == 0 {
				fmt.Println("---0>", product.URL)
			}
			for ilink, link := range product.PhotoLinks {
				BarProd.Prefix(fmt.Sprintf("[%d][%d/%d]", page, ilink+1, len(product.PhotoLinks)))
				ErrSavePhoto := dr.SavePhoto(link, fmt.Sprintf("mir-kubikov/%s/%d.jpeg", product.SKU, ilink))
				if ErrSavePhoto != nil {
					fmt.Println("--->", ErrSavePhoto)
				}
				product.PhotoPaths = append(product.PhotoPaths, fmt.Sprintf("mir-kubikov/%s/%d.jpeg", product.SKU, ilink))
				time.Sleep(500 * time.Millisecond)
			}

			xlsx.WriteXLSX(product)

			BarProd.Increment()
			// break
			time.Sleep(time.Second)
		}
		BarProd.Finish()

		// Если дальше страниц не будет
		if !NextPageIsExit {
			break
		}
		// break
		time.Sleep(time.Second)
	}

	//
	// Парсинг каждого товара

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
