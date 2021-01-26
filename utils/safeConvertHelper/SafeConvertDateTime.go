package safeConvertHelper

import "time"

type goDate struct {
	Time           time.Time
	TimeZone       *time.Location
	FirstDayOfWeek time.Weekday
}

const layOut = "2006-01-02 15:04:05"
const defaultDate = "2010-01-01 00:00:00"

//转换时间
// func GetSafeDateTime(o interface{}) time.Time {
// 	reslut, err := time.Parse(layOut, o)
// 	if err != nil {
// 		defaultResult, _ := time.Parse(layOut, defaultDate)
// 		return defaultResult
// 	}
// 	return reslut
// }

func (d *goDate) IsBefore(compare *goDate) bool {
	return d.Time.Before(compare.Time)
}
