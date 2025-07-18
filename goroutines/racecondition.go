package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"runtime"
)

func AtomicToDealWithRaceCondition() {
	var counter int64
	rtAmount := 10
	
	var wg sync.WaitGroup
	wg.Add(rtAmount)

	for i := 0; i < rtAmount; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			wg.Done()
		}() 	
	}

	wg.Wait()
	fmt.Println(counter)
}

func MutexToDealWithRaceCondition() {
	counter := 0
	rtAmount := 10
	
	var wg sync.WaitGroup
	wg.Add(rtAmount)

	var mutex sync.Mutex

	for i := 0; i < rtAmount; i++ {
		go func() {
			mutex.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mutex.Unlock()
			wg.Done()
		}() 	
	}

	wg.Wait()
	fmt.Println(counter)
}