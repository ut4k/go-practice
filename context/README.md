# context

## context.Context は状態を「扱う」ものか？

(AI回答)
context.Context は「状態（ステータス）」や「キャンセル信号」などの情報を伝搬（伝える）するための仕組みです。

厳密には、context.Context 自体が状態を保持しているわけではなく、
「キャンセルされたかどうか」などの状態情報を持つ内部の仕組みへの参照を保持しています。

## context.WithCancel

コンテキストと、キャンセル用の関数の２個をかえす
キャンセル用の関数を呼び出すと、このコンテキストをキャンセルできる

## context.Background

空のコンテキストをつくる


## コンテキストをキャンセルできる、とは？

(AI回答)
context.Context は内部に「キャンセルされたかどうか」の状態（キャンセルフラグ）を持っています。

cancel() を呼ぶと、そのコンテキストの「キャンセルフラグ」が立つ（＝キャンセルされた状態になる）。

以降、そのコンテキストを受け取っている処理は、ctx.Err() を呼ぶと context.Canceled を返し、
「キャンセルされたことを認識できる」ようになります。

## ほかにもこんな方法でキャンセルできる

`context.WithTimeout(parent, duration)`

`context.WithDeadline(parent, time)`
