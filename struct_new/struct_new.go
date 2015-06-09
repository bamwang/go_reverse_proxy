package main

import "fmt"

/*
Point koituhadame
dfasf
*/
type Point struct {
	x int
	y func() int
}

func main() {

	var pt = &Point{x: 30}
	var ppt = &pt
	var pppt = &ppt
	(***pppt).x = 3
	// pt.x = 3
	fmt.Print(pt.x)
}
