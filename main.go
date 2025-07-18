package main

import (
	"course/basics"
	"course/functions"
	"course/goroutines"
	"course/jsonHandling"
	"course/sorting"
)

func main() {
	functions.Functions()
	basics.Basics()
	basics.Pointer()
	jsonHandling.Handle()
	sorting.Sort()
	goroutines.Goroutines()
	goroutines.MutexToDealWithRaceCondition()
	goroutines.AtomicToDealWithRaceCondition()
}