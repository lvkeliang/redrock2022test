package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("出现")
		ch <- 1
	}()
	<-ch
}
