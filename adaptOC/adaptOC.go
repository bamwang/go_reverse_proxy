package main

import "fmt"

func main() {
	var sum = 0
	var ary = [...]int{1, 2, 3, 4, 5}
	for _, n := range ary {
		sum += n
	}
	fmt.Println(sum)

	color := make(map[string]string)
	color["key1"] = "red"
	color["key2"] = "blue"
	color["key3"] = "green"
	fmt.Println(color["key2"])
	color["key2"] = "yellow"
	fmt.Println(color["key2"])
	delete(color, "key3")
	fmt.Println(len(color))
	_, e := color["key3"]
	if e {
		fmt.Println("key3 exists")
	} else {
		fmt.Println("key3 does not exist")
	}
	color["key4"] = "white"
	fmt.Println(color)
	for key, value := range color {
		fmt.Println(key, value)
	}
}
