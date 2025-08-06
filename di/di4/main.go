package main

// Goでは構造体の中にインターフェースを埋め込むことができる
// これによって抽象的に依存した型を定義できる

type OrderService interface {
	Apply(int) error
}

type ServiceImpl struct{}

func (s *ServiceImpl) Apply(id int) error { return nil }

type Application struct {
	OrderService // 埋め込みインターフェース
}

func (app *Application) Run(id int) error {
	return app.Apply(id)
}

func main() {
	// 初期化時の宣言はオブジェクト初期化時にDIする方法と変わらない
	app := &Application{OrderService: &ServiceImpl{}}
	app.Run(19)
}
