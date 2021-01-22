package Services

import (
	model "GoMVC/models/BaseModel"
	"GoMVC/models/OrmModel"
	"fmt"

	"github.com/astaxie/beego/orm"
)

func Login(name string, pwd string) model.BaseResponse {
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
	o := orm.NewOrm()
	o.Using("default")
	perfile := new(OrmModel.Profile)
	perfile.Age = 30

	user := new(OrmModel.User)
	user.Name = "tom"
	user.Profile = perfile
	o.Insert(perfile)
	o.Insert(user)

	user.Name = "hezhixiong"
	num, err := o.Update(user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	o.Delete(&OrmModel.User{Id: 2})
}
