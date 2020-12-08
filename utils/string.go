package utils

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
