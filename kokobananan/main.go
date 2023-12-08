package main

import (
	"fmt"
	"math"
)

func canKOKOFinish(numbers []int, speed int, hours int) bool {
	hrs := 0

	for i := 0; i < len(numbers); i++ {
		hoursss := int(math.Ceil(float64(numbers[i]) / float64(speed)))
		hrs = hrs + hoursss
		if hrs > hours {
			return false
		}
	}
	return true
}

func returnMaxOfArr(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, returnMaxOfArr(piles)
	for left < right {
		mid := int(((left + right) / 2))
		if canKOKOFinish(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		fmt.Println(nums[mid])
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // Indicates that the target is not found
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := int(math.Ceil((float64(left) + float64(right)) / 2))
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func strStr(haystack string, needle string) int {
	lenth := len(haystack)
	needlength := len(needle)
	if lenth < needlength {
		return -1
	} else if lenth == needlength && haystack == needle {
		return 0
	}
	for i := 0; i <= lenth-needlength; i++ {
		if haystack[i:i+needlength] == needle {
			return i
		}
	}
	return -1
}

func plusOne(digits []int) []int {
	indexArr := []int{1}
	for i := 10; i < 100000; i = i * 10 {
		indexArr = append(indexArr, i)
	}
	mainNum := 0
	for i := len(digits) - 1; i >= 0; i-- {
		val := digits[i] * indexArr[len(digits)-i-1]
		mainNum = mainNum + val
	}
	secondaryNum := mainNum + 1
	secondaryNumArr := []int{}
	for secondaryNum > 0 {
		secondaryNumArr = append(secondaryNumArr, (secondaryNum % 10))
		secondaryNum = secondaryNum / 10
	}
	for i, j := 0, len(secondaryNumArr)-1; i < j; i, j = i+1, j-1 {
		secondaryNumArr[i], secondaryNumArr[j] = secondaryNumArr[j], secondaryNumArr[i]
	}
	fmt.Println(secondaryNumArr)
	return secondaryNumArr
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))
}
