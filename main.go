package main

import (
	"flag"
	"fmt"

	"github.com/atotto/clipboard"
)

func main() {
	// CLI flags
	length := flag.Int("length", 12, "Length of the password")
	useUpper := flag.Bool("uppercase", false, "Include uppercase letters")
	useNumbers := flag.Bool("numbers", false, "Include numbers")
	useSpecial := flag.Bool("special", false, "Include special characters")
	copyToClipboard := flag.Bool("copy", false, "Copy password to clipboard")

	flag.Parse()

	password := GeneratePassword(*length, *useUpper, *useNumbers, *useSpecial)

	fmt.Printf("🔐 Generated Password: %s\n", password)

	if *copyToClipboard {
		err := clipboard.WriteAll(password)
		if err != nil {
			fmt.Println("❌ Failed to copy to clipboard:", err)
		} else {
			fmt.Println("📋 Password copied to clipboard!")
		}
	}
}
