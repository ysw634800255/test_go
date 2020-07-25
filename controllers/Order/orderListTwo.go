package Order

import "github.com/astaxie/beego"

type ListTwoController struct {
	beego.Controller
}

func (o *ListTwoController) Get()  {
	o.TplName = "Order/order_list_two.html"
}
