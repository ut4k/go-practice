## 1個目

```go
func Apply() error {
    var err *MyErr = nil
    return err // <-- 型情報つき！
}
```

- `err`は`*MyErr`型の`nil`ポインタ
- しかし戻り値の型は`errorインターフェース`なのでgoは`err`を`error`に変換します。

## 2個目

```go
func Apply2() error {
    var err error = nil
    return err // <-- nil
}
```

- `err`は最初から`errorインターフェース`で、かつ中身も何もセットされていない`nil`。
- つまり (型: nil, 値: nil) なので、err == nil は true。

----

## 恐ろしい点😱

goは、「ある構造体がどのインターフェースを満たしているかは構造体定義のコードからは確認できない」。
不要に抽象化を行いまくると可読性を下げる。

### サンプル

```go
type MyType struct{}

func (m MyType) Read(p []byte) (int, error) { ... }
// このコードだけでは「MyType が io.Reader を実装しているか？」が一見して わからない。
```

### 対処法

#### 「明示的なコンパイル時アサーション」でインターフェース実装を記述する

```go
type MyType struct{}

var _ io.Reader = (*MyType)(nil)

func (m MyType) Read(p []byte) (int, error) { ... }
```

- これは 「MyType が io.Reader を満たしている」ことを明示し、かつコンパイル時に検証してくれる
  - （じゃあ結局implementsの明示に似てるじゃん！）

- 読む人にも親切（補助ドキュメントのようなもの）

#### インターフェースは最小に保つ（Interface Segregation）

Go では「インターフェースは定義者ではなく、利用者が定義する」スタイルが基本。

```go
// MyPackage では MyService が必要な振る舞いだけ定義
type MyService interface {
    Do()
}
```

→ これにより「必要な部分だけ」がインターフェースに現れ、読解しやすい。

#### インターフェースは使う側で定義する

Go では「インターフェースは定義者ではなく、利用者が定義する」スタイルが基本。

#### 具象型を隠す意味があるときだけ抽象化する

- 「あとで差し替えるかも」みたいな理由で何でもかんでも interface にしない

- 特に main() や cmd パッケージでは 具象型で完結するほうが読みやすい

#### コードリーディングでは go doc や gopls などを活用

----

## 具象型(ぐしょうがた）ってなに？

具象型（concrete type）とは、goにおいて「具体的な構造と実装を持つ型」のこと
（インターフェース（抽象型）にたいするクラスみたいなもん。インスタンス化できる。）

以下はすべて具象型。

```go
type User struct {
    Name string
    Age  int
}

type Point [2]int

type MyInt int

type MyReader struct{}

func (m MyReader) Read(p []byte) (int, error) {
    return 0, nil
}
```
