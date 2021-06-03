package services

type BaseService interface {
	Delete(a interface{}) bool
	DeleteById(ID int) bool
	GetModelById(ID int) interface{}
}
