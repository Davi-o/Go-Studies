package datagrouping

import "fmt"

var basicFiveElementsArray [5]int

func basicArrayFormat() {
	basicFiveElementsArray[0] = 1
	basicFiveElementsArray[1] = 10

	fmt.Println(basicFiveElementsArray)
	fmt.Println(len(basicFiveElementsArray))
}
