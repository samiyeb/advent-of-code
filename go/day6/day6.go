package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type Light struct {
	on bool
	brightness int
}

func turnOn() bool {
	return true
}

func turnOff() bool {
	return false
}

func toggle(l *Light) bool{
	if l.on {
		return false
	} else {
		return true
	}
}

func count(grid [1000][1000]*Light) (int, int) {
	result := 0
	resultv2 := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			resultv2 += grid[i][j].brightness
			if grid[i][j].on {
				result += 1
			}
		}
	}
	return result, resultv2
}

func initialize() [1000][1000]*Light {
	var grid [1000][1000]*Light

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			grid[i][j] = &Light{on: false, brightness: 0}
		}
	}
	return grid
}

func search (startingR int, startingC int, endingR int, endingC int, grid [1000][1000]*Light, flag string) {
	for i := startingR; i <= endingR; i++ {
		for j := startingC; j <= endingC; j++ {
			light := grid[i][j]
			if flag == "toggle" {
				light.on = toggle(light)
				light.brightness += 2

			} else if flag == "off" {
				light.on = turnOff()
				light.brightness -= 1

				if light.brightness < 0 {
					light.brightness = 0
				}

			} else if flag == "on" {
				light.on = turnOn()
				light.brightness += 1
			}
		}	
	}
}

func day6(ins [][]string) (int, int) {
	grid := initialize()

	for i := 0; i < len(ins); i++ {
		if ins[i][0] == "toggle" {
			flag := ins[i][0]
			startingRC := strings.Split(ins[i][1], ",")
			endingRC := strings.Split(ins[i][3], ",")

			startingR, err := strconv.Atoi(startingRC[0])
			startingC, err := strconv.Atoi(startingRC[1])
			endingR, err := strconv.Atoi(endingRC[0])
			endingC, err := strconv.Atoi(endingRC[1])

			if err != nil {
				log.Fatal(err)
			}

			search(startingR, startingC, endingR, endingC, grid, flag)

		} else {
			flag := ins[i][1]
			startingRC := strings.Split(ins[i][2], ",")
			endingRC := strings.Split(ins[i][4], ",")

			startingR, err := strconv.Atoi(startingRC[0])
			startingC, err := strconv.Atoi(startingRC[1])
			endingR, err := strconv.Atoi(endingRC[0])
			endingC, err := strconv.Atoi(endingRC[1])

			if err != nil {
				log.Fatal(err)
			}

			search(startingR, startingC, endingR, endingC, grid, flag)
		}
	}

	return count(grid)
}

func main() {

	if file, err := os.Open("day6.csv"); err != nil {
		log.Fatal(err)
	} else {
		scanner := bufio.NewScanner(file)
		ins := [][]string{}

		for scanner.Scan() {
			ins = append(ins, strings.Fields(scanner.Text()))
		}


		fmt.Println(day6(ins))
	}
}
