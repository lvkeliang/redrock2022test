package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int = 0
	//var input string = ""
	var time []int
	var quality []int
	var qualityType []int
	var qualityNum int = 0
	var output string = ""

	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Scanf("%v\n", &n)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	//fmt.Printf("scanf : %v\n", n)
	for i := 0; i < n; i++ {
		input, err := reader.ReadString('\n')
		//_, err := fmt.Scanf("%v", &input)
		//fmt.Printf("input%v : %v\n", i, input)
		if err != nil {
			fmt.Printf("err1\n")
			fmt.Printf("%v\n", err)
			return
		}
		arr := strings.Split(input, " ")
		length := len(arr)
		arr[length-1] = strings.TrimRight(arr[length-1], "\r\n")
		temp, err := strconv.Atoi(arr[0])
		if err != nil {
			fmt.Printf("err2\n")
			fmt.Printf("%v\n", err)
			return
		}
		time = append(time, temp)
		//fmt.Printf("arr[] : %v\n", arr)
		for i := 1; i < length; i++ {
			temp, err := strconv.Atoi(arr[i])
			if err != nil {
				fmt.Printf("err3\n")
				fmt.Printf("arr[%v] : %v\n", i, arr[i])
				fmt.Printf("%v\n", err)
				return
			}
			quality = append(quality, temp)
		}
		for _, value1 := range quality {
			flag := true
			for _, value2 := range qualityType {
				if value1 == value2 {
					flag = false
					break
				}
			}
			if flag {
				qualityType = append(qualityType, value1)
				qualityNum++
			}
		}
		output += strconv.Itoa(qualityNum)
		output += "\n"

	}
	fmt.Print(output)

}
