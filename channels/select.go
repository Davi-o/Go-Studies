package main 

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go func(){
		for i := 0; i < 2; i++ {
			ch1 <- 42
		}
	}()

	go func(){
		for i := 0; i < 2; i++ {
			ch2 <- 42
		}
	}()

	for i:= 0; i < 2; i++ {
		select{
			case c := <-ch1:
				fmt.Println("Channel one: ", c)
			case c := <-ch2:
				fmt.Println("Channel two: ", c)
		}
	}
}

