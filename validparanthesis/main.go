package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)
	stringMap := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, char := range s {
		if char == '{' || char == '(' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != stringMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0

}

func main() {
	fmt.Println("Hello world")
	fmt.Println(isValid("([])"))
}
