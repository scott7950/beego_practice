package routers

import (
	"github.com/astaxie/beego"
	"orm_test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
