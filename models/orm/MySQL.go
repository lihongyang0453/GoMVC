package dbmodel

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//_ "github.com/go-sql-driver/mysql"

func init() {
	username := beego.AppConfig.String("mysql_username")
	password := beego.AppConfig.String("mysql_password")
	ip := beego.AppConfig.String("mysql_ip")
	port := beego.AppConfig.String("mysql_port")
	dbName := beego.AppConfig.String("mysql_dbName")
	driverName := beego.AppConfig.String("mysql_driverName")

	mysqlConnectionStr := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	orm.RegisterDriver(driverName, orm.DRMySQL)
	orm.RegisterDataBase("default", driverName, mysqlConnectionStr)
}
