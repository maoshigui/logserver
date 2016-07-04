package controllers

type MslogController struct {
	baseController
}

func (c *MslogController) Get() {
	c.Data["IsMSLog"] = true
	c.Data["LogType"] = "mslog"
	c.TplName = "mslog.html"
}

func (this *MslogController) Log() {
	logchan := make(chan []byte, 10)
	this.setupWsForLog("ms", logchan)
}
