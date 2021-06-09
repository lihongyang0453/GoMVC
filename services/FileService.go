package services

import "os"

type FileService struct {
	BaseService
}

//判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

///创建文件夹
func MkDir(folderPath string) {
	os.Mkdir(folderPath, os.ModePerm)
}
