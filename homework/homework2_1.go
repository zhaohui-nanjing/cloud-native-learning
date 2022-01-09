package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var fmeng = false

func main() {

	chanStream := make(chan string, 10)

	//生产者和消费者做一个计数器
	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)

	for i := 0; i < 3; i++ {
		wgPd.Add(1)
		go producer(i, wgPd, chanStream)
	}

	for i := 0; i < 2; i++ {
		wgCs.Add(1)
		go consumer(wgCs, chanStream)
	}
	wgPd.Wait()

	go func() {
		time.Sleep(time.Second * 3)
		fmeng = true
	}()

	//生产完成，关闭channel
	close(chanStream)
	wgCs.Wait()
}

func producer(threadId int, wg *sync.WaitGroup, ch chan string) {
	count := 0

	for !fmeng {
		time.Sleep(time.Second * 1)
		count++
		data := strconv.Itoa(threadId) + "---" + strconv.Itoa(count)
		fmt.Printf("producer: %s\n", data)
		ch <- data
	}

	wg.Done()
}

func consumer(wg *sync.WaitGroup, ch chan string) {
	for data := range ch {
		time.Sleep(time.Second * 1)
		fmt.Printf("Consumer: %s\n", data)
	}
	wg.Done()
}
