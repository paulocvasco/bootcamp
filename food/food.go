package food

import "errors"

const (
	dog = "dog"
	cat = "cat"
	ham = "hamster"
	tar = "tarantula"
)

func animalDog(count int) float64 {
	return (float64(count) * 10.0)
}

func animalCat(count int) float64 {
	return (float64(count) * 5.0)
}

func animalHamster(count int) float64 {
	return (float64(count) * 0.25)
}

func animalTarantula(count int) float64 {
	return (float64(count) * 0.15)
}

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	case ham:
		return animalHamster, nil
	case tar:
		return animalTarantula, nil
	default:
		return nil, errors.New("Animal desconhecido")
	}
}
