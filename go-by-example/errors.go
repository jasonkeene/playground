package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 5, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 10, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e == nil {
			fmt.Println("f1 worked:", r)
		} else {
			fmt.Println("f1 failed:", e)
		}
		if r, e := f2(i); e == nil {
			fmt.Println("f2 worked:", r)
		} else {
			fmt.Println("f2 failed:", e)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
