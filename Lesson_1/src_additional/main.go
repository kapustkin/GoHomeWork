package main

import (
	"fmt"
)

func main() {
	p := fmt.Printf

	storage := map[string]string {
		"123": "321",
	}
	p(storage["123"])
}

func 

// Взаимодействие с хранилищем
type Shortener interface {
	Shorten(url string) string

	Resolve(url string) string
}

// Хранилище урлов
type UrlStruct struct {
	fullUrl, shortUrl string
}
	