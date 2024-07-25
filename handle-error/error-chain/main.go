package main

import (
	"errors"
	"fmt"
)

func test1() {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := errors.New("error3")

	err := errors.Join(err1, err2, err3)
	fmt.Println(err)

	errs, ok := err.(interface{ Unwrap() []error })
	if !ok {
		fmt.Println("not imple Unwrap []error")
		return
	}
	fmt.Println(errs.Unwrap())
}

func rootCause(err error) error {
	for {
		e, ok := err.(interface{ Unwrap() error })
		if !ok {
			return err
		}

		err = e.Unwrap()
		if err == nil {
			return nil
		}
	}
}

func test2() {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("2nd err: %w", err1)
	err3 := fmt.Errorf("3nd err: %w", err2)

	fmt.Println(err3)

	fmt.Println(rootCause(err1))
	fmt.Println(rootCause(err2))
	fmt.Println(rootCause(err3))

}

type MyError struct {
	err string
}

func (e *MyError) Error() string {
	return e.err
}

func test3() {
	err1 := &MyError{"temp error"}
	err2 := fmt.Errorf("2nd err: %w", err1)
	err3 := fmt.Errorf("3nd err: %w", err2)

	fmt.Println(err3)

	var e *MyError
	ok := errors.As(err3, &e)
	if ok {
		fmt.Println(e)
		return
	}
}

func main() {
	test1()

	test2()

	test3()
}
