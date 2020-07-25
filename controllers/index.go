package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}
func (l *IndexController) Get() {
	l.Data["Website"] = "beego.me"
	l.Data["Email"] = "ysw634800255@gmail.com"
	l.TplName = "index.html"
}

//我的桌面
type WelcomeController struct {
	beego.Controller
}

func (w *WelcomeController) Get() {
	w.TplName = "welcome.html"
}
