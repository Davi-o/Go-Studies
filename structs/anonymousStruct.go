package structs

import "fmt"

func AnonymousStruct(){
	anon := struct{
		name string
		age int
	}{
		"anon",
		35,
	}

	fmt.Println(anon)
}