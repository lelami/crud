package server

import (
	"github.com/valyala/fasthttp"
	"time"
)

var server *fasthttp.Server

func Run(addr string, handler fasthttp.RequestHandler) error {
	server = &fasthttp.Server{
		Handler:      handler,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
	}

	return server.ListenAndServe(addr)
}

func Stop() error {
	return server.Shutdown()
}
