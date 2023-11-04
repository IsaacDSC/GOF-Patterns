package main

import (
	"fmt"
	"sync"
	"time"
)

type Tools struct {
	wg        *sync.WaitGroup
	output    *string
	channel01 chan string
	channel02 chan string
}

func Request01(tools Tools) string {
	defer tools.wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("request01_success")
	*tools.output = "request01_success"
	// tools.channel01 <- "request01_success"
	return "request01_success"
}

func Request02(tools Tools) string {
	defer tools.wg.Done()
	time.Sleep(time.Second * 3)
	fmt.Println("request02_success")
	*tools.output = "request02_success"
	return "request02_success"
}

func main() {
	var (
		wg       sync.WaitGroup
		output01 string
		output02 string
	)
	ch01 := make(chan string)
	wg.Add(2)
	go Request01(Tools{wg: &wg, output: &output01, channel01: ch01})
	go Request02(Tools{wg: &wg, output: &output02})
	wg.Wait()
	fmt.Println("output01", output01)
	fmt.Println("output02", output02)
	fmt.Println("channel01", <-ch01)
}
