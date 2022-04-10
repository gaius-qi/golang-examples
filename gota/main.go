package main

import (
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
)

type User struct {
	Name string
	Age  int
}

func main() {
	f, err := os.Open("./data.csv")
	if err != nil {
		panic(err)
	}

	df1 := dataframe.ReadCSV(f)

	fmt.Println(df1)

	users := []User{{Name: "jack", Age: 18}, {Name: "lucy", Age: 20}}
	df2 := dataframe.LoadStructs(users)

	fmt.Println(df2)
}
