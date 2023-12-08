package main

import (
	"fmt"
	"strings"
)

func checkIfCharExist(arr []string, char string) int {
	for index, ele := range arr {
		if ele == char {
			return index
		}
	}
	return -1
}

func lengthOfLongestSubstring(s string) int {
	arr := strings.Split(s, "")
	top := 0
	longestSofar := 0
	longestArr := []string{}
	for i := 0; i < len(arr); i++ {
		if len(longestArr) == 0 {
			longestArr = append(longestArr, arr[i])
			longestSofar++
			if longestSofar > top {
				top = longestSofar
			}
			continue
		}
		val := checkIfCharExist(longestArr, arr[i])
		if val == -1 {
			longestArr = append(longestArr, arr[i])
			longestSofar++
			if longestSofar > top {
				top = longestSofar
			}
		} else {
			longestArr = longestArr[val+1:]
			if longestSofar > top {
				top = longestSofar
			}
			longestArr = append(longestArr, arr[i])
			longestSofar = len(longestArr)
		}
	}
	return top
}

func checkInclusion(s1 string, s2 string) bool {
	lenght := len(s1)
	s3 := strings.Split(s2, "")
	for i := 0; i < len(s3); i++ {
		newArr := s3[i:]
		pattern := strings.Join(newArr[:lenght], "")
		fmt.Println(pattern)
		fmt.Println(s1)
		if pattern == s1 {
			return true
		}

	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}
