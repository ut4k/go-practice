package main

import (
	"log"
	"net/http"
	"time"
)

// ミドルウェア: 処理時間を計測してログ出力
func MyMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 本来のハンドラを実行
		h.ServeHTTP(w, r)

		duration := time.Since(start).Microseconds()
		log.Printf("Request took %d ms", duration)
	})
}

// 実際の処理
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	// ハンドラにミドルウェアを適用
	http.Handle("/", MyMiddleware(http.HandlerFunc(HelloHandler)))

	// サーバー起動
	log.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// なぜかCtrl+c/Ctrl+dでサーバープロセス切れないので注意・・・
