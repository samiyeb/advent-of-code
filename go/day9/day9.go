package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Trip struct {
	start string
	end string
	distance int
}

func readingFile (fileName string) [][]string {
	strs := [][]string{}
	if file, err := os.Open(fileName); err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			strs = append(strs, strings.Fields(scanner.Text()))
		}
		return strs
	} else {
		log.Fatal(err)
	}
	return strs
}

func day9(listOfStrings [][]string) int {
	trips := make(map[string][]Trip)
	locationVisited := make(map[string]bool)

	for _, s := range listOfStrings {
		start := s[0]
		end := s[2]
		distance, err := strconv.Atoi(s[4])

		if err != nil {log.Fatal(err)}

		if _, ok := trips[start]; ok {
			trips[start] = append(trips[start], Trip{start: start, end: end, distance: distance})
		} else {
			trips[start] = []Trip{ {start: start, end: end, distance: distance} }
			locationVisited[start] = false
		}

		if _, ok := trips[end]; ok {
			trips[end] = append(trips[end], Trip{start: end, end: start, distance: distance})
		} else {
			trips[end] = []Trip{ {start: end, end: start, distance: distance} }
			locationVisited[end] = false
		}
	}

	global := math.MinInt

	var dfs func (visits int, currentPlace string, acc int)
	dfs = func (visits int, currentPlace string, acc int){
		locationVisited[currentPlace] = true
		if visits == len(locationVisited){
			if acc > global {
				global = acc
			}
		} else {
			for _, t := range trips[currentPlace] {
				if !(locationVisited[t.end]){
					dfs(visits+1, t.end, acc + t.distance)
					locationVisited[t.end] = false
				}
			}
		}
	}

	for l := range locationVisited {
		dfs(1, l, 0)
	}
	
	return global

    
}

func main() {
	fmt.Println(day9(readingFile("day9.csv")))
}