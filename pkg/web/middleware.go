package web

import "fmt"

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
