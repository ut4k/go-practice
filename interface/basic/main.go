package main

import "fmt"

// *Fileオブジェクトはio.Readerインターフェースを満たしている（暗黙のインターフェース定義）
// 他の言語みたいに明示的に「File implements Reader」とは書かない
type File struct{}

func (f *File) Read(p []byte) (n int, err error) { return 0, nil }

// ---
// ある構造体にたいして特定の責務だけのみ利用したい場合のインターフェース定義
// Goの思想では、「インターフェースは使う側が定義する（必要な最小限だけに絞る）」というスタイルが推奨されている
// Receiver インターフェースはその例です。
type Modem struct{}

func (Modem) Dial()   {}
func (Modem) Hangup() {}
func (Modem) Sender() {}
func (Modem) Recv() {
	fmt.Println("Receiving via Modem")
}

type Receiver interface {
	Recv()
}

// ---
// 他にも受信という機能をもつ構造体があった場合・・・
type Phone struct{}

func (Phone) Recv() {
	fmt.Println("Receiving via Phone")
}

func receiveData(r Receiver) {
	r.Recv()
}

// --- メイン ---
func main() {
	var m Modem
	var p Phone

	receiveData(m) // => Receiving via Modem
	receiveData(p) // => Receiving via Phone
}
