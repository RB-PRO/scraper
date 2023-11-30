package fogplay

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func Save(games []Game) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheet := "main"
	f.NewSheet(sheet)
	f.DeleteSheet("Sheet1")

	// Установленное значение ячейки
	f.SetCellValue(sheet, "A1", "Название игры")
	f.MergeCell(sheet, "A1", "A2")
	f.SetCellValue(sheet, "B1", "Кол-во доступных серверов за руб/час")
	f.MergeCell(sheet, "B1", "F1")
	f.SetCellValue(sheet, "B2", "<15")
	f.SetCellValue(sheet, "C2", "15-20")
	f.SetCellValue(sheet, "D2", "20-30")
	f.SetCellValue(sheet, "E2", "30-40")
	f.SetCellValue(sheet, "F2", ">40")

	for igame, game := range games {
		var a, b, c, d, e int
		for _, coast := range game.Coats {
			switch {
			case coast < 15:
				a++
			case coast >= 15 && coast < 20:
				b++
			case coast >= 20 && coast < 30:
				c++
			case coast >= 30 && coast < 40:
				d++
			case coast >= 40:
				e++
			}
		}
		row := strconv.Itoa(igame + 3)
		f.SetCellValue(sheet, "A"+row, game.Name)
		f.SetCellValue(sheet, "B"+row, a)
		f.SetCellValue(sheet, "C"+row, b)
		f.SetCellValue(sheet, "D"+row, c)
		f.SetCellValue(sheet, "E"+row, d)
		f.SetCellValue(sheet, "F"+row, e)
	}
	f.SaveAs("fogplay.xlsx")
}
