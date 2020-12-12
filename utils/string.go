package utils

import "strconv"

// UpperFirstChar 首字母大写
func UpperFirstChar(str string) string {
	if len(str) > 0 {
		runes := []rune(str)
		if int(runes[0]) >= 97 && int(runes[0]) <= 122 {
			runes[0] = rune(int(runes[0]) - 32)
			str = string(runes)
		}
	}
	return str
}

// Int642Str
func Int642Str(n int64) string {
	return strconv.FormatInt(n, 10)
}

// Str2Int64
func Str2Int64(s string) (i int64) {
	i, _ = strconv.ParseInt(s, 10, 8)
	return
}
