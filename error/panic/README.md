## panic💥 / recover 🤲

- あるゴルーチン内でpanicが実行された場合revocerでキャプチャする必要がある
- revocerはtry-catch-finallyでいうcatchにあたる
  - 原則defer文と一緒に利用する
- panicで中断されるとそこで処理が中断される
- recoverしないとキャッチされないpanicが伝搬しプログラムが異常終了する

## panicをパッケージ外に漏らさない！

deferでrecoverでキャプチャしたpanicをエラーにして返してあげる

https://go.dev/wiki/PanicAndRecover#usage-in-a-package

- パッケージ内でpanicを使ってもよいが、呼び出し元（パッケージ外）にpanicが届くと使いづらいので、
- パッケージの公開APIの入口でrecoverを使いpanicをerrorに変換することで安全に扱うことが推奨されている

## AI補足

### panic を使うべき場面

- **プログラムの継続が不可能な致命的なエラー**が起きたとき  
  例：想定外の内部エラー、論理矛盾、致命的なリソース不足など

- **通常のエラーハンドリングが意味をなさない時**  
  例：初期化処理で絶対失敗してはいけない箇所、外部依存の致命的障害

- **ライブラリなどで、利用者が必ず回避・対応すべき誤用を検出したい時**  
  例：引数が無効で「絶対にここを通ってはいけない」など

---

### panic を使わないほうがいい場面

- 一般的なエラー処理は `error` を返して呼び出し元で処理する  
- 予測可能なエラー（ファイルが無い、ネットワークが切れたなど）  
- ユーザ入力のバリデーション失敗など

---

### まとめ

- `panic` は基本的に「ここから先は処理続行できません」という意思表示  
- 通常は `error` を返して上位に処理委譲すべき  
- どうしても回復不能なら `panic` を使う  
- `recover` は `panic` 発生時に「落ちてくる panic をキャッチして安全に処理続行させる」ために使う

---

### 具体例

```go
func mustOpenFile(name string) *os.File {
    f, err := os.Open(name)
    if err != nil {
        panic("ファイルオープンに失敗: " + err.Error())
    }
    return f
}


