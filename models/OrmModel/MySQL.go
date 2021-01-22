package OrmModel

import (
	"strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username   = "root"
	password   = "!Aa123123"
	ip         = "localhost"
	port       = "3306"
	dbName     = "MemberMSDB"
	driverName = "mysql"
)

func init() {

	// username := beego.AppConfig.String("mysql_username")
	// password := beego.AppConfig.String("mysql_password")
	// ip := beego.AppConfig.String("mysql_ip")
	// port := beego.AppConfig.String("mysql_port")
	// dbName := beego.AppConfig.String("mysql_dbName")
	// driverName := beego.AppConfig.String("mysql_driverName")

	mysqlConnectionStr := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	orm.RegisterDriver(driverName, orm.DRMySQL)
	orm.RegisterDataBase("default", driverName, mysqlConnectionStr)

	// 需要在 init 中注册定义的 model
	orm.RegisterModel(new(User), new(Profile))
}
