package main

import "fmt"

type MyErr struct{}

func (me *MyErr) Error() string { return "" }

func Apply() error {
	var err *MyErr = nil
	// 戻り値に型情報が含まれているので予想外の挙動をする
	return err
}

func Apply2() error {
	var err error = nil
	// 結果的には問題がないが明示的にnilをreturnしてミスを防ぐほうがよい
	return err
}

func main() {
	fmt.Println(Apply() == nil)  // false
	fmt.Println(Apply2() == nil) // true
}
