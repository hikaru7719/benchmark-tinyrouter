package main

import "log"

func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	path := "github-openapi.json"
	dErr := DownloadFile(path)
	if dErr != nil {
		return dErr
	}

	rErr := ReadOpenAPI(path, "endpoint.txt")
	if rErr != nil {
		return rErr
	}

	return nil
}
