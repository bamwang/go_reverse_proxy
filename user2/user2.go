package main

/*
User Class
*/
type User struct {
	name string
}

func (u *User) speak() {
	println(u.name)
}

func (u *User) init(name string) *User {
	u.name = name
	return u
}

func main() {
	u := new(User).init("Gopher")
	u.speak()
}
