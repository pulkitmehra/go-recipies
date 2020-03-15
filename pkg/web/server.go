package web

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/valyala/fasthttp"
)

//MainWeb method for orchestrating
func MainWeb() {
	wrapperMain()
	callServer()
}

func callServer() {
	startServer()
	defer call("http://localhost:5678/close")
	fmt.Println(call("http://localhost:5678/test"))

}

func call(url string) string {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func startServer() {
	done := make(chan bool)
	go func() {
		fasthttp.ListenAndServe(":5678", func(ctx *fasthttp.RequestCtx) {
			switch string(ctx.Path()) {
			case "/test":
				ctx.SetBody([]byte(`recieved!`))
			case "/close":
				ctx.SetBody([]byte(`bye bye!`))
				close(done)
			}
		})

	}()
	fmt.Println("to kill call http://localhost:5678/close")
	<-done
}
