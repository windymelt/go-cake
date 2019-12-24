package main

import (
	"github.com/windymelt/go-cake/impl"
	"github.com/windymelt/go-cake/dummy"
	"fmt"
)

type App struct {
	*impl.RealDoAComponent
}
type App2 struct {
	*dummy.DummyDoAComponent
}

func main() {
	app := &App{}
	app.FieldDoA = &(impl.RealDoA{})
	fmt.Println(app.FieldDoA.A())

	app2 := &App2{}
	app2.FieldDoA = &(dummy.DummyDoA{})
	fmt.Println(app2.FieldDoA.A())
}
