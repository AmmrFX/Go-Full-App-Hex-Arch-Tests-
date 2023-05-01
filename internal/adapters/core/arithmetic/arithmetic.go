package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}
func (arithm Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}
func (arithm Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}
func (arithm Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}
func (arithm Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
