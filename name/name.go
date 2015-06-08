package main

import "fmt"

func main() {
	var age int
	var name string
	fmt.Printf("あなたの名前を入力してください。\n")
	fmt.Scan(&name)
	fmt.Printf("%sさん、こんにちは。\n", name)
	fmt.Printf("年齢を入力してください。\n")
	n, err := fmt.Scanf("%d", &age)
	fmt.Print(n, err)
	fmt.Printf("今 %d 歳とすると、10年後は %d 歳ですね。\n", age, age+10)
}
