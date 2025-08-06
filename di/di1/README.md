# ポイント

- Application は ServiceImpl の存在を知らなくていい（インターフェースに依存するだけ）。

- テスト用のダミー実装（モック）を渡せる

- 将来OrderService の実装を差し替えやすい

## DIじゃないとこうなる

```go
type ServiceImpl struct{}

func (s *ServiceImpl) Apply(di int) error { return nil }

type OrderService interface {
	Apply(int) error
}

type Application struct {
	os OrderService
}

func NewApplication() *Application {
	return &Application{
		os: &ServiceImpl{}, // ← ここで直接new
	}
}

func (app *Application) Apply(id int) error {
	return app.os.Apply(id)
}

func main() {
	app := NewApplication()
	app.Apply(19)
}
```

- Application は常に ServiceImpl に依存している

- ApplyをテストしたいときにもServiceImpl固定であるため変更できない

- 別の実装を使いたくてもServiceImplに依存してしまっているため、Applicationを直接変更しないといけない
