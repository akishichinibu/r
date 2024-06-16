# r

`r` is a library that helps you handle exceptions in a type-safe manner in Golang.
It is inspired by Rust's `Result` type, and built on top of [mo](https://github.com/samber/mo) library.

It provides a `ResultX`(`Result` ~ `Result5`) type that can be used to handle exceptions in a type-safe manner.

You can define your own result type that composited with the `ResultX` type and use it as the return type of your internal implementation. This can help you define the normal and exceptional return values of your function in a proper way. By utilizing the `Unpack` method, you can easily handle the normal and exceptional return values of your function.

## Install
```bash
go get github.com/akishichinibu/r
```

## Usage
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akishichinibu/r"
)

// define you own error types
type UserNotFoundErr struct {
	_      string
	UserID string
}

var _ error = UserNotFoundErr{}

func (e UserNotFoundErr) Error() string {
	return fmt.Sprintf("user not found: userID: %s", e.UserID)
}

type UserRestrictedErr struct {
	_      string
	UserID string
	Reason string
}

var _ error = UserRestrictedErr{}

func (e UserRestrictedErr) Error() string {
	return fmt.Sprintf("user restricted: userID: %s, reason: %s", e.UserID, e.Reason)
}

type GetUserBalanceResult struct {
	r.Result2[int, UserNotFoundErr, UserRestrictedErr]
}

// write your internal implementation
func getUserBalance(userID string) (GetUserBalanceResult, error) {
	var r GetUserBalanceResult
	if userID == "invalid" {
		r.FromFailure1(UserNotFoundErr{
			UserID: userID,
		})
		return r, nil
	}
	if userID == "restricted" {
		r.FromFailure2(UserRestrictedErr{
			UserID: userID,
			Reason: "dummpy reason",
		})
		return r, nil
	}
	if userID == "badguy" {
		return r, fmt.Errorf("internal error")
	}
	r.FromSuccess(1000)
	return r, nil
}

// write a service function
func serviceGetUserBalance(userID string) error {
	result, err := getUserBalance(userID)
	if err != nil {
		return err
	}

	balance, known, ok := result.Unpack()
	if ok {
		fmt.Printf("balance: %d\n", balance)
		return nil
	}

	if e1, ok := known.Failure1(); ok {
		return fmt.Errorf("service failed: %w", e1)
	}

	if e2, ok := known.Failure2(); ok {
		return fmt.Errorf("service failed: %w", e2)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: user_service <user_id>")
		os.Exit(1)
	}
	userID := os.Args[1]
	if err := serviceGetUserBalance(userID); err != nil {
		log.Fatalln(err)
	}
}
```
