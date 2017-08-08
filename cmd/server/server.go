package main

import (
	"fmt"
	"github.com/norunners/hellovecty/service"
	"github.com/norunners/hellovecty/util"
	"golang.org/x/net/websocket"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// main sets up a service, then listens and serves json rpc over websocket.
func main() {
	service := service.New()
	server := rpc.NewServer()
	server.Register(service)
	handle := handler(server)

	http.Handle("/ws-rpc", websocket.Handler(handle))
	err := http.ListenAndServe("localhost:1234", nil)
	util.Must(err)
}

// handler is a trivial rpc websocket wrapper.
func handler(server *rpc.Server) websocket.Handler {
	return func(conn *websocket.Conn) {
		fmt.Printf("Serve conn begin.\n")
		server.ServeCodec(jsonrpc.NewServerCodec(conn))
		fmt.Printf("Serve conn end.\n")
	}
}
