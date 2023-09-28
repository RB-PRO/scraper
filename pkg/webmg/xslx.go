package webmg

import (
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type XLSX struct {
	f    *excelize.File
	line map[string]int
}

// Создать книгу
func NewXLSX(PathFile string, Categorys []Category) *XLSX {
	f := excelize.NewFile()
	// defer f.Close()

	xlsx := XLSX{f: f, line: make(map[string]int)}

	for _, cat := range Categorys {
		f.NewSheet(cat.Name)
		xlsx.line[cat.Name] = 2

		f.SetCellValue(cat.Name, "A1", "Title")        // Заголовок
		f.SetCellValue(cat.Name, "B1", "Slug")         // Статья slug на русском, откуда взята картинка
		f.SetCellValue(cat.Name, "C1", "Description")  // Описание
		f.SetCellValue(cat.Name, "D1", "URL Category") // Ссылка на категорию
		f.SetCellValue(cat.Name, "E1", "Path")         // Путь к картинке в папке
		f.SetCellValue(cat.Name, "F1", "URL Photo")    // Ссылка на картинку в источнике

	}

	f.DeleteSheet("Sheet1")
	f.SaveAs(PathFile)
	return &xlsx
}

// Закрыть и сохранить файл
func (x *XLSX) CloceAndSaveXLSX() {
	x.f.Save()
	x.f.Close()
}

// Вписать данные в книгу
func (x *XLSX) WriteXLSX(cat Category, info Info, photos []Photo) {

	//
	PhotoPath := make([]string, len(photos))
	PhotoUrls := make([]string, len(photos))
	for i := range photos {
		PhotoPath[i] = photos[i].Path
		PhotoUrls[i] = photos[i].URL
	}

	x.f.SetCellValue(cat.Name, "A"+strconv.Itoa(x.line[cat.Name]), info.Title)                   // Заголовок
	x.f.SetCellValue(cat.Name, "B"+strconv.Itoa(x.line[cat.Name]), info.Slug)                    // Статья slug на русском, откуда взята картинка
	x.f.SetCellValue(cat.Name, "C"+strconv.Itoa(x.line[cat.Name]), info.Description)             // Описание
	x.f.SetCellValue(cat.Name, "D"+strconv.Itoa(x.line[cat.Name]), info.URL)                     // Ссылка на категорию
	x.f.SetCellValue(cat.Name, "E"+strconv.Itoa(x.line[cat.Name]), strings.Join(PhotoPath, ";")) // Путь к картинке в папке
	x.f.SetCellValue(cat.Name, "F"+strconv.Itoa(x.line[cat.Name]), strings.Join(PhotoUrls, ";")) // Ссылка на картинку в источнике

	// Иттерирование по строкам
	x.line[cat.Name]++
}
