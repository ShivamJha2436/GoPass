package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	length := flag.Int("length", 12, "Length of the password")
	useSpecial := flag.Bool("special", false, "Include special characters")
	useUpper := flag.Bool("uppercase", false, "Include uppercase letters")
	useNumbers := flag.Bool("numbers", false, "Include numbers")

	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	// ✅ This function must exist in generator.go
	password := GeneratePassword(*length, *useUpper, *useSpecial, *useNumbers)

	fmt.Println("🔐 Generated Password:", password)
}
