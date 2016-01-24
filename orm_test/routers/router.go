package routers

import (
	"github.com/astaxie/beego"
	"github.com/scott7950/beego_practice/orm_test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.MainController{}, "*:Login")
}
