## context.Contextの要点

- APIのリクエストスコープの値を含める
- 関数への追加引数は含めない
- `context.Context`型の値は複数のゴルーチンから同時に使われても安全

関数ロジックに関わる値を含めてはいけいない！
極端にいうと関数への外部データはすべて`context.Context`に含められるのだが、
必要な情報がその中に隠蔽されていた場合読みてはそのメソッドが必要となる情報を判断できなくなる。

## メモ書き

どこまでがコンテキストで、どこからが関数パラメータかの境界線は経験を積んでみないと判断むずそう・・・

