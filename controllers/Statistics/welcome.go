package Statistics

import "github.com/astaxie/beego"

type WelcomeController struct {
	beego.Controller
}

func (w *WelcomeController) Get()  {
	w.TplName = "Statistics/welcome_one.html"
}
