package simple

import "fmt"

type (
	aClazz struct {
		val int
	}
)

//PointerMain collection
func PointerMain() {
	n := 2
	getPointer(&n)
	fmt.Println(n)
	newClazz()
}

//you need * to get the value of a pointer, which is called dereferencing the pointer,
//and & to get the memory address of a non-pointer variable
func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := (n * n)
	return &v
}

func newClazz() {
	clz := new(aClazz)
	fmt.Println(clz)
	a := make([]aClazz, 2)
	fmt.Println(a)
}
