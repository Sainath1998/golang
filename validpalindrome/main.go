package main

import (
	"fmt"
	"strings"
	"unicode"
)

func removeAlphanumeric(s string) string {
	var result strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			result.WriteRune(char)
		}
	}
	return strings.ToLower(result.String())
}

func isPalindrome(s string) bool {
	s = removeAlphanumeric(s)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}

func main() {
	a := "A man, a plan, a canal: Panama"
	fmt.Println(removeAlphanumeric(a))
}
