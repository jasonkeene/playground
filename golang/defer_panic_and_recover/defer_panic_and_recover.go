package main

import (
	"errors"
	"fmt"
)

// This function's defer can mutate returned variables
func a() (x int) {
	x = 0
	defer func() {
		x++
	}()
	return x
}

// This function's defer can not mutate returned variables
func b() int {
	x := 0
	defer func() {
		x++
	}()
	return x
}

func safe_call(work func()) (x int, err error) {
	x = 42
	defer func() {
		if e := recover(); e != nil {
			x = 0
			err = errors.New(e.(string))
		}
	}()
	work()
	return
}

func main() {
	fmt.Println(a()) // outputs 1
	fmt.Println(b()) // outputs 0

	val, err := safe_call(func() {
		fmt.Println("Doing normal work")
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("val was: %d\n", val)
	}

	val, err = safe_call(func() {
		panic("OMG!")
	})
	if err != nil {
		fmt.Println(err) // outputs OMG! but doesn't halt
	} else {
		fmt.Printf("val was: %d\n", val)
	}
}
