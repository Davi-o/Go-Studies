package functions

import "fmt"

func callback() {
	fmt.Println(pairsSlice(Sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
}

func pairsSlice(function func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}

	return function(slice...)

}