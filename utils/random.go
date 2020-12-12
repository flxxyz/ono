package utils

import (
	"math/rand"
	"time"
)

var (
	letters      = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lettersUpper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandomNumber(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func RandomSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[RandomNumber(0, len(letters))]
	}
	return string(b)
}

func RandSeqUpper(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersUpper[RandomNumber(0, len(lettersUpper))]
	}
	return string(b)
}
