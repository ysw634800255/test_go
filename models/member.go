package models

import "github.com/astaxie/beego/orm"

type Member struct {
	Id 			int		`orm:"pk"`
	Email 		string
	Name 		string
	Sex	 		int
	Tel 		string
	Addres		string
	Password	string
	Add_time 	int
	Status 		int

}

func init() {
	orm.RegisterModel(new(Member))
}