package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Book struct {
	Id    string `json: "id"`
	Title string `json: "title"`
}

func main() {
	fileInfos, err := ioutil.ReadDir("./testdata")
	if err != nil {
		fmt.Println("Error in accessing directory:", err)
	}

	var books []Book
	for _, file := range fileInfos {
		if !file.IsDir() {
			b, err := ioutil.ReadFile(filepath.Join("./testdata", file.Name()))
			if err != nil {
				panic(err)
			}

			var book Book
			if err := json.Unmarshal(b, &book); err != nil {
				panic(err)
			}

			books = append(books, book)
			fmt.Println(books)
		}
	}
}
