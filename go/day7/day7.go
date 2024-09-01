package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readingFiles(fileName string) [][]string{
	var result [][]string

	if file, err := os.Open(fileName); err != nil {
		log.Fatal(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			result = append(result, strings.Fields(scanner.Text()))
		}
	}

	return result

}

func check(k string, v map[string]int ) bool {
	for k2 := range v {
		if k == k2 {
			return true
		}
	}
	return false
}

func day7(ins [][]string) int {
	vars := map[string]int{}
	temp := -1

	for t := 0; t < 2; t++ {
		for !(check("a", vars)) {
			for i := 0; i < len(ins); i++ {
				if ins[i][0] == "NOT" {
					key := ins[i][1]
	
					if check(key, vars) {
						target := ins[i][3]
						vars[target] = ^vars[key]
					}
		
				} else if val, err := strconv.Atoi(ins[i][0]); err == nil{
					if ins[i][1] != "->" {
						operator := ins[i][1]
						key := ins[i][2]
	
						if check(key, vars) {
							target := ins[i][4]
							if operator == "AND" {
								vars[target] = val & vars[key]
							} else if operator == "OR" {
								vars[target] = val | vars[key]
							} else if operator == "LSHIFT" {
								vars[target] = val << vars[key]
							} else if operator == "RSHIFT" {
								vars[target] = val >> vars[key]
							}
						}
					} else {
						if !(ins[i][2] == "b" && temp != -1){
							vars[ins[i][2]] = val
						}
					}
					
		
				} else if ins[i][1] == "->"{
					key := ins[i][0]
	
					if check(key, vars) {
						target := ins[i][2]
						vars[target] = vars[key]
					}
		
				} else {
					leftKey := ins[i][0]
					rightKey := ins[i][2]
					operator := ins[i][1]
	
					if val, err := strconv.Atoi(rightKey); err == nil {
						if check(leftKey, vars){
							target := ins[i][4]
							if operator == "AND" {
								vars[target] = vars[leftKey] & val
							} else if operator == "OR" {
								vars[target] = vars[leftKey] | val
							} else if operator == "LSHIFT" {
								vars[target] = vars[leftKey] << val
							} else if operator == "RSHIFT" {
								vars[target] = vars[leftKey] >> val
							}
						}
					} else {
						if check(leftKey, vars) && check(rightKey, vars) {
							target := ins[i][4]
							if operator == "AND" {
								vars[target] = vars[leftKey] & vars[rightKey]
							} else if operator == "OR" {
								vars[target] = vars[leftKey] | vars[rightKey]
							} else if operator == "LSHIFT" {
								vars[target] = vars[leftKey] << vars[rightKey]
							} else if operator == "RSHIFT" {
								vars[target] = vars[leftKey] >> vars[rightKey]
							}
						}
	
					}
	
				}
			}
		}

		if t == 1 { break }

		temp = vars["a"]
		vars = nil
		vars = map[string]int{"b": temp}


	}

	

	return vars["a"]
}

func main() {
	fmt.Println(day7(readingFiles("day7.csv")))
}