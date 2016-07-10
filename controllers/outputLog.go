package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

func outLog(ws *websocket.Conn, logchan chan []byte, logtype string, filterChan chan string) {
	var filters []string

	for {
		select {
		case msg := <-logchan:
			ws.WriteMessage(websocket.TextMessage, msg)
			beego.Trace(logtype, " output msg from logchan")

		case msg := <-filterChan:
			filters = append(filters, msg)
			beego.Trace(filters)

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
