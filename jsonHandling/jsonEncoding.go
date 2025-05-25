package jsonHandling

import (
	"course/structs"
	"os"
	"encoding/json"
)
	
func jsonEncoding(){
	address := structs.Address{
		Streetname: "some random street",
		Zipcode: "9000001",
		Number:  808,
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(address)
}