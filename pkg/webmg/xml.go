package webmg

import (
	"encoding/xml"
	"io"
	"os"
	"strings"
)

type DataXLM struct {
	Cat    Category `xml:"category"`
	Info   Info     `xml:"info"`
	Photos []Photo  `xml:"photos"`
}
type dataxmls struct {
	Data []dataxml `xml:"Articles"`
}
type dataxml struct {
	Cat  Category `xml:"category"`
	Info Info     `xml:"info"`
	URL  string   `xml:"url"`  // Ссылка на картинку в источнике
	Path string   `xml:"path"` // Путь к картинке в папке
}

func SaveXML(PathNameFile string, data []DataXLM) error {
	xmlFile, err := os.Create(PathNameFile)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	xmlWriter := io.Writer(xmlFile)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")

	var datas dataxmls
	datas.Data = make([]dataxml, len(data))
	for i := range data {
		datas.Data[i].Cat = data[i].Cat
		datas.Data[i].Info = data[i].Info
		URL := make([]string, 0, len(data[i].Photos))
		Path := make([]string, 0, len(data[i].Photos))
		for j := range data[i].Photos {
			URL = append(URL, data[i].Photos[j].URL)
			Path = append(Path, data[i].Photos[j].Path)
		}
		datas.Data[i].URL = strings.Join(URL, ";")
		datas.Data[i].Path = strings.Join(Path, ";")
	}

	if err := enc.Encode(datas); err != nil {
		return err
	}

	xmlFile.Close()
	return nil
}
