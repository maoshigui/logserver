package routers

import (
	"logserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	beego.Router("/goctlog", &controllers.GoctlogController{})
	beego.AutoRouter(&controllers.GoctlogController{})

	beego.Router("/mslog", &controllers.MslogController{})
	beego.AutoRouter(&controllers.MslogController{})

	beego.Router("/alllog", &controllers.AlllogController{})
	beego.AutoRouter(&controllers.AlllogController{})
}
