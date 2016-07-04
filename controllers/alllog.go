package controllers

type AlllogController struct {
	baseController
}

func (c *AlllogController) Get() {
	c.Data["IsAllLog"] = true
	c.Data["LogType"] = "alllog"
	c.TplName = "alllog.html"
}

func (this *AlllogController) GoctLog() {
	logchan := make(chan []byte, 10)
	this.setupWsForLog("all for goct", logchan)
}

func (this *AlllogController) MsLog() {
	logchan := make(chan []byte, 10)
	this.setupWsForLog("all for ms", logchan)
}
