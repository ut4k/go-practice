package myerror

import (
	"testing"
)

func TestMyError_Error(t *testing.T) {
	err := &MyError{Code: 123}
	want := "code: 123"
	got := err.Error()

	// 注意：t.Logfはgo test -vじゃないと成功時は何も出力されない
	if got == want {
		t.Logf("OK: got %q == want %q", got, want)
	}
}
