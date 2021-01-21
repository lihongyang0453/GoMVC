package model

type BaseResponse struct {
	// ID         uint      `json:"id"`
	// Username   string    `json:"username"`
	// Password   string    `json:"-"`
	// CreateTime time.Time `json:"create_time"`
	ErrCode int
	Msg     string
}
type BasePageItem struct {
	Total uint
}
