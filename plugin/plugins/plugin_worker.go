package main

import "fmt"

func init() {
	fmt.Println("init worker...")
}

type worker struct{}

func (w worker) Run() {
	fmt.Println("working...")
}

var Worker = worker{}
