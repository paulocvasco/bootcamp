package main

import (
	"bootcamp/aula_02_tarde/ex_2/types"
	"fmt"
)

func main() {

	var store types.Store

	type productData struct {
		size  string
		name  string
		price float64
	}

	products := []productData{
		{size: types.Small, name: "A", price: 89.34},
		{size: types.Medium, name: "B", price: 548.23},
		{size: types.Big, name: "C", price: 4390.82},
	}

	for _, product := range products {
		ok := store.Add(types.NewProduct(product.size, product.name, product.price))
		if !ok {
			fmt.Printf("Failed to add %s product", product.name)
		}
	}

	fmt.Printf("Total cost: %.2f\n", store.Total())
}
