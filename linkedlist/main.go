package main

import (
	"fmt"
	"strconv"
	"strings"
)

func facto(n int, less int) int {
	if n == less {
		return 0
	}
	return n + facto(n-1, less)
}

func totalMoney(n int) int {
	arr := []int{}
	for n >= 7 {
		arr = append(arr, 7)
		n = n - 7
	}
	arr = append(arr, n)
	fmt.Println(arr)
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + facto(arr[i]+i, i)
	}
	return sum
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	c := strings.Join(word1, "")
	d := strings.Join(word2, "")
	fmt.Println(c)
	fmt.Println(d)
	return c == d
}

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		intval, _ := strconv.Atoi(string(num[i]))
		if intval%2 != 0 {
			a := num[:i+1]
			return a
		}
	}
	return ""
}

func main() {
	// fmt.Println(facto(10))
	// fmt.Println()
	// s := totalMoney(10)
	// fmt.Println(s)
	// fmt.Println(facto(7))
	// word1 := []string{"ab", "c"}
	// word2 := []string{"a", "bc"}
	z := "239537672423884969653287101"

	c := largestOddNumber(z)
	fmt.Println(c)

}
