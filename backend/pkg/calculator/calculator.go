package calculator

import "errors"

// Calculator 接口定义
type Calculator interface {
    Add(a, b float64) float64
    Subtract(a, b float64) float64
    Multiply(a, b float64) float64
    Divide(a, b float64) (float64, error)
}

// CalculatorImpl 计算器实现
type CalculatorImpl struct{}

func (c *CalculatorImpl) Add(a, b float64) float64 {
    return a + b
}

func (c *CalculatorImpl) Subtract(a, b float64) float64 {
    return a - b
}

func (c *CalculatorImpl) Multiply(a, b float64) float64 {
    return a * b
}

func (c *CalculatorImpl) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
