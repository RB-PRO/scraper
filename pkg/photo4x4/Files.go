package photo4x4

import (
	"fmt"
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

	req.Header.Add("authority", "webmg.ru")
	req.Header.Add("accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("cookie", "_ym_uid=1695843846203795824; _ym_d=1695843846; _ym_isad=1; url=https://webmg.ru/wp-content/uploads/2023/05/10007-107-jpg.webp")
	req.Header.Add("referer", "https://webmg.ru/kartinki-na-den-kofe-56-otkrytok/")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "image")
	req.Header.Add("sec-fetch-mode", "no-cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.2271 YaBrowser/23.9.0.2271 Yowser/2.5 Safari/537.36")

	// Выполнить запрос
	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return ErrDo
	}
	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("SavePhoto: wrong status code: %d", res.StatusCode)
	}

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
