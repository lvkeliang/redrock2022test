package main

import (
	"fmt"
	"sync"
)

func number(num int) int {
	var res int = 0
	for ; num > 0; num /= 10 {
		res++
	}
	//fmt.Printf("number res : %v\n", res)
	return res
}

func pow(num int, n int) int {
	var res int = 1
	for i := 0; i < n; i++ {
		res *= num
	}
	//fmt.Printf("pow num : %v\n", num)
	//fmt.Printf("pow res : %v\n", res)
	return res
}

func isArmstrongNum(num int) bool {
	//fmt.Printf("isArmstrongNuming : %v\n", num)
	n := number(num)
	temp := num
	var sum int = 0
	for i := 0; i < n; i++ {
		a := temp % 10
		temp /= 10
		sum += pow(a, n)
	}
	if sum == num {
		return true
	}
	return false
}

var wg sync.WaitGroup

func main() {
	wg.Add(6)
	go func() {
		for i := 0; i < 10; i++ {
			if isArmstrongNum(i) {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 10; i < 100; i++ {
			if isArmstrongNum(i) {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 100; i < 1000; i++ {
			if isArmstrongNum(i) {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 1000; i < 10000; i++ {
			if isArmstrongNum(i) {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 10000; i < 100000; i++ {
			if isArmstrongNum(i) {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 100000; i < 1000000; i++ {
			if isArmstrongNum(i) {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
