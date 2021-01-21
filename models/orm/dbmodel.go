package dbmodel

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	// 需要在 init 中注册定义的 model
	orm.RegisterModel(new(User), new(Profile))
}

type User struct {
	Id      int `orm:"column(Id)"`
	Name    string
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

func (u *User) TableName() string {
	return "g_user"
}

type Profile struct {
	Id   int   `orm:"column(Id)"`
	Age  int16 `orm:"column(Age)"`
	User *User `orm:"reverse(one)"` // 设置反向关系（可选）
}

func (p *Profile) TableName() string {
	return "g_profile"
}
