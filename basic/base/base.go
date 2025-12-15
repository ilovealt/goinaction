package main

import (
	"bytes"
	"errors"
	"fmt"
	"iter"
	"log"
	"log/slog"
	"maps"
	"math"
	"os"
	"slices"
	"time"
)

func main() {

	fmt.Println("Hello Golang!")
	fmt.Println("")
	//aa()
	//bb()
	//cc()
	//dd()
	//ee()
	//ff()
	//gg()
	//hh()
	//ii()
	//jj()
	//kk()
	//ll()
	//mm()
	//nn()
	//oo()
	//pp()
	//qq()
	//rr()
	//ss()
	//tt()
	// uu()
	vv()
}

func vv() {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello")
	fmt.Print("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	myslog.Info("hello again", "key", "val", "age", 25)
}

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func uu() {
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return -1
}

type element[T any] struct {
	value T
	next  *element[T]
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{value: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{value: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) Pop() (v T, ok bool) {
	if lst.head == nil {
		var zero T
		return zero, false
	}
	v = lst.head.value
	lst.head = lst.head.next
	if lst.head == nil {
		lst.tail = nil
	}
	return v, true
}
func (lst *List[T]) IsEmpty() bool {
	return lst.head == nil
}

func (lst *List[T]) Size() int {
	size := 0
	for e := lst.head; e != nil; e = e.next {
		size++
	}
	return size
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.value)
	}
	return elems
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.value) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

// Range over Iterators
func tt() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}

func ss() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println(SlicesIndex(s, "bar"))
	fmt.Println(SlicesIndex(s, "baz"))

	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)

	fmt.Println("List size:", lst.Size())
	fmt.Println("All elements:", lst.AllElements())

	for !lst.IsEmpty() {
		v, _ := lst.Pop()
		fmt.Println("Popped:", v)
	}
	fmt.Println("List size after pops:", lst.Size())
}

type base struct {
	num int
}

type describer interface {
	describe() string
}

func (b base) describe() string {
	return fmt.Sprintf("base num = %d", b.num)
}

type container struct {
	base
	str string
}

func rr() {
	c := container{
		base: base{num: 12},
		str:  "hello",
	}
	fmt.Printf("c = %v\n", c)
	fmt.Printf("c={num: %v, str: %v}\n", c.num, c.str)
	fmt.Println("also num:", c.base.num)

	fmt.Printf("c.describe() = %s\n", c.describe())

	var d describer = c
	fmt.Printf("d.describe() = %s\n", d.describe())
}

type ServerState int

const (
	STATE_IDLE ServerState = iota
	STATE_CONNECTED
	STATE_ERROR
	STATE_RETRYING
)

func qq() {
	var state ServerState = STATE_CONNECTED

	switch state {
	case STATE_IDLE:
		fmt.Println("Server is idle")
	case STATE_CONNECTED:
		fmt.Println("Server is connected", STATE_CONNECTED)
	case STATE_ERROR:
		fmt.Println("Server has encountered an error")
	case STATE_RETRYING:
		fmt.Println("Server is retrying connection")
	default:
		fmt.Println("Unknown server state")
	}
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Printf("area: %.6f\n", g.area())
	fmt.Printf("perim: %.6f\n", g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("Circle radius:", c.radius)
	}
}

func pp() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}

func oo() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func nn() {

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := newPerson("Jonh")
	fmt.Println(s.name)

}

func mm() {
	fmt.Println("Factorial of 7 is:", fact(7))

	var fib func(n int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println("Fibonacci of 7 is:", fib(7))
}

func fact(n int) int {
	// 递归终止条件
	if n == 0 {
		return 1
	}
	// 递归调用,计算阶乘,n!
	return n * fact(n-1)
}

func ll() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt2 := intSeq()
	fmt.Println(newInt2())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func kk() {
	a, b := vals()
	fmt.Println(a, b)
}

func vals() (int, int) {
	return 3, 7
}

func jj() {
	//m := make(map[string]int) —— 可直接写入，运行时决定初始内部大小。
	//m := make(map[string]int, 1000) —— 运行时会尽量为约 1000 个元素预分配空间以减少扩容。
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	// 从 map 索引可以返回两个值：第一个是对应的值，第二个是表示键是否存在的布尔值 (ok)。
	// 示例：同时接收值和值存在标志并检查键是否存在。
	v2, ok := m["k1"]
	fmt.Println("v2:", v2, "exists:", ok)

	// 如果只关心是否存在，可以用 '_' 忽略第一个返回值，只保留存在标志：
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

func ii() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t equal t2")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func hh() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func gg() {
	switch os := "linux"; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type")
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func ff() {
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func ee() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 虽然上面已经使用，但是这里会重新初始化i的范围
	for i := range 3 {
		fmt.Println("range", i)
	}

	for j := 0; j <= 6; j++ {
		fmt.Println(j)
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func dd() {
	// 函数外部定义的常量可以调用，但是函数内部的常量不能在外部调用
	fmt.Println("Hello", PI)
}

const PI = 3.14

func cc() {
	fmt.Println("PI =", PI)

	const World = "世界"
	fmt.Println("Hello", World)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}

func bb() {
	var a = "hello"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var e = true
	fmt.Println(e)

	d := 3.14
	fmt.Println(d)

}

func aa() {
	fmt.Println("go" + "lang")

	fmt.Println("1 + 1 =", 1+1)

	fmt.Println("7.0 / 3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
