package controllers

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	c.Data["Website"] = "bbbbbbbbbbbbbeego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	c.TplName = "index.tpl"
	c.Render()
}

func (this *UserController) Get2() {
	// v := this.GetSession("asta")
	// if v == nil {
	// 	this.SetSession("asta", int(1))
	//    this.Data["num"] = 0
	// } else {
	// 	this.SetSession("asta", v.(int)+1)
	// 	this.Data["num"] = v.(int)
	// }
	//c.TplName = "index.tpl"

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
}
