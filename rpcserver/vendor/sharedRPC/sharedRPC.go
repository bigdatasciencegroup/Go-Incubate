package sharedRPC

type MyFloats struct {
	A1, A2 float64
}

type MyInterface interface {
	Multiply2222(arguments *MyFloats, reply *float64) error
	Power2222(arguments *MyFloats, reply *float64) error
}
