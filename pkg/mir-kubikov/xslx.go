package mirkubikov

import (
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type XLSX struct {
	f         *excelize.File
	line      int
	SheetName string
}

// Создать книгу
func NewXLSX(PathNameFile string) *XLSX {
	f := excelize.NewFile()
	// defer f.Close()

	SheetName := "main"
	f.NewSheet(SheetName)

	f.SetCellValue(SheetName, "A1", "Name")        // Название товара
	f.SetCellValue(SheetName, "B1", "SKU")         // Артикул
	f.SetCellValue(SheetName, "C1", "Price")       // Цена
	f.SetCellValue(SheetName, "D1", "Description") // Описание товара
	f.SetCellValue(SheetName, "E1", "PhotoLinks")  // Ссылки на фото источника
	f.SetCellValue(SheetName, "F1", "PhotoPaths")  // Ссылки на локальные файлы
	f.SetCellValue(SheetName, "G1", "URL")         // Ссылки на товар

	f.DeleteSheet("Sheet1")
	f.SaveAs(PathNameFile)
	return &XLSX{f: f, line: 2, SheetName: SheetName}
}

// Закрыть и сохранить файл
func (x *XLSX) CloceAndSaveXLSX() {
	x.f.Save()
	x.f.Close()
}

// Вписать данные по товару в книгу
func (x *XLSX) WriteXLSX(Products Product) {

	// Записать данные
	x.f.SetCellValue(x.SheetName, "A"+strconv.Itoa(x.line), Products.Name)
	x.f.SetCellValue(x.SheetName, "B"+strconv.Itoa(x.line), Products.SKU)
	x.f.SetCellValue(x.SheetName, "C"+strconv.Itoa(x.line), Products.Price)
	x.f.SetCellValue(x.SheetName, "D"+strconv.Itoa(x.line), Products.Description)
	x.f.SetCellValue(x.SheetName, "E"+strconv.Itoa(x.line), strings.Join(Products.PhotoLinks, ";"))
	x.f.SetCellValue(x.SheetName, "F"+strconv.Itoa(x.line), strings.Join(Products.PhotoPaths, ";"))
	x.f.SetCellValue(x.SheetName, "G"+strconv.Itoa(x.line), Products.URL)

	x.line++ // Иттерирование по строкам
}
