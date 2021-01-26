package OrmModel

type Profile struct {
	Id  int   `orm:"column(Id)"`
	Age int16 `orm:"column(Age)"`
	//User *User `orm:"reverse(one)"` // 设置反向关系（可选）不会创建字段,可选tag
}

func (p *Profile) TableName() string {
	return "g_profile"
}
