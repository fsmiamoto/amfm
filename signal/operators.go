package signal

type binaryOp func(a, b float64) float64
type unaryOp func(a float64) float64

func add(a, b float64) float64 {
	return a + b
}

func multiply(a, b float64) float64 {
	return a * b
}
