package simple

import "fmt"

//Hello to outside word
func Hello() string {
	return "hello"
}

//CopyArrayAndSlice helper fn
func CopyArrayAndSlice(a, b []int) {
	copy(a, b)
	fmt.Printf("a %v \n", a)
	fmt.Printf("b %v \n", b)
}
