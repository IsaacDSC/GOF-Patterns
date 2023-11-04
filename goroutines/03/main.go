package main

import (
	"fmt"
	"time"
)

type Tools struct {
	operationName string
}

func main() {
	ch01 := make(chan Tools)
	ch02 := make(chan Tools)
	go func() {
		res := Request01()
		ch01 <- Tools{operationName: res}
	}()
	go func() {
		res := Request02()
		ch02 <- Tools{operationName: res}
	}()
	value01, ok := <-ch01
	if ok {
		fmt.Println("ch01", value01)
	}
	value02, ok := <-ch02
	if ok {
		fmt.Println("ch02", value02)
	}

}

func Request01() string {
	time.Sleep(time.Second * 3)
	return "request01_success"
}

func Request02() string {
	time.Sleep(time.Second * 3)
	return "request02_success"
}
