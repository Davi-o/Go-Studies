package main

import "fmt"

func main() {
	ch := make(chan int)
	go send(ch)
	receive(ch)
}

func send(s chan<- int) {
	s <- 42
}

func receive(r <-chan int) {
	fmt.Println(<-r)
}