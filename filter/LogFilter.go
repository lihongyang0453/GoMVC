package filter

import (
	"context"
)

var FilterLog = func(ctx *context.Context) {
	// url, _ := json.Marshal(ctx.Input.Data()["RouterPattern"])
	// params, _ := json.Marshal(ctx.Request.Form)
	// outputBytes, _ := json.Marshal(ctx.Input.Data()["json"])
	// divider := " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -"
	// topDivider := "┌" + divider
	// middleDivider := "├" + divider
	// bottomDivider := "└" + divider
	// outputStr := "\n" + topDivider + "\n│ 请求地址:" + string(url) + "\n" + middleDivider + "\n│ 请求参数:" + string(params) + "\n│ 返回数据:" + string(outputBytes) + "\n" + bottomDivider
	// logHelper.LogInfo(outputStr)
}

func init() {
	//beego.InsertFilter("/*", beego.FinishRouter, FilterLog, false)
}
