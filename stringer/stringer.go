package main

import "fmt"

type Test struct {
	data string
}
type TestIF interface {
	fmt.Stringer
	fmt.GoStringer
}

func (this Test) String() string {
	return this.data + "1"
}

func (this Test) GoString() string {
	return this.data + "2"
}

func main() {
	var aI TestIF
	aI = Test{data: "test"}
	fmt.Println(aI)
	fmt.Printf("%#v", aI)
}
