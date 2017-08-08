package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/websocket"
	"github.com/norunners/hellovecty/app"
	"github.com/norunners/hellovecty/dao"
	"github.com/norunners/hellovecty/util"
	"net/rpc/jsonrpc"
)

// main sets up the application and hands it to vecty to render.
func main() {
	conn, err := websocket.Dial("ws://localhost:1234/ws-rpc")
	util.Must(err)
	client := jsonrpc.NewClient(conn)

	dao := dao.New(client)
	app := app.New(dao)

	vecty.SetTitle("Hello Vecty!")
	vecty.RenderBody(app)
}
