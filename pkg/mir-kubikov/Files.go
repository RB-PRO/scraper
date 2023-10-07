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

	req.Header.Add("Accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8")
	req.Header.Add("Accept-Language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "BITRIX_SM_SALE_UID=1342262728; _ga=GA1.1.2087411697.1696525926; tmr_lvid=a713f136e8c395fa5febb356d7571f8e; tmr_lvidTS=1696525925788; _ym_uid=1696525926329875913; _ym_d=1696525926; _gpVisits={\"isFirstVisitDomain\":true,\"idContainer\":\"10002548\"}; popmechanic_sbjs_migrations=popmechanic_1418474375998%3D1%7C%7C%7C1471519752600%3D1%7C%7C%7C1471519752605%3D1; adspire_uid=AS.1308471874.1696525927; ads_adware=true; _ym_isad=1; PHPSESSID=GYGNlgVTNcY21AyU9Eg9Rv1XWr53zuUO; _ym_visorc=b; _gp10002548={\"hits\":89,\"vc\":1,\"ac\":1,\"a6\":1}; _ga_61W59B3053=GS1.1.1696623419.6.1.1696625026.0.0.0; mindboxDeviceUUID=277600e1-2469-4aea-8214-06dfbed47d1f; directCrm-session=%7B%22deviceGuid%22%3A%22277600e1-2469-4aea-8214-06dfbed47d1f%22%7D")
	req.Header.Add("Referer", "https://mir-kubikov.ru/")
	req.Header.Add("Sec-Fetch-Dest", "image")
	req.Header.Add("Sec-Fetch-Mode", "no-cors")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.967 YaBrowser/23.9.1.967 Yowser/2.5 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
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
