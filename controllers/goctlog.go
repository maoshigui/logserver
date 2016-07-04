package controllers

type GoctlogController struct {
	baseController
}

func (c *GoctlogController) Get() {
	c.Data["IsGoCTLog"] = true
	c.Data["LogType"] = "goctlog"
	c.TplName = "goctlog.html"
}

func (this *GoctlogController) Log() {
	logchan := make(chan []byte, 10)
	this.setupWsForLog("goct", logchan)
}
