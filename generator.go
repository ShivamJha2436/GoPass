package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}<>?/|"
)

// GeneratePassword returns a random password of the given length using selected character sets.
func GeneratePassword(length int, useUpper, useNumbers, useSpecial bool) string {
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
		fmt.Println("❌ Error: No character sets selected")
		os.Exit(1)
	}

	var password strings.Builder
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password.WriteByte(charset[index.Int64()])
	}

	return password.String()
}
