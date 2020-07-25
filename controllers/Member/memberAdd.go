package Member

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
	"unsafe"
	"ysw/models"
)

type MemberAddController struct {
	beego.Controller
}

func (m *MemberAddController) Get()  {
	c, err := redis.Dial("tcp", "localhost:6379", redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	result,err := redis.Values(c.Do("hgetall","menber:13"))//读

	for _, v := range result {
		fmt.Printf("%s ", v.([]byte))
	}
	fmt.Println(result)
	//_, _ = c.Do("expire", "menber:12", 10000)
	m.TplName = "Member/member_add.html"
}

type MemberAdDataController struct {
	beego.Controller
}

func (m *MemberAdDataController) Post()  {
	o := orm.NewOrm()
	c, err := redis.Dial("tcp", "localhost:6379", redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	name := m.GetString("name")
	email := m.GetString("email")
	tel := m.GetString("tel")
	addres := m.GetString("addres")
	Password := fmt.Sprintf("%x", md5.Sum([]byte(m.GetString("password"))))

	//类型转换
	sexString := m.GetString("sex")
	idInt, _ := strconv.ParseInt(sexString, 10, 64)
	sex := *(*int)(unsafe.Pointer(&idInt))

	//类型转换
	timeInt     := time.Now().Unix()
	addTime := *(*int)(unsafe.Pointer(&timeInt))

	addValue := models.Member{Name:name, Email:email, Addres:addres, Add_time: addTime, Tel:tel, Sex:sex, Password:Password, Status:0}
	id, err := o.Insert(&addValue)
	if err != nil {
		m.TplName = "error.html"
	}

	//插入redis,menber:1 先将id转为字符串，再拼接member
	member := "menber:" + strconv.FormatInt(id,10)
	_, err = c.Do("HMSet", member,
		"name", name, "email", email, "email", addres, "addres", addres, "addTime", addTime, "tel", tel, "Password", Password, "Status", 0)
	//设定过期时间
	_, _ = c.Do("expire", member, 100)

	//读取redis
	res, err := c.Do("hGetAll", "menber:10")
	fmt.Println(res)
	m.TplName = "success.html"
}

