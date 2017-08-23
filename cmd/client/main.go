// Package main initializes the client.
package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/websocket"
	"github.com/norunners/hellovecty/app"
	"github.com/norunners/hellovecty/dao"
	"github.com/norunners/hellovecty/util"
	"net/rpc"
)

// main sets up the client and application and hands it over to vecty for rendering.
func main() {
	conn, err := websocket.Dial("ws://localhost:1234/ws-rpc")
	util.Must(err)
	client := rpc.NewClient(conn)

	dao := dao.New(client)
	app := app.New(dao)

	vecty.SetTitle("Hello Vecty!")
	vecty.RenderBody(app)
}
