package simple

import (
	"fmt"
	"math/rand"
	"time"
)

//FnMain functiuon
func FnMain() {
	randomStringPwd()
}

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

func randomStringPwd() {

	MIN := 0
	MAX := 94
	SEED := time.Now().Unix()
	rand.Seed(SEED)
	var length int64 = 8

	startChar := "!"
	var i int64 = 1

	for {
		myRand := rand.Intn(MAX-MIN) + MIN
		newChar := string(startChar[0] + byte(myRand))
		fmt.Print(newChar)
		if i == length {
			break
		}
		i++
	}
	fmt.Println()
}
