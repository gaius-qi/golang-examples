package main

import "fmt"

type App struct {
	Name string
}

func main() {
	var app1 *App
	var app2 App

	app2 = *app1
	fmt.Println(app2)
}
