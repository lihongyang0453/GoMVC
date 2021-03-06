package OrmModel

import (
	"time"
)

type User struct {
	Id            int       `orm:"column(Id);pk;auto"`
	CreatedTime   time.Time `orm:"column(CreatedTime)"`
	DisplayName   string    `orm:"column(DisplayName)"`
	UserName      string    `orm:"column(UserName)"`
	Pwd           string    `orm:"column(Pwd)"`
	Mobile        string    `orm:"column(Mobile)"`
	DateLastVisit time.Time `orm:"column(DateLastVisit)"`
	IsDelete      int       `orm:"column(IsDelete);default(0)"`
}

// Profile *Profile `orm:"rel(one)"` // OneToOne relation 创建对应关系的字段:表_id,即对应主键,有唯一约束
func (u *User) TableName() string {
	return "t_userinfo"
}
