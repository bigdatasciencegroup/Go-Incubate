package sharedRPC

type MyFloats struct {
	A1, A2 float64
}

type MyInterface interface {
	Multiply112222(arguments *MyFloats, reply *float64) error
	Power112222(arguments *MyFloats, reply *float64) error
}
