package main

import (
	"fmt"
	"time"
	"sync"
)

func printNumbers1() {
	for i := 0; i < 10; i++{
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++{
	}
}

func print1() {
	printLetters1()
	printNumbers1()
}

func goPrint1() {
	go printLetters1()
	go printNumbers1()
}


func printNumbers2() {
	for i := 0; i < 10; i++{
		time.Sleep(time.Microsecond)
	}
}

func printLetters2() {
	for i := 'A'; i < 'A'+10; i++{
		time.Sleep(time.Microsecond)
	}
}

func print2() {
	printLetters2()
	printNumbers2()
}

func goPrint2() {
	go printLetters2()
	go printNumbers2()
}

func printNumbers3(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++{
		time.Sleep(time.Microsecond)
		fmt.Printf("%d", i)
	}
	wg.Done()
}

func printLetters3(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++{
		time.Sleep(time.Microsecond)
		fmt.Printf("%c", i)
	}
	wg.Done()
}

func printNumber4(ch chan bool) {
	for i := 0; i < 10; i++{
		time.Sleep(time.Microsecond)
		fmt.Printf("%d", i)
	}
	ch <- true
}

func printLetters4(ch chan bool) {
	for i := 'A'; i < 'A'+10; i++{
		time.Sleep(time.Microsecond)
		fmt.Printf("%c", i)
	}
	ch <- true
}

func main (){
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//	go printLetters3(&wg)
//	go printNumbers3(&wg)
//	wg.Wait()
	ch1, ch2 := make(chan bool), make(chan bool)
	go printNumber4(ch1)
	go printLetters4(ch2)
	<- ch1
	<- ch2
}