package ports

//"context"

type GRPCPorts interface {
	Run()
	GetAddition()
	GetSubtraction()
	GetMultiplication()
	GetDivision()
}
