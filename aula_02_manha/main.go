package main

import (
	"bootcamp/aula_02_manha/food"
	"fmt"
)

func main() {

	var amount float64

	animals := []struct {
		name string
		num  int
	}{
		{"dog", 10}, {"papagaio", 3}, {"hamster", 8}, {"tarantula", 3}, {"cat", 11}, {"cavalo", 2},
	}

	for _, animal := range animals {
		quant, err := food.Animal(animal.name)
		if err != nil {
			fmt.Printf("%s: %s\n", animal.name, err.Error())
			continue
		}
		amount += quant(animal.num)
	}

	fmt.Printf("Total: %.2f\n", amount)

}
