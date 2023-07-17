package ego

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func DateNowFormatA() string {
	return time.Now().Format("2006/1/2")
}

func DateNowFormatB() string {
	return time.Now().Format("2006/01/02")
}

func DateNowFormatC() string {
	return time.Now().Format("20060102150405")
}

func DateNowFormatD(date time.Time) string {
	return date.Format("2006-01-02")
}

func DateNowFormatE(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func DateNowFormatF(date time.Time) string {
	return date.Format("200601021504")
}

// TimestampToFormat @description: 10位时间戳转换为2006-01-02 15:04:05格式时间字符串
// @parameter timeUnix(例1680001686)
// @return string
func TimestampToFormat(timeUnix int64) string {
	return time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
}

// TimeStrToFormat @description: 2006-01-02T15:04:05+08:00转换为2006-01-02 15:04:05
// @parameter timeStr
// @return string
func TimeStrToFormat(timeStr string) string {
	t, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", timeStr, time.Local)
	if nil != err {
		log.Println("TimeStrToFormat(),err:", err)
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

// TimeStrToFormat2 @description:2006-01-02 15:04:05转换为20060102150405
// @parameter timeStr
// @return string
func TimeStrToFormat2(timeStr string) string {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if nil != err {
		log.Println("TimeStrToFormat2(),err:", err)
		return ""
	}
	return t.Format("20060102150405")
}

// StrToTimeA @description: 2006/1/2 15:04:05转换为time.Time
// @parameter timeStr
// @return time.Time
func StrToTimeA(timeStr string) time.Time {
	t, err := time.ParseInLocation("2006/1/2 15:04:05", timeStr, time.Local)
	if nil != err {
		log.Println("StrToTimeA(),err:", err)
		return time.Now()
	}
	return t
}

// StrToTimeB @description: 2006-01-02 15:04:05转为time.Time
// @parameter timeStr
// @return time.Time
func StrToTimeB(timeStr string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if nil != err {
		log.Println("StrToTimeA(),err:", err)
		return time.Now()
	}
	return t
}

// IsToday @description: 判断时间点处于今天
// @parameter d
// @return bool
func IsToday(d time.Time) bool {
	now := time.Now()
	return d.Year() == now.Year() && d.Month() == now.Month() && d.Day() == now.Day()
}
