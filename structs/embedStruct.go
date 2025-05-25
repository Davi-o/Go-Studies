package structs

import "fmt"

type building struct {
	room []Address
	floors int
}

func EmbedStruct(){
	firstBuilding := building {
		room: []Address{
			{
				Streetname: "street name",
				Zipcode: "9120300",
				Number:  150,
			},
		}, 
		floors: 3,
	}

	secondBuilding := building {
		room: []Address{
			{
				"street name",
				"9120420",
				22,
			},
			{
				"street name",
				"9120420",
				32,
			},
			{
				"street name",
				"9120420",
				42,
			},
		}, 
	}
	fmt.Println(firstBuilding)
	
	for _, value := range secondBuilding.room{
		fmt.Println(value)
	}

}

func MapEmbedStruct(){
	type someone struct{
		name string
		favoriteFoods []string
	}

	someMap := make(map[string]someone)

	someMap["first"] = someone{
		"john",
		[]string{
			 "lasagna",
			 "pasta",
			 "cake",
		},
	}

	someMap["second"] = someone{
		"mary",
		[]string{
			 "lasagna",
			 "pasta",
			 "cake",
		},
	}

	for _, anyone := range someMap{
		fmt.Println(anyone.name)
		for _, name := range anyone.favoriteFoods{
			fmt.Println(name)
		}
	}
}