package myerror

import "fmt"

type ErrCode int
type MyError struct {
	Code ErrCode
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code: %d", e.Code)
}
