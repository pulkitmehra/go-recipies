package main

import (
	"fmt"

	"github.com/pulkitmehra/go-recipies/pkg/simple"
)

func main() {
	fmt.Println(simple.Hello())
	simple.FnMain()
	/* simple.CopyArrayAndSlice([]int{1, 2, 3, 4}, []int{5, 6, 7, 8})
	simple.SortMain()
	simple.PointerMain()
	simple.StringMain() */
}
