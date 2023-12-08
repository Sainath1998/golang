package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func permute(nums []int) [][]int {
	ans := [][]int{}
	var backTrack func(nums []int, index int)
	backTrack = func(nums []int, index int) {
		if len(nums) == index {
			temp := make([]int, len(nums))
			copy(temp, nums)
			ans = append(ans, temp)
		}
		for i := index; i < len(nums); i++ {
			nums[i], nums[index] = nums[index], nums[i]
			backTrack(nums, index+1)
			nums[index], nums[i] = nums[i], nums[index]
		}
	}
	backTrack(nums, 0)
	return ans
}

func permuteStrings(s1string string, target string) []string {
	chars := strings.Split(s1string, "")
	ans := []string{}
	var backTrack func(chars []string, index int)
	backTrack = func(chars []string, index int) {
		if index == len(chars) {
			temp := make([]string, len(chars))
			copy(temp, chars)
			fmt.Println(strings.Join(temp, ""))
			fmt.Println(target)
			if strings.Join(temp, "") == target {
				ans = append(ans, strings.Join(temp, ""))
			}
			// ans = append(ans, strings.Join(temp, ""))
		}
		for i := index; i < len(chars); i++ {
			chars[i], chars[index] = chars[index], chars[i]
			backTrack(chars, index+1)
			chars[index], chars[i] = chars[i], chars[index]
		}
	}
	backTrack(chars, 0)
	return ans
}

func charsInclude(char1arr string, char2arr string) bool {
	runes1 := []rune(char1arr)
	runes2 := []rune(char2arr)
	sort.Slice(runes1, func(i, j int) bool {
		return runes1[i] < runes1[j]
	})
	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] < runes2[j]
	})
	return string(runes1) == string(runes2)
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// Initialize character frequency maps
	s1Freq := make(map[byte]int)
	s2Freq := make(map[byte]int)

	// Initialize frequency maps for the first window of size len(s1)
	for i := 0; i < len(s1); i++ {
		s1Freq[s1[i]]++
		s2Freq[s2[i]]++
	}

	// Check if the frequency maps match for the first window
	if reflect.DeepEqual(s1Freq, s2Freq) {
		return true
	}

	// Slide the window through the rest of s2
	for i := len(s1); i < len(s2); i++ {
		// Update s2Freq for the new window
		s2Freq[s2[i]]++
		s2Freq[s2[i-len(s1)]]--

		// Check if the frequency maps match for the current window
		if reflect.DeepEqual(s1Freq, s2Freq) {
			return true
		}
	}

	return false
}

func main() {
	s2 := "cdef"
	s1 := "ef"
	result := checkInclusion(s1, s2)
	fmt.Println(result)
}
