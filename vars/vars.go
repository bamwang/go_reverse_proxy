package main

import "fmt"

func main() {
	var a int
	a = 10
	fmt.Print(a)

	b := 20
	fmt.Print(b)

	{

		var a = 30
		fmt.Print(a)
	}
	fmt.Print(a)
}
