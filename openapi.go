package main

import (
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
)

func ReadOpenAPI(path string, textPath string) error {
	t, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		return err
	}

	f, err := os.Create(textPath)
	defer f.Close()
	if err != nil {
		return err
	}

	for p, item := range t.Paths {
		if item.Get != nil {
			printMethodAndPath(f, "GET", p)
		}

		if item.Post != nil {
			printMethodAndPath(f, "POST", p)
		}

		if item.Put != nil {
			printMethodAndPath(f, "PUT", p)
		}

		if item.Patch != nil {
			printMethodAndPath(f, "PATCH", p)
		}

		if item.Delete != nil {
			printMethodAndPath(f, "DELETE", p)
		}

		if item.Connect != nil {
			printMethodAndPath(f, "CONNECT", p)
		}

		if item.Head != nil {
			printMethodAndPath(f, "HEAD", p)
		}

		if item.Options != nil {
			printMethodAndPath(f, "OPTIONS", p)
		}
	}
	return nil
}

func printMethodAndPath(file *os.File, method string, path string) {
	fmt.Fprintf(file, "%s %s\n", method, path)
}
