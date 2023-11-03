package tax

import "time"

func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0
	} else if amount >= 1000 && amount < 2000 {
		return 10
	} else if amount >= 2000 {
		return 20
	}
	return 5
}

func CalculateTax2(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount == 0 {
		return 0
	} else if amount >= 1000 {
		return 10
	}
	return 5
}
