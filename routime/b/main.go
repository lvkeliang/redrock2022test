package main

import (
	"fmt"
	"time"
)

// var wg sync.WaitGroup
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	done := make(chan int)
	//wg.Add(3)
	go Work("goroutine1", ch1, ch2)
	go Work("goroutine2", ch2, ch3)
	go Work("goroutine3", ch3, done)
	ch1 <- 1
	<-done
	//wg.Wait()
	fmt.Println("successful")
}
func Work(workName string, chin chan int, chou chan int) {
	<-chin
	time.Sleep(time.Second) // 模拟逻辑处理
	fmt.Println(workName)
	chou <- 1
	//wg.Done()
}
