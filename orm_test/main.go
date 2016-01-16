package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"orm_test/models"
	_ "orm_test/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "scott:password@/orm_test?charset=utf8")
}

func main() {
	orm.Debug = true

	o := orm.NewOrm()
	o.Using("default")

	//var prof *models.Profile
	profile := new(models.Profile)
	profile.Age = 30

	user := new(models.User)
	user.Profile = profile
	user.Name = "slene"

	profile.User = user

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))

	beego.Run()
}
