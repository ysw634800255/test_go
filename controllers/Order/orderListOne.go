package Order

import "github.com/astaxie/beego"

type ListOneController struct {
	beego.Controller
}

func (o *ListOneController) Get()  {
	o.TplName = "Order/order_list_one.html"
}
