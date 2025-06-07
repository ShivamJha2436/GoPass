package main

import (
	"math/rand"
	"strings"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:',.<>?/"
)

func GeneratePassword(length int, useUpper, useSpecial, useNumbers bool) string {
	charset := lowerChars
	if useUpper {
		charset += upperChars
	}
	if useNumbers {
		charset += numberChars
	}
	if useSpecial {
		charset += specialChars
	}

	if len(charset) == 0 {
		return ""
	}

	var password strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		password.WriteByte(charset[randomIndex])
	}
	return password.String()
}
