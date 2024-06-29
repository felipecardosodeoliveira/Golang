package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result := CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected %f but got %f ", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{amount: 500.0, expected: 5.0},
		{amount: 690.0, expected: 5.0},
		{amount: 1200.0, expected: 10.0},
		{amount: 0.0, expected: 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but got %f ", item.expected, result)
		}
	}
}

func BanchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}
