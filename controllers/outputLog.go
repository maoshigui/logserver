package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

func outLog(ws *websocket.Conn, logchan chan []byte, logtype string) {
	for {
		select {
		case msg := <-logchan:
			ws.WriteMessage(websocket.TextMessage, msg)
			beego.Trace(logtype, " output msg from logchan")

		case <-time.After(time.Second * 5):
			if ws != nil {
				if ws.WriteMessage(websocket.TextMessage, []byte("hello "+logtype)) != nil {
					beego.Trace(logtype, "websocket closed, i will quit now")
					return
				}
				beego.Trace(logtype, "out put msg:", "hello", logtype)
			}
		}
	}
}
