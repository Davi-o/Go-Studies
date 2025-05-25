package jsonHandling
import (
	"encoding/json"
	"fmt"
	"os"
)

func MarshallingUnmarshalling(){
	marshalling()
	unmarshalling()
}

func marshalling() {
	type structExample struct {
		Id     int
		Name   string
	}
	group := structExample{ 1, "Example" }
	sliceOfByte, error := json.Marshal(group)
	if error != nil {
		fmt.Println("error:", error)
	}
	os.Stdout.Write(sliceOfByte)
}

func unmarshalling(){
	type Infos struct {
		Name string `json:"nome"`
		Age  int `json:"idade"`
		Car  any `json:"carro"`
	}
	sliceOfBytes := []byte(`[{"nome":"John", "idade":30, "carro":"uno"}]`)
	var infos []Infos
	error := json.Unmarshal(sliceOfBytes, &infos)
	if error != nil {
		fmt.Println("error: ", error)
	}
	fmt.Printf("%+v", infos)

}