package routers

import (
	"github.com/astaxie/beego"
	"src/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
