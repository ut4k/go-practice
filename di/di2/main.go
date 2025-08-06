package main

import "fmt"

//　DIその２:setterでDI

type ServiceImpl struct{}

func (s *ServiceImpl) Apply(di int) error { return nil }

type OrderService interface {
	Apply(int) error
}

type Application struct {
	os OrderService
}

func (app *Application) Apply(id int) error {
	return app.os.Apply(id)
}

func (app *Application) SetService(os OrderService) {
	app.os = os
}

func main() {
	app := &Application{}
	// ここでDIしている
	app.SetService(&ServiceImpl{})
	app.Apply(19)

	fmt.Printf("%#v\n", app.os)
}
