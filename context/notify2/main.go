package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	task := make(chan int)
	go func() {
		//無限ループ
		for {
			select {
			// チャネルに何かが届いてる
			case <-ctx.Done():
				return
			// taskチャネルに何かが届いてる
			case i := <-task:
				fmt.Println("受信した:", i)
			default:
				fmt.Println("キャンセルされていない")
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()
	time.Sleep(time.Second)
	for i := 0; 5 > i; i++ {
		task <- i
	}
	cancel()
}
