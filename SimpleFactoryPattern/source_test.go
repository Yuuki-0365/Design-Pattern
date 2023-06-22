package SimpleFactoryPattern

import (
	"testing"
)

func TestFactory_NewProduct(t *testing.T) {
	factory := new(Factory)
	product := factory.NewProduct("A")
	product.Use()
	product = factory.NewProduct("B")
	product.Use()
}
