package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/scott7950/beego_practice/orm_test/routers"
)

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "scott:password@/orm_test?charset=utf8")
}

func main() {
	orm.Debug = true

	beego.Run()
}
