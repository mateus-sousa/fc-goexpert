package tax

import "errors"

type Repository interface {
	SaveTax(tax float64) error
}

func TaxCalculateAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)
	return repository.SaveTax(tax)
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0 {
		return 0
	} else if amount >= 1000 && amount < 2000 {
		return 10
	} else if amount >= 2000 {
		return 20
	}
	return 5
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("value should be greater than 0")
	} else if amount >= 1000 && amount < 2000 {
		return 10, nil
	} else if amount >= 2000 {
		return 20, nil
	}
	return 5, nil
}
