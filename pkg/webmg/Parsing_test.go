package webmg

import (
	"fmt"
	"strings"
	"testing"
)

func TestParserPhoto(t *testing.T) {
	url := "https://webmg.ru/kartinki-na-den-kofe-56-otkrytok/"
	photos, info, err := ParsePhoto(url)
	if err != nil {
		t.Error()
	}
	fmt.Println("->", info.Title)
	fmt.Println("->", info.Slug)
	fmt.Println("->", info.URL)
	fmt.Println("->", info.Description)
	fmt.Println("Всего фото:", len(photos))
	fmt.Println(photos[0].URL)
	fmt.Println()

	iPhoto := 0
	dr := NewDir("webmg/")
	FileName := photos[iPhoto].URL
	FileName = strings.ReplaceAll(FileName, "https://webmg.ru/wp-content/uploads/", "")
	FileName = strings.ReplaceAll(FileName, "/", "")

	dr.MakeDir("one" + "/")
	Direction, _ := dr.MakeDir("one" + "/" + "two" + "/")

	fmt.Println(">>>", Direction+"\\"+FileName)
	photos[iPhoto].Path = Direction + "\\" + FileName

	dr.SavePhoto(photos[iPhoto].URL, photos[iPhoto].Path)
	fmt.Println()
}
