package main

import "fmt"

//　DIその３:メソッド呼び出し時にDI

type ServiceImpl struct{}

func (s *ServiceImpl) Apply(di int) error { return nil }

type OrderService interface {
	Apply(int) error
}

type Application struct {
	os OrderService
}

func (app *Application) Apply(os OrderService, id int) error {
	return os.Apply(id)
}

func main() {
	app := &Application{}
	// ここでDIしている
	app.Apply(&ServiceImpl{}, 19)
}
