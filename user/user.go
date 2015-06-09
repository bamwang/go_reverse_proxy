package main

/*
User Class
*/
type User struct {
	Name string
}

func (u *User) Speak() {
	println(u.Name)
}

func main() {
	u := new(User)
	u.Name = "Gopher"
	u.Speak()
}
