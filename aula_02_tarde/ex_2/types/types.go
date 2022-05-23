package types

import "fmt"

const (
	Small  = "small"
	Medium = "medium"
	Big    = "big"
)

type product struct {
	Type  string
	Name  string
	Price float64
}

type Store struct {
	products []product
}

type Product interface {
	Cost() float64
}

type Ecommerce interface {
	Add(Product) bool
	Total() float64
}

func (p product) Cost() float64 {
	switch p.Type {
	case Small:
		return p.Price
	case Medium:
		return p.Price * 1.03
	case Big:
		return p.Price*1.06 + 2500
	default:
		fmt.Printf("%s product invalid", p.Name)
		return 0
	}
}

func (s *Store) Add(newProduct Product) bool {
	p, ok := newProduct.(product)
	if !ok {
		return false
	}

	s.products = append(s.products, p)

	return true
}

func (s Store) Total() float64 {
	var total float64
	for _, p := range s.products {
		total += p.Cost()
	}

	return total
}

func NewProduct(productType string, name string, price float64) Product {
	newProduct := product{
		Type:  productType,
		Name:  name,
		Price: price,
	}

	return newProduct
}
