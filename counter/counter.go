package main

import (
	"bufio"
	"os"
	"time"
)

func numChan() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; true; i++ {
			time.Sleep(100 * time.Millisecond)
			ch <- i
		}
	}()
	return ch
}

func interaptChan() chan int {
	ch := make(chan int)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			ch <- 1
		}
	}()
	return ch
}

func timerChan() <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(10 * time.Second)
		ch <- 1
	}()
	return ch
}

func main() {
	numCh := numChan()
	interaptCh := interaptChan()
	timerCh := timerChan()
	for {
		select {
		case n := <-numCh:
			println(n)
		case <-interaptCh:
			os.Exit(1)
		case <-timerCh:
			os.Exit(1)
		}
	}
}
