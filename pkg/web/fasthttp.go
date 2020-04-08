package web

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type (
	//MyHandler func
	MyHandler struct {
		foobar string
	}
)

//HandleFastHTTP override func
func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
			ctx.Path(), h.foobar)
	case "/bar":
		fmt.Fprintf(ctx, "Hello, world bar! Requested path is %q. Foobar is %q",
			ctx.Path(), h.foobar)
	default:

	}
}

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

//MainFastHTTP func
func MainFastHTTP2() {
	myHandler := &MyHandler{
		foobar: "foobar",
	}
	fmt.Println("starting server...")
	fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)
}

//MainFastHTTP func
func MainFastHTTP() {
	route := routing.New()
	route.Get("/", func(c *routing.Context) error {
		fmt.Fprintf(c, "Hello, world!")
		return nil
	})
}
