package main

import (
	"fmt"
	"time"
)

func thrower(ch chan int){
	for i := 0; i < 5; i++{
		ch <- i
		fmt.Println("Threw >>", i)
	}
}

func catcher(ch chan int){
	for i := range(ch){
		fmt.Println("Catch <<", i)
	}
}

func main() {
	ch := make(chan int, 3)
	go thrower(ch)
	go catcher(ch)
	time.Sleep(100 * time.Millisecond)
}