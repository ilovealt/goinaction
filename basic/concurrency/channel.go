package main

import (
	"fmt"
	"time"
)

// 无缓存，主协程发送，从协程接收
func zero_test1() {
	done := make(chan int)
	go func() {
		fmt.Println("你好, 世界")
		<-done
	}()
	done <- 1
}

// 有一个缓存，主协程发送，从协程接收
func one_test2() {
	done := make(chan int, 1)
	go func() {
		fmt.Println("你好, 世界")
		<-done
	}()
	done <- 1
	fmt.Println("你好, 世界1")
	time.Sleep(time.Second)
}

// 有一个缓存，主协程接收，从协程发送
func one_test3() {
	done := make(chan int, 1)
	go func() {
		fmt.Println("你好, 世界")
		done <- 1
	}()
	<-done
	fmt.Println("你好, 世界1")
}

// 有10个缓存，主协程接收，从协程发送
func ten_test4() {
	done := make(chan int, 10)
	for i := 0; i < cap(done); i++ {
		go func(i int) {
			fmt.Println("你好, 世界 ", i)
			done <- i
		}(i)
	}

	for i := 0; i < cap(done); i++ {
		s := <-done
		fmt.Println("== 你好, 世界 ", s)
	}
}
