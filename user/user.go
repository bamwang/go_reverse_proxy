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

func (u *User) init(name string) {
	u.name = name
}

func main() {
	u := new(User)
	u.name = "Gopher"
	u.speak()
	u2 := &User{"Gopher2"}
	u2.speak()
	u3 := new(User)
	u3.init("Gopher3")
	u3.speak()
}
