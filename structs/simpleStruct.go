package structs

import "fmt"

type Address struct {
	Streetname string
	Zipcode string
	Number  int
}

func SimpleStruct() {
	firstHouse := Address{
		Streetname: "street name",
		Zipcode: "9120300",
		Number:  312,
	}

	secondHouse := Address{
		Streetname: "street name",
		Zipcode: "9120310",
		Number:  210,
	}

	fmt.Println(firstHouse, secondHouse)

}