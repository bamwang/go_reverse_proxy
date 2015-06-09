package main

type MyString string

func (this *MyString) Length() int {
	sum := 0
	for range *this {
		sum++
	}
	return sum
}

func main() {
	a := MyString("aa")
	println(a.Length())

	b := MyString("知らない")
	println(b.Length())
}
