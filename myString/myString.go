package main

type MyString string

func (this *MyString) Length() (sum int) {
	for range *this {
		sum++
	}
	return
}

func main() {
	a := MyString("aa")
	println(a.Length())

	b := MyString("知らない")
	println(b.Length())
}
