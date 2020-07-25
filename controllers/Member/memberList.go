package Member

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"ysw/models"
)

type MemberListController struct {
	beego.Controller
}

func (m *MemberListController) Get()  {
	o := orm.NewOrm()

	star := m.GetString("start")
	end := m.GetString("end")
	name := m.GetString("name")

	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	timeTemplate1 := "2006-01-02 15:04:05" //常规类型endTime

	//开始时间
	starTime, _ := time.ParseInLocation(timeTemplate1, star + " 00:00:00", time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	endTime, _ := time.ParseInLocation(timeTemplate1, end + " 00:00:00", time.Local) //使用parseInLocation将字符串格式化返回本地时区时间

	list := []models.Member{}
	menber := new(models.Member)
	qs := o.QueryTable(menber)
	nil := ""
	//开始时间为空，则不执行时间条件
	if star != nil {
		qs = qs.Filter("add_time__gte", starTime.Unix())

		//如果结束时间为空则定位到当前时间
		if end != nil {
			qs = qs.Filter("add_time__lte", endTime.Unix())
		} else {
			qs = qs.Filter("add_time__lte", time.Now())
		}
	}

	//名称条件判断
	if name != nil {
		qs = qs.Filter("name__contains", name)
	}
	_, _ = qs.OrderBy("-id").All(&list)
	m.Data["data"] = list
	m.Data["search_start"] = star
	m.Data["search_end"] = end
	m.Data["search_name"] = name
	m.TplName = "Member/member_list.html"
}
