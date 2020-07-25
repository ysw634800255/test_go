package routers

import (
	"github.com/astaxie/beego"
	"ysw/controllers"
	"ysw/controllers/Member"
	"ysw/controllers/Order"
	"ysw/controllers/ysw"
)

func init() {

	//首页
    beego.Router("/", &controllers.IndexController{})
	beego.Router("/welcome", &controllers.WelcomeController{})

	//会员管理
	beego.Router("/member_list", &Member.MemberListController{})
	beego.Router("/member_add", &Member.MemberAddController{})
	beego.Router("/member_add_data", &Member.MemberAdDataController{})

    //订单管理
    beego.Router("/order_list_one", &Order.ListOneController{})
	beego.Router("/order_list_two", &Order.ListTwoController{})

    beego.Router("/ysw", &ysw.YeController{})

	beego.Router("/add_list", &ysw.AddListYeController{})

	beego.Router("/add_ysw", &ysw.AddYeController{})

    beego.Router("/del_ysw", &ysw.DelController{})
}
