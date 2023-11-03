package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expectedtax := 5.0

	actualTax := CalculateTax(amount)

	if actualTax != expectedtax {
		t.Errorf("Expected value %v doesn't match with actual value %v", expectedtax, actualTax)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type Amount struct {
		amount float64
		expect float64
	}

	table := []Amount{
		{500, 5.0},
		{999, 5.0},
		{1000, 10.0},
		{1500, 10.0},
		{0, 0},
	}
	for _, item := range table {
		actualTax := CalculateTax(item.amount)

		if actualTax != item.expect {
			t.Errorf("Expected value %v doesn't match with actual value %v", item.expect, actualTax)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(1000)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-10, -5, 500, 1000, 1400}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Expected value 0 doesn't match with actual value %v", result)
		}
		if amount >= 2000 && result != 20 {
			t.Errorf("Expected value 20 doesn't match with actual value %v", result)
		}
	})
}
