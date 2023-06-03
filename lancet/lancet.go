package lancet

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/datetime"
	"time"
)

//这里只转写部分函数，用于工具记录。如果需使用更多函数，请直接使用github.com/duke-git/lancet
//参考文档 https://zhuanlan.zhihu.com/p/543391872
//https://github.com/duke-git/lancet/blob/main/README_zh-CN.md
//convertor文档 https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToInt

func ToInt(value interface{}) (int64, error) {
	return convertor.ToInt(value)
}

func ToString(value interface{}) string {
	return convertor.ToString(value)
}

func ToBytes(value interface{}) ([]byte, error) {
	return convertor.ToBytes(value)
}

func FormatTimeToStr(t time.Time, format string) string {
	return datetime.FormatTimeToStr(t, format)
}

// FormatStrToTime convert string to time
func FormatStrToTime(str, format string) (time.Time, error) {
	return datetime.FormatStrToTime(str, format)
}

// AddMinute add or sub minute to the time
func AddMinute(t time.Time, minute int64) time.Time {
	return datetime.AddMinute(t, minute)
}

// AddHour add or sub hour to the time
func AddHour(t time.Time, hour int64) time.Time {
	return datetime.AddHour(t, hour)
}

// AddDay add or sub day to the time
func AddDay(t time.Time, day int64) time.Time {
	return datetime.AddDay(t, day)
}

func HmacSha256(data, key string) string {
	return cryptor.HmacSha256(data, key)
}

func Md5(s string) string {
	return cryptor.Md5String(s)
}
