package main

import "fmt"

func main() {
	ns := []int{1, 2, 3}
	ns[1] *= 10
	fmt.Println("ns:", ns)
	fmt.Println("ns:", len(ns), cap(ns))

	var ary [5]int = [...]int{1, 1, 2, 3, 5}
	fmt.Println("ary:", ary)
	fmt.Println("ary:", len(ary), cap(ary))

	slc2 := ary[1:3]
	fmt.Println("slc2:", slc2)
	fmt.Println("slc2:", len(slc2), cap(slc2))

	slc2[0] = 100
	fmt.Println("slc2:", slc2)
	fmt.Println("ary:", ary)

	slc2 = append(slc2, 10)
	fmt.Println("slc2:", slc2)
	fmt.Println("ary:", ary)

	slc2 = append(slc2, 20)
	fmt.Println("slc2:", slc2)
	fmt.Println("ary:", ary)

	slc2 = append(slc2, 30)
	fmt.Println("slc2:", slc2)
	fmt.Println("ary:", ary)

	slc2[0] = 500
	fmt.Println("slc2:", slc2)
	fmt.Println("ary:", ary)
}
