package main

import "fmt"

// 执行if语句中的return后，向上依次执行defer语句，所以只执行到了输出1的defer，没执行下面输出3的defer
func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()
	if a {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}
