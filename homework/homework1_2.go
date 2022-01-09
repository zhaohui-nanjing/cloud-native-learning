package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	myChan := make(chan int, 10)
	defer close(myChan)

	//Consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			fmt.Println("receive: ", <-myChan)
		}
	}()

	//Producer
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10)
		myChan <- n
		time.Sleep(1 * time.Second)
	}
}
