package main
import (
	"fmt"
)

func callerA(ch chan string){
	ch <- "Hello World"
	close(ch)
}

func callerB(ch chan string){
	ch <- "Ling Huang"
	close(ch)
}

func main(){
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)

	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2{
		select{
		case msg, ok1 = <-a:
			if ok1{
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2{
				fmt.Printf("%s from B\n", msg)
			}
		}
	}
}