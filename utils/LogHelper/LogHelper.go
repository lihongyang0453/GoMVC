package logHelper

import (
	"encoding/json"

	logconf "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

func LogError(msg string) {
	logs.Error(msg)
}
func LogInfo(msg string) {
	logs.Info(msg)
}

func LogWarning(msg string) {
	logs.Warning(msg)
}
func LogDebug(msg string) {
	logs.Debug(msg)
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
