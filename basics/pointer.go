package basics
import "fmt"

type person struct {
	personName string
}

func changeStructByPointer(p *person){
	p.personName = "john doe"
	fmt.Println(p)
}

func Pointer() {
	x := 10
	y := &x
	fmt.Println(*y)
	fmt.Printf("%T, %T", x, y)
	jonathan := person{"johnathan doe"}
	fmt.Println(jonathan)
	changeStructByPointer(&jonathan)
	fmt.Println(jonathan)
}
