package WebClientHelper

import (
	logHelper "GoMVC/utils/LogHelper"

	"github.com/astaxie/beego/httplib"
)

const connectTimeout = 120
const readWriteTimeout = 120

func GetResponse(url string, header map[string]string) string {
	req := httplib.Get(url)
	req.SetTimeout(connectTimeout, readWriteTimeout)
	if header != nil {
		for key, val := range header {
			req.Header(key, val)
		}
	}

	str, err := req.String()
	if err != nil {
		logHelper.LogError_e(err)
	}
	return str

}
func GetResponseByte(url string, header map[string]string) []byte {
	req := httplib.Get(url)
	req.SetTimeout(connectTimeout, readWriteTimeout)
	if header != nil {
		for key, val := range header {
			req.Header(key, val)
		}
	}

	bt, err := req.Bytes()
	if err != nil {
		logHelper.LogError_e(err)
	}
	return bt

}
func PostResponse(url string, header map[string]string) string {
	req := httplib.Post(url)
	req.SetTimeout(connectTimeout, readWriteTimeout)
	if header != nil {
		for key, val := range header {
			req.Header(key, val)
		}
	}

	str, err := req.String()
	if err != nil {
		logHelper.LogError_e(err)
	}
	return str

}
func PostResponseByte(url string, header map[string]string) []byte {
	req := httplib.Post(url)
	req.SetTimeout(connectTimeout, readWriteTimeout)
	if header != nil {
		for key, val := range header {
			req.Header(key, val)
		}
	}

	bt, err := req.Bytes()
	if err != nil {
		logHelper.LogError_e(err)
	}
	return bt

}
