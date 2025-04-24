package calculator

import "testing"

func TestAdd(t *testing.T) {
    calc := &CalculatorImpl{}
    result := calc.Add(1, 2)
    if result != 3 {
        t.Errorf("Expected 3, got %v", result)
    }
}
