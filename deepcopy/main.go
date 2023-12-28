package main

import "fmt"

type User struct {
	Username string
	Password string
}

func main() {
	user1 := &User{"user1", "pass1"}
	user2 := new(User)
	*user2 = *user1

	user2.Username = "user2"
	user2.Password = "pass2"

	fmt.Println(user1)
	fmt.Println(user2)
}
