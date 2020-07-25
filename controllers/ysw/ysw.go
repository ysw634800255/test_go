package ysw

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"unsafe"
	"ysw/models"
)

type YeController struct {
	beego.Controller
}

func (y *YeController) Get() {
	o := orm.NewOrm()
	//age := y.GetString("age")
	//name := y.GetString("name")
	//插入单条
	//ysw := models.Ysw{ Age:age, Name:name}
	//id, err := o.Insert(&ysw)
	//fmt.Println(id)
	//fmt.Println(err)

	//插入多条
	//ysw := []models.Ysw{{Age:"11",Name:"aa-aa"},{Age:"22",Name:"bb-bb"},{Age:"33",Name:"cc-cc"}}
	//count, err := o.InsertMulti(100, ysw)
	//fmt.Println(count)
	//fmt.Println(err)

	//查询单条
	//yswData := models.Ysw{Age:"2"}
	//err := o.Read(&yswData, "age")
	//if err != nil {
	//	fmt.Println("不存在的数据")
	//}
	//y.Data["Id"] = yswData.Id
	//y.Data["Age"] = yswData.Age
	//y.Data["Name"] = yswData.Name

	//复杂查询
	//yswData := models.Ysw{}
	//var yswArray []models.Ysw
	yswArray := []models.Ysw{}

	ysw := new(models.Ysw)
	qs := o.QueryTable(ysw)

	//qs.Filter("name__gt", "ysw").All(&yswData)
	//qs.Filter("age__gt", 7).One(&yswData)
	//qs.Filter("age__gt", 5).Filter("id__gt",35).One(&yswData, "id")

	qs.All(&yswArray)
	y.Data["data"] = yswArray
	fmt.Println(yswArray)

	y.TplName = "ysw/ysw.html"
}

type AddListYeController struct {
	beego.Controller
}


func (l *AddListYeController) Get()  {
	l.TplName="ysw/add.html"
}

type AddYeController struct {
	beego.Controller
}

func (a *AddYeController) Post()  {
	o := orm.NewOrm()
	age := a.GetString("age")
	name := a.GetString("name")
	ysw := models.Ysw{Age:age, Name:name}
	id, err := o.Insert(&ysw)
	fmt.Println(id)
	fmt.Println(err)
	a.TplName="success.html"
}

type DelController struct {
	beego.Controller
}

func (d *DelController) Get()  {
	o :=orm.NewOrm()
	idString := d.GetString("id")

	idInt, _ := strconv.ParseInt(idString, 10, 64)
	id := *(*int)(unsafe.Pointer(&idInt))

	delInfo := models.Ysw{Id:id}
	num, err := o.Delete(&delInfo)
	fmt.Println(err)
	fmt.Println(num)
	if num == 0{
		//d.Ctx.WriteString("删除失败")
		d.TplName = "error.html"
	} else {
		d.TplName = "success.html"
	}

}