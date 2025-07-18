package main

import "fmt"

func main(){
	ch := make(chan int)
	go loop(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func loop(s chan<- int) {
	for i := 0; i< 15; i++ {
		s <- i
	}
	close(s)
}