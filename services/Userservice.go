package services

import (
	ormModel "GoMVC/models/OrmModel"
	logHelper "GoMVC/utils/LogHelper"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type UserService struct {
	User *ormModel.User
}

func (b *UserService) Insert(model ormModel.User) bool {
	o := orm.NewOrm()
	err := o.Begin() // 事务处理过程
	_, ierror := o.Insert(&model)
	if ierror != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	return err == nil
}

func (b *UserService) Delete(model ormModel.User) bool {
	o := orm.NewOrm()

	err := o.Begin() // 事务处理过程
	_, ierror := o.Delete(&model)
	if ierror != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	return err == nil
}
func (b *UserService) DeleteById(id int) bool {
	o := orm.NewOrm()
	err := o.Begin() // 事务处理过程
	model := ormModel.User{}
	model.Id = id
	_, ierror := o.Delete(&model)
	if ierror != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	return err == nil
}
func (b *UserService) FindById(id int) ormModel.User {
	o := orm.NewOrm()
	model := ormModel.User{}
	model.Id = id
	ierror := o.Read(&model)
	if ierror != nil {
		logHelper.LogError_e(ierror)
	}
	return model
}

func (b *UserService) Update(model ormModel.User) bool {
	o := orm.NewOrm()
	err := o.Begin() // 事务处理过程
	_, ierror := o.Update(&model)
	if ierror != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	return err == nil
}

//返回DataTable
func (b *UserService) QueryAll() orm.QuerySeter {
	return orm.NewOrm().QueryTable(b.User)
}

func (b *UserService) QueryListPaged(pageIndex int, pageSize int, filter map[string]interface{}, orderBy string, totalCount *int) []ormModel.User {
	var user []ormModel.User

	sql := "select * from t_userInfo where 1=1 "

	o := orm.NewOrm()
	num, err := o.Raw(sql, 0).QueryRows(&user)
	if err != nil {
		logHelper.LogError(strconv.FormatInt(num, 10))
	}
	return user
}

func (b *UserService) GetList(pageIndex int, pageSize int, filter map[string]interface{}, orderBy string, totalCount *int) []ormModel.User {
	var user []ormModel.User
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
	num, err := o.Raw(sql, 0).QueryRows(&user)
	if err != nil {
		logHelper.LogError(strconv.FormatInt(num, 10))
	}
	return user
}
