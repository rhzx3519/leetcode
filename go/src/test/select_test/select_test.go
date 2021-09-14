package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel1(t *testing.T) {
	c := make(chan int)
	defer close(c)
	go func() { c <- 3 + 4 }()
	i := <-c
	fmt.Println(i)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func TestFibonacci(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func TestTimeout(t *testing.T) {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}

func TestSwitchCase2(t *testing.T) {
	var c1 = make(chan int)
	var quit = make(chan int)
	var x int

	go func() {
		time.Sleep(time.Second * 2)
		quit <- 1
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c1)
		}
		//c1 <- 2
	}()

	for {
		select {
		case c1 <- x:
			fmt.Println(x)
		case <- quit:
			return
		}
	}
}

func TestQuit(t *testing.T) {
	quit := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		quit <- 1
	}()
	fmt.Println(<- quit)
}

func TestWait(t *testing.T) {
	quit := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		quit <- 1
	}()

	x := <- quit
	fmt.Println(x)
	//select {
	//case x := <- quit:
	//	fmt.Println(x)
	//	break
	//}
}

