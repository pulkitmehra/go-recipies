package web

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

//MainWeb method for orchestrating
func MainWeb() {
	startServer()
}

func startServer() {
	done := make(chan bool)
	go func() {
		fasthttp.ListenAndServe(":8081", func(ctx *fasthttp.RequestCtx) {
			switch string(ctx.Path()) {
			case "/test":
				ctx.SetBody([]byte(`recieved!`))
			case "/close":
				ctx.SetBody([]byte(`bye bye!`))
				close(done)
			}
		})

	}()
	fmt.Println("to kill call http://localhost:8081/close")
	<-done
}
