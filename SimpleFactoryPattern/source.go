package SimpleFactoryPattern

import "fmt"

type Factory struct {
}

type Product interface {
	Use()
}

type ProductA struct {
}

func (p *ProductA) Use() {
	fmt.Printf("product a use\n")
}

type ProductB struct {
}

func (p *ProductB) Use() {
	fmt.Printf("product b use\n")
}

func (f *Factory) NewProduct(param string) Product {
	var product Product = nil
	if param == "A" {
		product = &ProductA{}
	} else if param == "B" {
		product = &ProductB{}
	}
	return product
}
