package sorting

import (
	"fmt"
	"sort"
)

func Sort() {
	sortingAlphabetically()
	sortingNumber()
	CustomSort()
	HashingWithBCrypt()
}

func sortingAlphabetically() {
	sliceOfString := []string{"b first", "c second", "a third"}
	sort.Strings(sliceOfString)
	fmt.Println(sliceOfString)
}

func sortingNumber() {
	sliceOfInt := []int{100, 3, 8, 1, -1, 3000, 10, 2}
	sort.Ints(sliceOfInt)
	fmt.Println(sliceOfInt)
}
