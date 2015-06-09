package main

import "fmt"

/*
var point struct {
	x int
	y int
}
*/

type Point struct {
	x int
	y int
}

func main() {
	/*	point.x = 30
		point.y = 20
		fmt.Printf("x=%d, y=%d\n", point.x, point.y)
	*/
	pt := Point{x: 30}
	fmt.Printf("x=%d, y=%d\n", pt.x, pt.y)
}
