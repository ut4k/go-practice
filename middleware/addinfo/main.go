package main

import (
	"fmt"
	"log"
	"net/http"
)

type AppVersion string

// 追加情報を付与
func VersionAdder(v AppVersion) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//リクエストヘッダに付与
			r.Header.Add("App_Version", string(v))
			//レスポンスヘッダに付与
			w.Header().Set("App-Version", string(v))
			next.ServeHTTP(w, r)
		})
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	ver := r.Header.Get("App_Version")
	fmt.Fprintf(w, "Hello! App Version in request header: %s\n", ver)
}

func main() {
	version := AppVersion("v1.2.3")

	// ハンドラにミドルウェアを適用
	mux := http.NewServeMux()
	mux.Handle("/", VersionAdder(version)(http.HandlerFunc(HelloHandler)))

	log.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
