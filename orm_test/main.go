package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "orm_test/models"
	_ "orm_test/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	//orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "scott:password@/orm_test?charset=utf8")
}

func main() {

	o := orm.NewOrm()
	o.Using("default")

	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))

	beego.Run()
}
