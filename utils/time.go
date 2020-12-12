package utils

import (
	"strconv"
	"time"
)

// GetCurrTimestamp 返回当前时间戳(数值)
func GetCurrTimestamp() int64 {
	return time.Now().Unix()
}

// CurrTsString 返回当前时间戳(字符)
func CurrTsString() string {
	return strconv.FormatInt(GetCurrTimestamp(), 10)
}
