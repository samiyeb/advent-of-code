package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func day1 (d string) (int, int) {
	result := 0
	position := -1
	notVisited := true
	for i := 0; i < len(d); i++ {
		if result == -1 && notVisited {
			position = i
			notVisited = false
		}
		if d[i] == '(' {
			result += 1
		} else {
			result -= 1
		}
	}
	return result, position
}

func main() {
	file, err := os.Open("day1.csv")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	directions := ""

	for scanner.Scan() {
		directions += scanner.Text()
	}

	fmt.Println(day1(directions))
}