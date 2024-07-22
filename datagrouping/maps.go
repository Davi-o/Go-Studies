package datagrouping
import "fmt"

func mapsRange(){
	anyMap := map[int]string{
		1:"a",
		2:"b",
		3:"c",
		4:"d",
	}
	
	delete(anyMap, 1)

	for k, v := range anyMap{
		fmt.Println(k, v)
	}
}

func maps( ){
	agenda := map[string]int{
		"john": 1234456,	
		"mary": 1234456,	
	}
	agenda["mark"] = 11112222

	fmt.Println(agenda)
} 

func commaOk(){
	names := map[int]string{
		1:"john",
		2:"mary",
		3:"joseph",
	}
	
	exists, ok := names[3]

	if(ok){
		fmt.Println(exists)
	}
}