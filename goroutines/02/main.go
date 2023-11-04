package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch01 := make(chan string)
	ch02 := make(chan string)
	go func() {
		defer wg.Done()
		res := Request01()
		fmt.Println("res01", res)
		ch01 <- res
	}()
	go func() {
		defer wg.Done()
		res := Request02()
		fmt.Println("res02", res)
		ch02 <- res
	}()
	fmt.Println("ch01", <-ch01)
	fmt.Println("ch02", <-ch02)
	wg.Wait()

}

func Request01() string {
	time.Sleep(time.Second * 1)
	return "request01_success"
}

func Request02() string {
	time.Sleep(time.Second * 3)
	return "request02_success"
}
