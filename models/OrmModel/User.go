package OrmModel

import (
	"time"

	"github.com/astaxie/beego/orm"
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

func (b *User) Insert() bool {
	o := orm.NewOrm()

	err := o.Begin() // 事务处理过程
	_, ierror := o.Insert(b)
	if ierror == nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	return err != nil
}

func (b *User) Delete() bool {
	o := orm.NewOrm()

	err := o.Begin() // 事务处理过程
	_, ierror := o.Delete(b)
	if ierror == nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	return err != nil
}

func (b *User) Update(fields ...string) bool {

	o := orm.NewOrm()

	err := o.Begin() // 事务处理过程
	_, ierror := o.Update(b, fields...)
	if ierror == nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	return err != nil
}

//返回DataTable
func (b *User) QueryAll() orm.QuerySeter {
	return orm.NewOrm().QueryTable(b)
}

// func (b *User) QueryListPaged(pageIndex int, pageSize int, filter map[string]interface{}, orderBy string, totalCount *int) orm.QuerySeter {
// 	user := new([]User)
// 	totalCount = orm.NewOrm().QueryTable(b).Filter(filter).Count()

// 	if len(orderBy) == 0 {
// 		return orm.NewOrm().QueryTable(b).Filter(filter).Limit(pageSize, (pageIndex-1)*pageSize).All(user)
// 	} else {
// 		return orm.NewOrm().QueryTable(b).Filter(filter).OrderBy(orderBy).Limit(pageSize, (pageIndex-1)*pageSize).All(user)
// 	}
// }

func (b *User) GetList(pageIndex int, pageSize int, filter map[string]interface{}, orderBy string, totalCount *int) []User {
	user := new([]User)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("t_userInfo").
		//InnerJoin("profile").On("user.id_user = profile.fk_user").
		Where("isDelete = ?").
		OrderBy("CreatedTime").Desc().
		Limit(pageSize).Offset((pageIndex - 1) * pageSize)
	// 导出 SQL 语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql, 0).QueryRows(&user)
	return *user
}
