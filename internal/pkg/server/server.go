package server

import (
	"github.com/gorilla/mux"
	sw "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"net/http"
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

func ServerDocs(host string) error {
	router := mux.NewRouter()
	router.PathPrefix("/swagger").Handler(sw.WrapHandler)
	return http.ListenAndServe(host, router)
}
