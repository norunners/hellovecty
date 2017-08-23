// Package main initializes the server.
package main

import (
	"fmt"
	"github.com/norunners/hellovecty/service"
	"github.com/norunners/hellovecty/util"
	"github.com/pdf/websocketrwc"
	"net/http"
	"net/rpc"
)

// main sets up a service, then listens and serves json rpc over websocket.
func main() {
	service := service.New()
	server := rpc.NewServer()
	server.Register(service)
	handler := handler(server)

	http.HandleFunc("/ws-rpc", handler)
	err := http.ListenAndServe("localhost:1234", nil)
	util.Must(err)
}

// handler is a trivial rpc websocket wrapper function.
func handler(server *rpc.Server) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Serve conn begin.")
		conn, err := websocketrwc.Upgrade(res, req, res.Header(), nil)
		util.Must(err)
		server.ServeConn(conn)
		fmt.Println("Serve conn end.")
	}
}
