package userinfo

import (
	"github.com/astaxie/beego/orm"
)

type UserInfo struct {
	id   int
	name string
	age  int
}

func init() {
	// 需要在 init 中注册定义的 model
	orm.RegisterModel(new(User))
}

func AddUser(name string, age int) {

	o := orm.NewOrm()
	o.Using("default")

	userinfo := new(userinfo.UserInfo)
	userinfo.name = name
	userinfo.age = age

	o.Insert(userinfo)
}
