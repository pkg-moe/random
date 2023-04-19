package random

import (
	"math/rand"
	"strings"
	"time"
)

func NonceCustom(length int, chars string) string {
	rand.Seed(time.Now().UnixNano())
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(rune(chars[rand.Intn(len(chars))]))
	}
	return b.String()
}

func Nonce(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789",
	)

	return NonceCustom(length, string(chars))
}

func NonceLowercase(length int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyz" +
		"0123456789",
	)

	return NonceCustom(length, string(chars))
}
