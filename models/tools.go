package models

import (
	"time"
)

// 把时间戳换算成日期
func UnixToTime(timesmap int) string {
	t := time.Unix(int64(timesmap), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳
func DateToTime(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	} else {
		return t.Unix()
	}
}

// 获取Unix当前时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取当前日期
func GetDate() string {
	template := "2002-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}
