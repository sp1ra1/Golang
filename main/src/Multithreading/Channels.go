package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// The select statement lets a goroutine wait on multiple communication operations(here is to listen to two channels)
// A select blocks until one of its cases can run, then execute that case. Randomly picked one if multiple are ready
func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	go say("gg")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	// buffered channel
	channel := make(chan int, 3)
	channel <- 2
	channel <- 3
	channel <- 4
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	// sender could close a channel, receiver could test whether a channel has been closed by: v, ok := <- ch
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	// The range could keep receiving values from the channel until it is closed
	for i := range ch {
		fmt.Println(i)
	}

	// select statement
	chan1 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-chan1)
		}
		quit <- 0
	}()
	fibonacciSelect(chan1, quit)
}
