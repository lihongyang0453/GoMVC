package OrmModel

import (
	"strings"

	appconf "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	conf, err1 := appconf.NewConfig("ini", "conf/db.conf")
	if err1 != nil {
		return
	}
	username := conf.String("mysql_username")
	password := conf.String("mysql_password")
	ip := conf.String("mysql_ip")
	port := conf.String("mysql_port")
	dbName := conf.String("mysql_dbName")
	driverName := conf.String("mysql_driverName")

	mysqlConnectionStr := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	orm.RegisterDriver(driverName, orm.DRMySQL)
	orm.RegisterDataBase("default", driverName, mysqlConnectionStr)

	// 需要在 init 中注册定义的 model
	orm.RegisterModel(new(User), new(Profile))
}
