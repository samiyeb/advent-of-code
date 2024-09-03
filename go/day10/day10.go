package main

import (
	"fmt"
	"strconv"
)

func day10(input string) (string, int) {
	result := ""
	resultCount := 0

	i := 0
	for i < len(input) {
		curr := input[i]
		count := 0
		for i < len(input) && curr == input[i] {
			i++
			count++
		}
		result += string(strconv.Itoa(count)) + string(curr)
		resultCount += 2
	}

	return result, resultCount
}

func main() {
	input := "3113322113"
	inputCount := 12

	for i := 0; i < 50; i++ {
		temp, tempCount := day10(input)
		input = temp
		inputCount = tempCount
	}
	

	fmt.Println(inputCount)
}