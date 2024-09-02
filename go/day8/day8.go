package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func newTotal (fileName string) int {
	if file, err := os.Open(fileName); err == nil {
		scanner := bufio.NewScanner(file)
		result := 0

		for scanner.Scan() {
			i := 0
			temp := 0
			for i < len(scanner.Text()) {
				c := string(scanner.Text()[i])

				if c == `\` {
					if i + 1 < len(scanner.Text()) - 1 && string(scanner.Text()[i+1]) == `x` {
						i += 4
						temp += 4
					} else {
						i += 2
						temp += 3
					}
				} else {
					i++
				}
				temp += 1
			}
			result += temp + 4
		}

		return result
	} else {
		log.Fatal(err)
	}

	return -1

}

func total (fileName string) int {
	if file, err := os.Open(fileName); err == nil {
		scanner := bufio.NewScanner(file)
		result := 0

		for scanner.Scan() {
			result += len(scanner.Text())
		}

		return result
	
	} else {
		log.Fatal(err)
	}

	return -1

}

func inMemory (fileName string) int {
	if file, err := os.Open(fileName); err == nil {
		scanner := bufio.NewScanner(file)
		result := 0

		for scanner.Scan() {
			i := 1
			temp := 0
			for i < len(scanner.Text()) - 1 {
				c := string(scanner.Text()[i])

				if c == `\` {
					if i + 1 < len(scanner.Text()) - 1 && string(scanner.Text()[i+1]) == `x` {
						i += 4
					} else {
						i += 2
					}
				} else {
					i++
				}
				temp += 1
			}
			result += temp
		}

		return result
	} else {
		log.Fatal(err)
	}

	return -1

}

func main() {
	fmt.Println(total("day8.csv") - inMemory("day8.csv"))
	fmt.Println(newTotal("day8.csv") - total("day8.csv"))
	
}
