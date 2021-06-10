package controllers

import(
	model "GoMVC/models/BaseModel"
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
func (r *RoleController)SaveModel(data interface{}) *model.BaseResponse  {
 	var result *model.BaseResponse = new(model.BaseResponse)

	result.ErrCode=1
	result.Msg="error"
	return result
}