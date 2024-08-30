package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
	"encoding/hex"
)

func day4(m string) int{
	length := len(m)

	endingStr := "9"

	for i := 0; i < length - 1; i++ {
		endingStr += "9"
	}

	endingNum, err := strconv.Atoi(endingStr) 

	if err != nil {
		log.Fatal(err)
	}

	for num := 0; num <= endingNum; num++ {
		secondHalf := strconv.Itoa(num)

		data := md5.Sum([]byte(m + secondHalf))

		var input []byte

		for i := 0; i < len(data); i++ {
			input = append(input, data[i])
		}

		h := hex.EncodeToString(input)
		

		if h[0] == '0' && h[1] == '0' && h[2] == '0' && h[3] == '0' && h[4] == '0' && h[5] == '0' {
			return num
		}
		


	} 

	return -1

}

func main(){
	fmt.Println(day4("yzbqklnj"))
}