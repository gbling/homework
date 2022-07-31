package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var fstop = false

func producer(threadId int, wg *sync.WaitGroup, ch chan string) {
	count := 0
	for !fstop {
		time.Sleep(time.Second * 1)
		count++
		data := strconv.Itoa(threadId) + "---" + strconv.Itoa(count)
		fmt.Printf("producer, %s\n", data)
		ch <- data
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup, ch chan string) {
	for data := range ch {
		time.Sleep(time.Second * 1)
		fmt.Printf("consumer %s\n", data)
	}
	wg.Done()
}

func main() {

	chanSteam := make(chan string, 10)

	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)

	// 5个生产者
	for i := 0; i < 5; i++ {
		wgPd.Add(1)
		go producer(i, wgPd, chanSteam)
	}

	// 3个消费者
	for j := 0; j < 3; j++ {
		wgCs.Add(1)
		go consumer(wgCs, chanSteam)
	}

	// 独立一个协程设置超时
	go func() {
		time.Sleep(time.Second * 3)
		fstop = true
	}()

	wgPd.Wait()
	close(chanSteam)
	wgCs.Wait()
}
