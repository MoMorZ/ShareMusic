package routers

import (
	"github.com/astaxie/beego"
	"sharemusic/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
