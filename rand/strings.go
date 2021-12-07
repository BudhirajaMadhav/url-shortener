package rand

import (
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	byteSlice := make([]byte, length)

	for index := range byteSlice {
		byteSlice[index] = charset[seededRand.Intn(len(charset))]
	}

	return string(byteSlice)

}

func String(length int) string {
	charset := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	return StringWithCharset(length, charset)
}