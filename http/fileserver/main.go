package main

import "net/http"

func main() {
	http.Handle("/fileserver/", http.StripPrefix("/fileserver/", http.FileServer(http.Dir("./testdata"))))

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
