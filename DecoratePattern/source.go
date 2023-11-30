package DecoratePattern

import (
	"fmt"
)

type Component interface {
	Operation()
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Operation() {
	fmt.Println("具体对象的操作~")
}

type Decorator struct {
	component Component
}

func (d *Decorator) Operation() {
	if d.component != nil {
		d.component.Operation()
	}
}

func (d *Decorator) SetComponent(component Component) {
	d.component = component
}
