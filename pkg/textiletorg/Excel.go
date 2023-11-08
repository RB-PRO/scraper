package textiletorg

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

/*
func WriteOneLine(f *excelize.File, ssheet string, row int, SearchBasicRes SearchBasicResponse, SearchBasicIndex int, GetPartsRemainsByCodeRes GetPartsRemainsByCodeResponse, GetPartsRemainsByCodeIndex int) {
	// SearchBasic
	writeHeadOne(f, ssheet, 1, row, SearchBasicRes.Data.Items[SearchBasicIndex].Code, "")
}
*/

// Сохранить данные в Xlsx
func SaveXlsx(filename string, prods []Product) error {
	// Создать книгу
	book, makeBookError := MakeWorkBook()
	if makeBookError != nil {
		return makeBookError
	}

	wotkSheet := "main"
	setHead(book, wotkSheet, 1, "Название товара")      // Catalog
	setHead(book, wotkSheet, 2, "Артикул")              // PodCatalog
	setHead(book, wotkSheet, 3, "Ссылка")               // PodCatalog
	setHead(book, wotkSheet, 4, "Цена")                 // Section
	setHead(book, wotkSheet, 5, "Ссылки на картинки")   // PodSection
	setHead(book, wotkSheet, 6, "Ссылки на сохранёнки") // PodSection
	setHead(book, wotkSheet, 7, "Описание")             // Name

	startIndexCollumn := 8

	// // Создаём мапу, которая будет содержать значения номеров колонок
	colName := make(map[string]int)
	for indexItem, valItem := range prods {
		setCell(book, wotkSheet, indexItem+2, 1, valItem.Name)
		setCell(book, wotkSheet, indexItem+2, 2, valItem.SKU)
		setCell(book, wotkSheet, indexItem+2, 3, valItem.URL)
		setCell(book, wotkSheet, indexItem+2, 4, valItem.Price)
		setCell(book, wotkSheet, indexItem+2, 5, strings.Join(valItem.PhotoURL, ";"))
		setCell(book, wotkSheet, indexItem+2, 6, strings.Join(valItem.PhotoDir, ";"))
		setCell(book, wotkSheet, indexItem+2, 7, valItem.Deskription)

		// Обработка мапы доп полей
		for key, val := range valItem.TS {
			if _, ok := colName[key]; ok { // Если такое значение существует(т.е. существует колонка)
				setCell(book, wotkSheet, indexItem+2, colName[key], val)
			} else {
				colName[key] = startIndexCollumn
				setHead(book, wotkSheet, colName[key], key)
				setCell(book, wotkSheet, indexItem+2, colName[key], val)
				startIndexCollumn++
			}
		}
	}

	// Закрыть книгу
	closeBookError := CloseXlsx(book, filename)
	if closeBookError != nil {
		return closeBookError
	}
	return nil
}

/*
// Добавить ссылку в массив строк
func addURL_toLink(links []string) []string {
	for index := range links {
		links[index] = URL + links[index]
	}
	return links
}
*/

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
