package main

import "fmt"

func test() {
	fmt.Println(10)
	panic("panic!!!!!")
	fmt.Println(20)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println(2)
	}()
	fmt.Println(1)
	test()
}
