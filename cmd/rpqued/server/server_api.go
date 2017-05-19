package server

import (
	"net"
	"net/http"
	"time"

	"github.com/funkygao/httprouter"
)

type apiServer struct {
	router *httprouter.Router

	httpListener net.Listener
	httpServer   *http.Server
}

func newApiServer() *apiServer {
	router := httprouter.New()
	return &apiServer{
		router: router,
		httpServer: &http.Server{
			Addr:           "httpAddr", // FIXME
			Handler:        router,
			ReadTimeout:    time.Second * 4,
			WriteTimeout:   time.Second * 4,
			MaxHeaderBytes: 1024,
		},
	}
}

func (api *apiServer) Router() *httprouter.Router {
	return api.router
}

func (api *apiServer) Start() error {
	listener, err := net.Listen("tcp", api.httpServer.Addr)
	if err != nil {
		return err
	}

	api.httpListener = listener
	return api.httpServer.Serve(api.httpListener)
}
