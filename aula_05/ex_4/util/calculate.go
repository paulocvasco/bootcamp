package util

import "errors"

func CalculateSalary(hours int, value float64) (float64, error) {
	if hours < 80 {
		return 0, errors.New("o trabalhador nao pode trabalhar menos de 80 horas")
	}
	total := float64(hours) * value

	if total >= 20000 {
		return total * 0.8, nil
	}
	return total, nil
}
