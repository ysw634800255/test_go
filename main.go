package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "ysw/routers"
)

func init()  {

	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	databases := beego.AppConfig.String("databases")

	//coon := "root:root@tcp(localhost:3306)/test?charset=utf8"
	coon := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + databases + "?charset=utf8"
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", coon)
}

func main() {

	orm.Debug = true

	// 通过 /img/资源路径  可以访问static/img目录的内容
	beego.SetStaticPath("/img","static/img")

	// 通过 /css/资源路径  可以访问static/css目录的内容
	beego.SetStaticPath("/css","static/css")

	// 通过 /js/资源路径  可以访问static/js目录的内容
	beego.SetStaticPath("/js","static/js")

	// 通过 /lib/资源路径  可以访问static/lib目录的内容
	beego.SetStaticPath("/lib","static/lib")

	// 通过 /fonts/资源路径  可以访问static/fonts目录的内容
	beego.SetStaticPath("/fonts","static/fonts")
	beego.Run()
}

