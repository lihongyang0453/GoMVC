package controllers

import (
	service "GoMVC/services"
	logHelper "GoMVC/utils/LogHelper"
)

type FileController struct {
	BaseController
}

func (c *FileController) UploadFile() {
	logHelper.LogInfo("UploadFile start...")
	f, h, err := c.GetFile("uploadname")
	logHelper.LogInfo(h.Filename)
	if err != nil {
		logHelper.LogError("getfile err")
	}
	defer f.Close()
	folder := "/Users/baobaob/go/src/GoMVC/static/upload/"
	flag, _ := service.PathExists(folder)
	if !flag {
		service.MkDir(folder)
	}
	er := c.SaveToFile("uploadname", folder+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	if er != nil {
		logHelper.LogError_e(er)
		return
	}
	logHelper.LogInfo("UploadFile end...")
}
