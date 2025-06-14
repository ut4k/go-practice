package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	go func() { fmt.Println("別ゴルーチン") }()
	fmt.Println("STOP")
	// キャンセルやタイムアウトされるまでここで止まる
	<-ctx.Done() // チャネルから値を受け取る（待つ）

	fmt.Println("そして時は動き出す")
}
