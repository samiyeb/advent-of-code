package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Node struct {
	visits int
}

type Location struct {
	X int
	Y int
}



func day3(directions string) int{
	g := map[Location]Node{{X:0, Y:0}: {visits:1}}	
	var currXs, currYs int = 0, 0
	var currXr, currYr int = 0, 0
	santa := 1
	robo := 0
	flag := true

	


	for i := 0; i < len(directions); i++{
		d := string(directions[i])

		if d == "^" {
			if flag {
				if node, ok := g[Location{X:currXs, Y:currYs+1}]; ok {
					node.visits += 1
	
				} else {
					g[Location{X:currXs, Y:currYs+1}] = Node{visits:1}
					santa += 1
				}
				currYs += 1
				flag = false

			} else {
				if node, ok := g[Location{X:currXr, Y:currYr+1}]; ok {
					node.visits += 1
	
				} else {
					g[Location{X:currXr, Y:currYr+1}] = Node{visits:1}
					santa += 1
				}
				currYr += 1
				flag = true
			}
			

		} else if d == ">" {
			if flag {
				if node, ok := g[Location{X:currXs+1, Y:currYs}]; ok {
					node.visits += 1
	
				} else {
					g[Location{X:currXs+1, Y:currYs}] = Node{visits:1}
					santa += 1
				}
				currXs += 1
				flag = false

			} else {
				if node, ok := g[Location{X:currXr+1, Y:currYr}]; ok {
					node.visits += 1
	
				} else {
					g[Location{X:currXr+1, Y:currYr}] = Node{visits:1}
					santa += 1
				}
				currXr += 1
				flag = true
			}
			

		} else if d == "v" {
			if flag {
				if node, ok := g[Location{X:currXs, Y:currYs-1}]; ok {
					node.visits += 1
	
				} else {
					g[Location{X:currXs, Y:currYs-1}] = Node{visits:1}
					santa += 1
				}
				currYs -= 1
				flag = false

			} else {
				if node, ok := g[Location{X:currXr, Y:currYr-1}]; ok {
					node.visits += 1
	
				} else {
					g[Location{X:currXr, Y:currYr-1}] = Node{visits:1}
					santa += 1
				}
				currYr -= 1
				flag = true
			}
			

		} else if d == "<" {
			if flag {
				if node, ok := g[Location{X:currXs-1, Y:currYs}]; ok {
					node.visits += 1
	
				} else {
					g[Location{X:currXs-1, Y:currYs}] = Node{visits:1}
					santa += 1
				}
				currXs -= 1
				flag = false

			} else {
				if node, ok := g[Location{X:currXr-1, Y:currYr}]; ok {
					node.visits += 1
	
				} else {
					g[Location{X:currXr-1, Y:currYr}] = Node{visits:1}
					santa += 1
				}
				currXr -= 1
				flag = true
			}
			

		}

	}

	return santa + robo


}

func main(){
	file, err := os.Open("day3.csv")
	if err != nil{
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	directions := ""

	for scanner.Scan(){
		line := scanner.Text()
		directions += line
	}

	


	fmt.Println(day3(directions))
	
	
}