package basics

import (
	"fmt"
	"runtime"
)
const (
	_ = iota
	g
	h
	i
	j
	k
	l = iota + 10
)
var variable interface{}

func Basics() {
	switchCase()
	switchCaseByType()
	forLoop()
	primitiveTypesOrSo()
}

func switchCaseByType() {
	variable = "batman"

	switch variable.(type) {
	case float64:
		fmt.Println("its a float number")
	case int:
		fmt.Println("its a integer number")
	case string:
		fmt.Printf("the string contains: %v", variable)
	}
}

func switchCase() {
	
	switch false {
	case g < l:
		fmt.Printf("%v is smaller than %v", g, l)
	case l == k:
		fmt.Printf("%v is different from %v \n", l, k)
		fallthrough
	case l > g:
		fmt.Printf("and %v is greater than %v", l, g)
	default:
		fmt.Println("thats all folks")
	}
}

func forLoop(){
	for i := 50; i <= 60; i++ {
		fmt.Printf("%c \n", i)
	}
}

func primitiveTypesOrSo() {
	x := "e"
	y := "é"
	z := "😁"

	a := []byte(x)
	b := []byte(y)
	c := []byte(z)

	fmt.Println(a, b, c)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println(g, h, i)
	fmt.Printf("\n %T \t %T \t %T", g, a, y)
}
