package main

import "fmt"

//　DIその１:オブジェクト初期化時にDIする方法

type ServiceImpl struct{}

func (s *ServiceImpl) Apply(di int) error { return nil }

type OrderService interface {
	Apply(int) error
}

type Application struct {
	os OrderService
}

func NewApplication(os OrderService) *Application {
	return &Application{os: os}
}

func (app *Application) Apply(id int) error {
	return app.os.Apply(id)
}

func main() {
	// ここでDIしている
	app := NewApplication(&ServiceImpl{})
	app.Apply(19)

	fmt.Printf("%#v\n", app.os)
}
