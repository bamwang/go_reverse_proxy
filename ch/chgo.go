package main

import "fmt"

func gen() chan string {
	ch := make(chan string, 1)
	go func() {
		for {
			ch <- "y"
		}
	}()
	return ch
}
func main() {
	g := gen()
	fmt.Println(<-g)
	fmt.Println(<-g)
	fmt.Println(<-g)

}
