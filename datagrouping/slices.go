package datagrouping

import "fmt"

var sliceOfFruits = []string{
	"apple",
	"banana", 
	"pear", 
	"pineapple", 
	"tomato",
} 

func multidimensionalSlice(){
	slice := [][]string{
		{"a","b","c"},
		{"x","y","z"},
		{"g","h","i"},
	}

	for _ , value := range slice{
		fmt.Println(value)
	}

}

func makeSlice(){
	slice := make([]int, 5, 10)
	
	slice[0], slice[1], slice[2], slice[3],	slice[4] = 1, 2 , 3, 4, 5
	
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	fmt.Println("slice:", slice, "lenght:", len(slice), "cap:", cap(slice))
	
	slice = append(slice, 11)
	
	fmt.Println("slice:", slice, "lenght:", len(slice), "cap:", cap(slice))

}

func appendSliceOnSlice(){
	slice := []string{"rice","bean"}
	sliceOfFruits = append(sliceOfFruits, slice...) // the three dots are important, they will unfurl the slice 
	fmt.Println(sliceOfFruits)
}

func sliceSliceExample(){
	slice := sliceOfFruits[2:] // the same as [2:len(slice)]
	fmt.Println(slice)
}

func sliceRangeExample(){
	sliceOfFruits = append(sliceOfFruits,"coconut", "grape")
	for _, value := range sliceOfFruits {
		fmt.Println("fruit:", value)
	}
}

func basicSliceFormat() {
	array := [5]int{1,2,3,4,5}
	fmt.Println(array)
	slice := []int{1,2,3,4,5}
	sliceTwo := append(slice, 6)
	fmt.Println(sliceTwo, len(sliceTwo))
}


