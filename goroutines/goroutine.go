package goroutines

import (
	"fmt"
	"sync"
)

var waitingGroup sync.WaitGroup

var f = func(x int, text string) {
	for i := 0; i <= x; i++ {
		fmt.Printf("%s: %d \n", text, i)
	}
}

var y = func(x int, text string){
	f(x, text)
	waitingGroup.Done()
}

func Goroutines() {
	waitingGroup.Add(2)

	go y(5, "i'll just wait ")
	f(6, "running as soon as possible ")

	go y(5, "waiting too ")
	f(4, "also running asap ")

    waitingGroup.Wait()
}