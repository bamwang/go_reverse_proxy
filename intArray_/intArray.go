package main

import "github.com/foo/collection"

func main() {
	ary := collection.IntArray{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ary.Map(func(n int) int {
		return n * n
	})
	// .ForEach(func(n int) {
	//     fmt.Println(n)
	// })
}
