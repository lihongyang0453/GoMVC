package OrmModel

type User struct {
	Id      int      `orm:"column(Id)"`
	Name    string   `orm:"column(Name)"`
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

func (u *User) TableName() string {
	return "g_user"
}
