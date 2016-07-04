package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type baseController struct {
	beego.Controller
}

func (this *baseController) setupWsForLog(logtype string, logchan chan []byte) {
	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		beego.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	go outLog(ws, logchan, logtype)

	defer func() {
		ws.Close()
		this.Redirect("/", 302)
	}()

	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		beego.Trace(string(p))
	}
}
