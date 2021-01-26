package dateTimeHelper

import "time"

func GetCurrent() time.Time {
	return time.Now()
}

// 获取当前时间的时间戳  s
func GetCurrentTimeStamp() int64 {
	return time.Now().Unix()
}

// 获取当前时间的时间戳纳秒  ns
func GetCurrentTimeStamp_ns() int64 {
	return time.Now().UnixNano()
}

//时间戳转日期
func TimeStampToTime(stamp int64) time.Time {
	return time.Unix(int64(stamp), 0)
}
