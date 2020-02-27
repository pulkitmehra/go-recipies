package simple

import (
	"fmt"
	"math/rand"
	"time"
)

//FnMain functiuon
func FnMain() {
	randomStringPwd()
	wrapperMain()
}

//middleware func
type (
	adapter func(handler) handler

	handler interface {
		call(b []byte) []byte
	}

	handlerFunc func(b []byte) []byte
)

func (mf handlerFunc) call(b []byte) []byte {
	return mf(b)
}

//middleware func end

func log(mw handler) handler {
	fmt.Println("setting up log")
	return handlerFunc(func(b []byte) []byte {
		fmt.Println("before log")
		resp := mw.call(b)
		fmt.Println("after log")
		return resp
	})
}

func wrapperMain() {
	fn := handlerFunc(func(b []byte) []byte {
		hi := []byte(` Hi`)
		return append(b, hi...)
	})

	v := log(fn).call([]byte(`Pulkit`))

	fmt.Printf("%s \n", string(v))
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
