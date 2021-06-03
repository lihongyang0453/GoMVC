package controllers

import(

	"GoMVC/model"
)

type RoleController struct{
	BaseController
}

func (r *RoleController)getByID(id int32) string{
	return ""
}
func (r *RoleController)getByName(username string) string{
	return ""
}
func (r *RoleController)SaveModel(data interface{}) model.BaseResponse  {
	result:=make(model.BaseResponse)

	return result
}