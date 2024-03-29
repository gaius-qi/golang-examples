package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwordOK := "admin"
	passwordERR := "password"

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	encodePW := string(hash)
	log.Println(encodePW)

	if err := bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOK)); err != nil {
		log.Println("password wrong")
	} else {
		log.Println("password ok")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordERR)); err != nil {
		log.Println("password wrong")
	} else {
		log.Println("password ok")
	}

	c, err := bcrypt.Cost([]byte(encodePW))
	if err != nil {
		log.Fatal(err)
	}

	if c != bcrypt.DefaultCost {
		log.Printf("expected cost is %d, bug got %d", bcrypt.DefaultCost, c)
	}
}
