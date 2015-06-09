package main

import "fmt"

/*
User Class
*/
type User struct {
	name string
}

func (u *User) Speak() {
	fmt.Println(u.name)
}

// func (u *User) init(name string) *User {
// 	u.name = name
// 	return u
// }

func NewUser(name string) *User {
	return &(User{name: name})
}

func main() {
	// u := new(User).init("Gopher")
	u := NewUser("Gopher")
	u.Speak()
}
