package textiletorg

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func SavePhoto(link string, FilePath string) error {
	client := &http.Client{}
	// client.Timeout = time.Minute
	req, ErrNewRequest := http.NewRequest(http.MethodGet, link, nil)
	if ErrNewRequest != nil {
		return ErrNewRequest
	}

	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.837 YaBrowser/23.9.4.837 Yowser/2.5 Safari/537.36")

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
