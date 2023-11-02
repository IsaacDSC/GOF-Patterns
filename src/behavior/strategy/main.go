package main

import (
	"fmt"
	"strings"
)

type Calculator struct {
	operationFunction map[string]func(numb01, numb02 int) int
}

func NewCalculator() *Calculator {
	calculator := new(Calculator)
	calculator.registerOperationRules()
	return calculator
}

func (c *Calculator) Calculate(operationType string, numb01, numb02 int) int {
	function := c.operationFunction[strings.ToLower(operationType)]
	if function == nil {
		panic("Operation type not registered")
	}
	return function(numb01, numb02)
}

func (c *Calculator) registerOperationRules() {
	c.operationFunction = make(map[string]func(numb01, numb02 int) int)
	c.operationFunction["sum"] = c.sum
	c.operationFunction["sub"] = c.sub
}

func (*Calculator) sum(numb01, numb02 int) int {
	return numb01 + numb02
}

func (*Calculator) sub(numb01, numb02 int) int {
	return numb01 - numb02
}

func main() {
	calc := NewCalculator()
	result := calc.Calculate("sum", 10, 10)
	fmt.Println("Result", result)
	result = calc.Calculate("sub", 10, 10)
	fmt.Println("Result", result)
}
