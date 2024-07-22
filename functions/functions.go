package functions

import "fmt"

func Functions() {
	fmt.Println(Sum(1,2))
	fmt.Println(LargeSum(10, 15, 20, 5))
	deferExample()
	methodCall()
	interfaceMethodCall()
	anonymous()
	callReturnedFunction()
	callback()
	closuredContextFunction()
}



