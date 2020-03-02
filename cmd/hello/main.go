package main

import (
	"github.com/pulkitmehra/go-recipies/pkg/simple"
	"github.com/pulkitmehra/go-recipies/pkg/web"
)

func main() {

	simple.FnMain()
	web.MainWeb()

	/* simple.CopyArrayAndSlice([]int{1, 2, 3, 4}, []int{5, 6, 7, 8})
	simple.SortMain()
	simple.PointerMain()
	simple.StringMain() */
}
