package main

import (
	"fmt"
	"log"
	"time"
)

type MyError struct {
	Time   time.Time
	Reason string
}

func (myerr *MyError) Error() string {
	return fmt.Sprintf("[%v] %s", myerr.Time, myerr.Reason)
}

func badFunction() {
	fmt.Printf("Select Panic type (0=no panic, 1=int, 2=runtime panic)\n")
	var choice int
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		panic(MyError{time.Now(), "Invalid choice"})
	case 2:
		var invalid func()
		invalid()
	}
}

func main() {
	//defer func() {
	//	if pan := recover(); pan != nil {
	//		err := pan.(MyError)
	//		fmt.Printf("Function panicked with an error: %s\n", err)
	//		fmt.Printf("Timestamp: %v", err.Time)
	//
	//		//switch err := pan.(type) {
	//		//default:
	//		//	panic(pan)
	//		//case MyError:
	//		//	fmt.Printf("Function panicked with an error: %s\n", err)
	//		//	fmt.Printf("Timestamp: %v", err.Time)
	//		//}
	//	}
	//}()
	//badFunction()
	//fmt.Printf("Program exited normally\n")
	server()
}

// More examples:
type Work interface{}

var do = func() {
	panic(&MyError{time.Now(), "Work task error"})
}

func server() {
	//for work := range workChan {
	err := safelyDo()
	if err != nil {
		log.Println("Work failed:", err)
	}
	//}
}

func safelyDo() (myErr error) {
	defer func() {
		if err := recover(); err != nil {
			myErr = err.(error)
			return
		}
	}()
	do()
	return
}

// Error is the type of a parse error; it satisfies the error interface.
type Error string

func (e Error) Error() string {
	return string(e)
}

//
//// error is a method of *Regexp that reports parsing errors by
//// panicking with an Error.
//func (regexp *Regexp) error(err string) {
//	panic(Error(err))
//}
//
//// Compile returns a parsed representation of the regular expression.
//func Compile(str string) (regexp *Regexp, err error) {
//	regexp = new(Regexp)
//	// doParse will panic if there is a parse error.
//	defer func() {
//		if e := recover(); e != nil {
//			regexp = nil    // Clear return value.
//			err = e.(Error) // Will re-panic if not a parse error.
//		}
//	}()
//	return regexp.doParse(str), nil
//}
