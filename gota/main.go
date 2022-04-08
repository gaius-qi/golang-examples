package main

import (
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	f, err := os.Open("./data.csv")
	if err != nil {
		panic(err)
	}

	df := dataframe.ReadCSV(f)
	fmt.Println(df)
}
