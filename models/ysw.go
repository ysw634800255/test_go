package models

import "github.com/astaxie/beego/orm"

type Ysw struct {
	Id int
	Age string
	Name string
}

func init() {
	orm.RegisterModel(new(Ysw))
}