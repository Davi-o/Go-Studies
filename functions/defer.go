package functions
import "fmt"

func deferExample() {
	defer fmt.Println("this will be showed after")
	fmt.Println("and this will be showed before")
}
