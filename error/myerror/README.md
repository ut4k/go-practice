## 独自のエラー構造体

- 独自のエラーを識別したいときは独自の構造体を用意する
- 慣習としてプレフィックスは`Err-`をつけるとよし
- もしエラーチェーンの詳細を保持しておく場合は`error`インターフェースを満たすために`Error`、`Unwrap`、`As`、`Is`を実装する必要あり
- 独自エラーを使う場合でも、関数やメソッドの戻り値は`error`インターフェースをつかう慣習にすべき。
  - 呼び出し側が特定の具体型に依存しなくなる
  - 慣習だから。様々なパッケージも`error`インターフェースを返すつくりになってる


### 具体型を返す場合（X）

```go
package main

import "fmt"

type MyError struct {
	Code int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error code: %d", e.Code)
}

// MyErrorを返す
func DoSomething(bad bool) *MyError {
	if bad {
		return &MyError{Code: 123}
	}
	return nil
}

func main() {
	err := DoSomething(true) // 戻り値がMyError固定
	if err != nil {
		fmt.Println("具体型のエラー発生:", err)
	}
}
```


### インターフェースを返す場合（O）

```go
package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error code: %d", e.Code)
}

// errorを返す
func DoSomething(bad bool) error {
	if bad {
		return &MyError{Code: 123}
	}
	return nil
}

func main() {
	err := DoSomething(true) // 戻り値はなんらかのerrorインターフェース
	if err != nil {
		fmt.Println("error インターフェースのエラー発生:", err)
	}

	// 型アサーションで詳細情報を取得可能
	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Printf("MyErrorの詳細コード: %d\n", myErr.Code)
	}
}
```
