package main

import (
	"fmt"
	"time"
)

// 生产者
func producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func run_pro_con() {
	ch := make(chan int, 64)

	go producer(3, ch)
	go producer(5, ch)
	go consumer(ch)

	time.Sleep(5 * time.Millisecond)
}
