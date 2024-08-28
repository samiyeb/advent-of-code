package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	
)

func min(val1 int, val2 int, val3 int) int{
	result := val1

	if result > val2{
		result = val2
	}
	if result > val3{
		result = val3
	}

	return result
}

func day2(presentSizes []string) (int, int){
	result := 0
	resultv2 := 0
	for i := 0; i < len(presentSizes); i++ {
		temp := strings.Split(presentSizes[i], "x")
		var l, w, h int
		var err error
		
		if l, err = strconv.Atoi(temp[0]); err != nil {
			log.Fatal("0")
		}
		if w, err = strconv.Atoi(temp[1]); err != nil {
			log.Fatal("1")
		}
		if h, err = strconv.Atoi(temp[2]); err != nil {
			log.Fatal("2")
		}

		bowLen := l*w*h
		wrapLen := min(2*l+2*w, 2*w+2*h, 2*h+2*l)

		totalRibbonLen := bowLen + wrapLen

		totalSurfaceArea := 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
		
		result += totalSurfaceArea
		resultv2 += totalRibbonLen

		
	} 
	return result, resultv2
}

func main() {
	file, err := os.Open("day2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var presentSizes [1000]string
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		presentSizes[i] = line
		i += 1
	}

	part1, part2 := day2(presentSizes[:])

	fmt.Println(part1)
	fmt.Println(part2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}