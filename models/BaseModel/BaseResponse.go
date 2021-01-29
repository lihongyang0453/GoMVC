package model

type BaseResponse struct {
	ErrCode int
	Msg     string
}
type BasePageItem struct {
	Total uint
	Items interface{}
}
