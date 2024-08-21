package main

import (
	"fmt"
	"sync"
	"time"
)

func worker_one(cancel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello...")
		case <-cancel:
			fmt.Println("end ...")
			return
		}
	}
}

func Test_one() {
	cancel := make(chan bool)
	go worker_one(cancel)

	time.Sleep(time.Second)
	cancel <- true
}

func worker_many(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello...")
		case <-cancel:
			fmt.Println("end ...")
			return
		}
	}
}

// 主线程等待子线程结束
func Test_many() {
	cancel := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker_many(&wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)

	wg.Wait()
}
