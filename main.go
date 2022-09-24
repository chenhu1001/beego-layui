package main

import (
	"github.com/astaxie/beego"
	"github.com/chenhu1001/beego-layui/models"
	_ "github.com/chenhu1001/beego-layui/routers"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}
