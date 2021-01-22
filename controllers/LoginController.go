package controllers

type LoginController struct {
	BaseController
}

// func (c *LoginController) URLMapping() {
// 	c.Mapping("StaticBlock", c.StaticBlock)
// }

// @router /staticblock/:key [get]
func (this *LoginController) StaticBlock() {
	// orm.Debug = true
	// // 自动建表
	// orm.RunSyncdb("default", false, true)
	// o := orm.NewOrm()
	// o.Using("default")
	// perfile := new(dbmodel.Profile)
	// perfile.Age = 30

	// user := new(dbmodel.User)
	// user.Name = "tom"
	// user.Profile = perfile
	// o.Insert(perfile)
	// o.Insert(user)

	// user.Name = "hezhixiong"
	// num, err := o.Update(user)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// o.Delete(&dbmodel.User{Id: 2})

	//this.ServeJSON()
}
