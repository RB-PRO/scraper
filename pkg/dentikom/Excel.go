package dentikom

import (
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Создать книгу
func MakeWorkBook() (*excelize.File, error) {
	// Создать книгу Excel
	f := excelize.NewFile()
	// Create a new sheet.
	_, err := f.NewSheet("main")
	if err != nil {
		return f, err
	}
	f.DeleteSheet("Sheet1")
	return f, nil
}

// Сохранить и закрыть файл
func CloseXlsx(f *excelize.File, filename string) error {
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// Сохранить данные в Xlsx
func SaveXlsx(filename string, prods []Product) error {
	// Создать книгу
	book, makeBookError := MakeWorkBook()
	if makeBookError != nil {
		return makeBookError
	}

	wotkSheet := "main"
	setHead(book, wotkSheet, 1, "Название товара")
	setHead(book, wotkSheet, 2, "Артикул")
	setHead(book, wotkSheet, 3, "Ссылка")
	setHead(book, wotkSheet, 4, "Категория")
	setHead(book, wotkSheet, 5, "Цена")
	setHead(book, wotkSheet, 6, "Картинки")
	setHead(book, wotkSheet, 7, "Описание")

	// // Создаём мапу, которая будет содержать значения номеров колонок
	for indexItem, valItem := range prods {
		setCell(book, wotkSheet, indexItem+2, 1, valItem.Name)
		setCell(book, wotkSheet, indexItem+2, 2, valItem.SKU)
		setCell(book, wotkSheet, indexItem+2, 3, valItem.URL)
		setCell(book, wotkSheet, indexItem+2, 4, valItem.Category)
		setCell(book, wotkSheet, indexItem+2, 5, valItem.Price)
		setCell(book, wotkSheet, indexItem+2, 6, strings.Join(valItem.Images, ";"))
		setCell(book, wotkSheet, indexItem+2, 7, valItem.Deskription)

	}

	// Закрыть книгу
	closeBookError := CloseXlsx(book, filename)
	if closeBookError != nil {
		return closeBookError
	}
	return nil
}

// Вписать значение в ячейку
func setCell(file *excelize.File, wotkSheet string, y, x int, value interface{}) {
	collumnStr, _ := excelize.ColumnNumberToName(x)
	file.SetCellValue(wotkSheet, collumnStr+strconv.Itoa(y), value)
}

// Вписать значение в название колонки
func setHead(file *excelize.File, wotkSheet string, x int, value interface{}) {
	collumnStr, _ := excelize.ColumnNumberToName(x)
	file.SetCellValue(wotkSheet, collumnStr+"1", value)
}
