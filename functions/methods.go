package functions

import "fmt"

type person struct {
	name string
	surname string
	age int
}

func (p person) wish(something string) string {
	return p.name + " wish " + something
}

func methodCall(){
	someone := person{"john","doe", 45}
	fmt.Println(someone.wish("you a goodnight"));
}