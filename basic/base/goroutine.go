package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main2() {
	fmt.Println("Hello Golang!")
	fmt.Println("")

	//gaa()
	//gbb()
	//gcc()
	//gdd()
	//gee()
	//gff()
	//ggg()
	//ghh()
	//gii()
	//gjj()
	//gkk()
	//gll1()
	//gll2()
	//gmm()
	// gnn()
	// goo()
	// gpp()
	// gqq()
	// grr()
	// gss()
	// gtt()
	// guu()
	// gww()
	gxx()
}

func gxx() {

}

func gww() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func gvvWriteFile() {
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat1")
	err := os.WriteFile(path1, d1, 0644)
	check(err)
}

func gvv() {
	// gvvWriteFile()
}

func guu() {
	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}

func gtt() {
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}

func gss() {
	p := fmt.Println
	now := time.Now()
	p("Now:", now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("Then:", then)

	p("Year:", then.Year())
	p("Month:", then.Month())
	p("Day:", then.Day())
	p("Hour:", then.Hour())
	p("Minute:", then.Minute())
	p("Second:", then.Second())
	p("Nanosecond:", then.Nanosecond())
	p("Location:", then.Location())

	p("Weekday:", then.Weekday())

	p("Unix:", then.Unix())
	p("UnixNano:", then.UnixNano())

	p("Format:", then.Format(time.RFC3339))
	// 自定义格式化输出, 注意格式必须是固定的时间点
	// 2006-01-02 15:04:05 Monday 是 Go 语言中时间格式化的参考时间
	// 通过这种方式，Go 语言可以识别出各个时间组件的位置和格式
	// 然后根据这个参考时间来格式化实际的时间值
	// 这种设计虽然有些独特，但一旦理解了参考时间的含义，就能灵活地进行各种时间格式化操作
	p("Format:", then.Format("2006-01-02 15:04:05 Monday"))

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p("Diff:", diff)

	p("Hours:", diff.Hours())
	p("Minutes:", diff.Minutes())
	p("Seconds:", diff.Seconds())
	p("Nanoseconds:", diff.Nanoseconds())

	p("Add:", then.Add(diff))
	p("Add:", then.Add(-diff))
	p("AddDate:", then.AddDate(1, 2, 3))

}

func grr() {
	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("Count:     ", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("Index:     ", strings.Index("test", "e"))
	fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", strings.Repeat("a", 5))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", strings.ToLower("TEST"))
	fmt.Println("ToUpper:   ", strings.ToUpper("test"))
}

func mayPanic() {
	panic("a problem")
}

func gqq() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")
}

func createFile(p string) *os.File {
	fmt.Println("Creating file:", p)
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing to file:", f.Name())
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing file:", f.Name())
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func gpp() {
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

type Person struct {
	name string
	age  int
}

func goo() {
	fruits := []string{"peach", "banana", "kiwi"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	people := []Person{
		{"Bob", 31},
		{"John", 22},
		{"Alice", 27},
	}

	slices.SortFunc(people, func(a, b Person) int {
		if c := cmp.Compare(a.name, b.name); c != 0 {
			return c
		}
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
}

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[name]++
}

func gnn() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("a", 10000)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("a", 10000)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("b", 10000)
	}()

	wg.Wait()
	fmt.Println(c.counters)

}

func gmm() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for range 50 {
		go func() {
			wg.Add(1)
			defer wg.Done()
			for range 1000 {
				ops.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops.Load())
}

// 速率限制器 示例
// 通过使用 time.Tick 函数创建一个定时器通道，我们可以限制对某个资源的访问速率
// 在这个例子中，我们创建了一个每200毫秒发送一个时间戳的通道
// 然后在处理请求时，我们从这个通道接收时间戳，从而实现了对请求处理的速率限制
// 这样可以防止系统过载，确保资源被合理利用
func gll1() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 创建一个Ticker
	limiter := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}

func gll2() {

	burstyLimiter := make(chan time.Time, 3)

	for range 3 {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func gkk() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker1(i, &wg)
	}
	wg.Wait()

}

func worker1(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func gjj() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
	close(results)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func gii() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func ghh() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// Stop timer2 before it fires
	stoped := timer2.Stop()
	if stoped {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

func ggg() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	fmt.Println("queue length:", len(queue))

	close(queue)

	for elem := range queue {
		fmt.Println("Received:", elem)
	}
}

func gff() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	<-done

	_, ok := <-jobs
	fmt.Println("Channel closed:", !ok)
}

func gee() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 3")
	}

}

// Select语句用于在多个通道操作中选择一个可用的操作
// 当有多个通道准备好时，select会随机选择一个执行
// 如果没有通道准备好，select会阻塞直到至少有一个通道准备好
func gdd() {
	// 创建两个通道
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("Received from c1:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received from c2:", msg2)
		}
	}

}

func gcc() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// 关闭通道后，不能再发送数据到通道
	// 但可以从通道中接收数据
	// 如果通道已经关闭，接收操作会立即返回零值
	// 如果通道未关闭，接收操作会阻塞直到有数据可接收
	// close() 函数用于关闭通道
	close(pings)
	close(pongs)
	fmt.Println("Channels closed")
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func gbb() {
	// 创建一个带缓冲区的通道，缓冲区大小为2
	// 缓冲区允许发送和接收操作异步进行
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func gaa() {
	// 通道没有缓冲，所以发送和接收必须是同时发生的
	message := make(chan string)

	go func() {
		fmt.Println("准备 发送。。。")
		message <- "ping"
		fmt.Println("发送 完成")
	}()

	fmt.Println("准备 接收。。。")
	msg := <-message
	fmt.Println(msg)
	fmt.Println("接收 完成")

}
