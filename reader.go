package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadFile(path string) error {
	req, _ := http.NewRequest("GET", "https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/api.github.com/api.github.com.json", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("can't get openapi file with %w", err)
	}
	writeErr := WriteFile(path, res.Body)
	defer res.Body.Close()
	if writeErr != nil {
		log.Println(err)
		return fmt.Errorf("can't write file with %s", writeErr)
	}
	return nil
}

func WriteFile(path string, body io.ReadCloser) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	_, copyErr := io.Copy(file, body)
	if copyErr != nil {
		return copyErr
	}
	return nil
}
