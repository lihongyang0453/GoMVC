package logHelper

import (
	"encoding/json"

	logconf "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

func LogError(msg string) {
	logs.Error(msg + "\n")
}
func LogInfo(msg string) {
	logs.Info(msg + "\n")
}

func LogWarning(msg string) {
	logs.Warning(msg + "\n")
}
func LogDebug(msg string) {
	logs.Debug(msg + "\n")
}

func init() {
	conf, err1 := logconf.NewConfig("ini", "conf/log.conf")
	if err1 != nil {
		return
	}
	config := make(map[string]interface{})
	// 	filename 保存的文件名
	// maxlines 每个文件保存的最大行数，默认值 1000000
	// maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	// daily 是否按照每天 logrotate，默认是 true
	// maxdays 文件最多保存多少天，默认保存 7 天
	// rotate 是否开启 logrotate，默认是 true
	// level 日志保存的时候的级别，默认是 Trace 级别 logs.LevelDebug
	// perm 日志文件权限
	config["filename"] = conf.String("log_path")
	config["level"] = conf.DefaultInt(conf.String("log_level"), 6)

	configStr, err := json.Marshal(config)
	if err != nil {
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.SetLogFuncCall(true)
	logs.EnableFuncCallDepth(true) //输出调用的文件名和文件行号
	return

}

//发送电子邮件
func LogEmail(data string) {
	conf, err1 := logconf.NewConfig("ini", "conf/app.conf")
	if err1 != nil {
		return
	}
	log := logs.NewLogger()
	config_smtp := make(map[string]interface{})
	config_smtp["username"] = conf.String("smtp_username")
	config_smtp["password"] = conf.String("smtp_password")
	config_smtp["host"] = conf.String("smtp_host")
	config_smtp["fromAddress"] = conf.String("smtp_fromAddress")
	config_smtp["sendTos"] = []string{conf.String("smtp_sendTos")}
	configSmtp, errsmtp := json.Marshal(config_smtp)
	if errsmtp != nil {
		return
	}
	//str := `{"username":"lihongyang0453@126.com","password":"li8978591","host":"smtp.126.com:25","fromAddress":"lihongyang0453@126.com","sendTos":["lihy@crealifemed.com"]}`
	log.SetLogger(logs.AdapterMail, string(configSmtp))
	//log.SetLogFuncCall(true)
	log.EnableFuncCallDepth(false) //输出调用的文件名和文件行号
	log.Info(data)
}

///给传入的地址发送邮件
func LogEmailWithAddress(data string, receiver []string) {
	conf, err1 := logconf.NewConfig("ini", "conf/app.conf")
	if err1 != nil {
		return
	}
	if len(receiver) <= 0 { //没传给个默认接受人
		receiver = []string{conf.String("smtp_sendTos")}
	}
	log := logs.NewLogger()
	config_smtp := make(map[string]interface{})
	config_smtp["username"] = conf.String("smtp_username")
	config_smtp["password"] = conf.String("smtp_password")
	config_smtp["host"] = conf.String("smtp_host")
	config_smtp["fromAddress"] = conf.String("smtp_fromAddress")
	config_smtp["sendTos"] = receiver
	configSmtp, errsmtp := json.Marshal(config_smtp)
	if errsmtp != nil {
		return
	}
	//str := `{"username":"lihongyang0453@126.com","password":"li8978591","host":"smtp.126.com:25","fromAddress":"lihongyang0453@126.com","sendTos":["lihy@crealifemed.com"]}`
	log.SetLogger(logs.AdapterMail, string(configSmtp))
	log.EnableFuncCallDepth(false) //输出调用的文件名和文件行号
	log.Info(data)
}
