package mirkubikov

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Фунционал файл отвечает за работу директорией, создание, отслеживание, скачивание данных по позициям

// Структура управления файловой системы парсера
//
// Под капотом директория пути до самой папки с фотографиями
// и методы, которые отвечают за сохранение данных
type Direction struct {
	zeroPath string // Путь до папки, в которую будут сохраняться данные
}

// Созждать обхект работы с файловой системой
func NewDir(zeroPath string) *Direction {
	dr := &Direction{zeroPath: zeroPath}
	dr.MakeDir(zeroPath) // psd add error
	return dr
}

// Пересоздать папку
func (dr *Direction) MakeDir(Path string) (string, error) {

	// Абсолютный путь до папки. Если его нет, то удаляем всё
	absFolderPath, _ := filepath.Abs(dr.zeroPath + Path)

	// Если папка существует - удаляем
	if _, err := os.Stat(absFolderPath); err == nil {
		os.RemoveAll(dr.zeroPath + Path)
	}

	// Создание пути
	ErrMkdirAll := os.MkdirAll(dr.zeroPath+Path, 0777)
	if ErrMkdirAll != nil {
		return "", ErrMkdirAll
	}
	return absFolderPath, nil
}

func (dr *Direction) SavePhoto(link string, FilePath string) error {
	client := &http.Client{}
	// client.Timeout = time.Minute
	req, ErrNewRequest := http.NewRequest(http.MethodGet, link, nil)
	if ErrNewRequest != nil {
		return ErrNewRequest
	}

	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("Referer", "https://mir-kubikov.ru/")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.967 YaBrowser/23.9.1.967 Yowser/2.5 Safari/537.36")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

	// Выполнить запрос
	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return ErrDo
	}
	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// // В случае положительного результата
	// if res.StatusCode != http.StatusOK {
	// 	return fmt.Errorf("SavePhoto: wrong status code: %d", res.StatusCode)
	// }

	//open a file for writing
	file, ErrCreate := os.Create(FilePath)
	if ErrCreate != nil {
		return ErrCreate
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, ErrCopy := io.Copy(file, res.Body)
	if ErrCopy != nil {
		return ErrCopy
	}
	return nil
}
