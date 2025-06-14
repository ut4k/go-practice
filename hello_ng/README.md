- 同一パッケージ内に２回`func main`は定義してはならない。(エントリーポイントが２個になるから)

- `go run .`は以下のエラーになり動作しない。

```
./hello2.go:5:6: main redeclared in this block
	./hello.go:5:6: other declaration of main
```


