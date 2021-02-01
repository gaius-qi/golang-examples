package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/avast/retry-go"
)

func main() {
	num := 0
	url := "http://dsafdsafdsafdsa.com"
	var body []byte

	err := retry.Do(
		func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			num += 1
			fmt.Printf("num: %v\n", num)
			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("success:", string(body))

}
