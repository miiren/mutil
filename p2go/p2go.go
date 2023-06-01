package p2go

import (
	"github.com/syyongx/php2go"
)

func Strlen(s string) int {
	return php2go.Strlen(s)
}

func Implode(glue string, pieces []string) string {
	return php2go.Implode(glue, pieces)
}

func Md5(str string) string {
	return php2go.Md5(str)
}

func Strtotime(format, strtime string) (int64, error) {
	return php2go.Strtotime(format, strtime)
}

func Date(format string, timestamp int64) string {
	return php2go.Date(format, timestamp)
}

// Sleep sleep()
func Sleep(t int64) {
	php2go.Sleep(t)
}

// Usleep usleep()
func Usleep(t int64) {
	php2go.Usleep(t)
}

func Strpos(haystack, needle string, offset int) int {
	return php2go.Strpos(haystack, needle, offset)
}
