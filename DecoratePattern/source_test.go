package DecoratePattern

import (
	"fmt"
	"testing"
)

type ConcreteDecoratorA struct {
	base Decorator
}

func (c *ConcreteDecoratorA) Operation() {
	c.base.Operation()
	fmt.Println("具体装饰对象A的操作")
}

type ConcreteDecoratorB struct {
	base Decorator
}

func (c *ConcreteDecoratorB) Operation() {
	c.base.Operation()
	fmt.Println("具体装饰对象B的操作")
}

func (c *ConcreteDecoratorB) AddedBehavior() {

}

func TestDecorator_Operation(t *testing.T) {
	c := &ConcreteComponent{}
	d1 := &ConcreteDecoratorA{}
	d2 := &ConcreteDecoratorB{}

	d1.base.SetComponent(c)
	d2.base.SetComponent(d1)
	d2.Operation()
}
