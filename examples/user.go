package examples

import (
	"fmt"

	"github.com/akishichinibu/r"
)

func empty[T any]() (t T) {
	return
}

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

// Service Layer
func GetUserBalance(userID string) (GetUserBalanceResult, error) {
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

// Handler Layer
func UserGetUserBalance() error {
	userID := "invalid"
	result, err := GetUserBalance(userID)
	if err != nil {
		return err
	}

	balance, known, ok := result.Unpack()
	if ok {
		fmt.Printf("balance: %d\n", balance)
		return nil
	}

	if e1, ok := known.Failure1(); ok {
		fmt.Println(e1.Error())
		// or return a 400 response
		return nil
	}

	if e2, ok := known.Failure2(); ok {
		fmt.Println(e2.Error())
		return nil
	}

	return nil
}
