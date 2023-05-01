package api

import (
	"hex/internal/ports"
)

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}
func (apiA Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apiA.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(answer, "Addition")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
func (apiA Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apiA.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(answer, "Subtraction")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
func (apiA Adapter) getMultiplication(a, b int32) (int32, error) {
	answer, err := apiA.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(answer, "Multiplication")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
func (apiA Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apiA.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(answer, "Division")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
