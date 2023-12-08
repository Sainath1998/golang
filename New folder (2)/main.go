// all permutations of array of integer
package main

import "fmt"

func permuteStrings(data []string) [][]string {
	result := [][]string{}
	var backTrack func(data []string, index int)
	backTrack = func(data []string, index int) {
		if len(data) == index {
			temp := make([]string, len(data))
			copy(temp, data)
			result = append(result, temp)
		}
		for i := index; i < len(data); i++ {
			data[i], data[index] = data[index], data[i]
			backTrack(data, index+1)
			data[index], data[i] = data[i], data[index]
		}
	}
	backTrack(data, 0)
	return result
}

func main() {
	// set := []int{1, 2, 3, 4}
	// finalResult := permute(set)
	// fmt.Println(finalResult)
	stringResult := []string{"bathing", "suit", "strtch"}
	finalResult := permuteStrings(stringResult)
	fmt.Println(finalResult)
}
