package structs

import "fmt"

type address struct {
	streetname string
	zipcode string
	number  int
}

func SimpleStruct() {
	firstHouse := address{
		streetname: "street name",
		zipcode: "9120300",
		number:  312,
	}

	secondHouse := address{
		streetname: "street name",
		zipcode: "9120310",
		number:  210,
	}

	fmt.Println(firstHouse, secondHouse)

}