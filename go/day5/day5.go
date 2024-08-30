package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readingFile(fileName string) []string {
	var strings []string

	if file, err := os.Open(fileName); err != nil {
		log.Fatal(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			strings = append(strings, scanner.Text())
		}
	}

	return strings

}

// Part 1 Code

func threeVowels (s string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	count := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(vowels); j++ {
			if string(s[i]) == vowels[j] {
				count += 1
			}
		}

		if count == 3 { break }
	}

	return count >= 3
}

func back2back(s string) bool {
	if len(s) < 2 { return false }

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] { return true }
	}

	return false
}

func isPure(s string) bool {
	if len(s) < 2 { return true }

	for i := 1; i < len(s); i++ {
		var current string = string(s[i-1]) + string(s[i])

		if current == "ab" || current == "cd" || current == "pq" || current == "xy" {
			return false
		}
	}

	return true
}

// Part 2 Code

func analysis (pairs []string) bool{
	if len(pairs) < 2 { return false }

	for i := 0; i < len(pairs); i++ {
		for j := i + 2; j < len(pairs); j++ {
			if pairs[i] == pairs[j] {
				return true
			}
		}
	}

	return false

}

func offset(s string) []string {
	var pairs []string

	for i := 2; i < len(s); i += 2 {
		pairs = append(pairs, string(s[i-1]) + string(s[i]))
	}

	return pairs

}

func original(s string) []string {
	var pairs []string

	for i := 1; i < len(s); i += 2 {
		pairs = append(pairs, string(s[i-1]) + string(s[i]))
	}

	return pairs
}

func havePair(s string) bool {
	if len(s) < 4 { return false }

	ult := []string{}
	pairs1 := original(s)
	pairs2 := offset(s)

	if len(s) % 2 != 0 {
		for i := 0; i < len(pairs1); i++ {
			ult = append(ult, pairs1[i])
			ult = append(ult, pairs2[i])
		}
	} else {
		for i := 0; i < len(pairs2); i++ {
			ult = append(ult, pairs1[i])
			ult = append(ult, pairs2[i])

			if i == len(pairs2) - 1 {
				ult = append(ult, pairs1[i+1])
			}
		}
	}

	return analysis(ult)


}

func repeater(s string) bool {

	// original
	if len(s) < 3 { return false }

	for i := 2; i < len(s); i++ {
		if string(s[i]) == string(s[i-2]){
			return true
		}
	}

	// offset
	if len(s) < 4 { return false }

	for i := 3; i < len(s); i++ {
		if string(s[i]) == string(s[i-2]){
			return true
		}
	}

	return false

}



// Testing the both parts 

func day5 (strings []string) (int, int) {
	result := 0
	resultv2 := 0

	for i := 0; i < len(strings); i++ {
		if threeVowels(strings[i]) && back2back(strings[i]) && isPure(strings[i]) {
			result += 1
		}
		if havePair(strings[i]) && repeater(strings[i]){
			resultv2 += 1
		}

	}

	return result, resultv2
}

func main() {
	strings := readingFile("day5.csv")

	fmt.Println(day5(strings))
}