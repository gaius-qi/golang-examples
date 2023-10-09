package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/go-http-utils/headers"
)

func main() {
	s := server()
	defer s.Close()

	fmt.Println(s.URL)

	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fi.Name())
	if err != nil {
		panic(err)
	}
	part.Write(fileContents)

	if err := writer.Close(); err != nil {
		panic(err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPut, s.URL, body)
	if err != nil {
		panic(err)
	}
	req.Header.Add(headers.ContentType, writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		panic(fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}

	fmt.Println(resp.StatusCode)
}

func server() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, fileHeader, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}

		go func() {
			file, err := fileHeader.Open()
			if err != nil {
				panic(err)
			}
			defer file.Close()
			body, err := io.ReadAll(file)
			if err != nil {
				panic(err)
			}

			fmt.Println("11111111", string(body))
		}()

		go func() {
			file, err := fileHeader.Open()
			if err != nil {
				panic(err)
			}
			defer file.Close()

			body, err := io.ReadAll(file)
			if err != nil {
				panic(err)
			}

			fmt.Println("2222222", string(body))
		}()

		file, err := fileHeader.Open()
		if err != nil {
			panic(err)
		}
		defer file.Close()

		body, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		fmt.Println("3333333333", string(body))

		w.WriteHeader(http.StatusOK)
	})

	return httptest.NewServer(handler)
}
