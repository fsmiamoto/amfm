package signal

type binaryOp func(a, b float64) float64
type unaryOp func(a float64) float64

func (s Signal) applyUnaryOperator(op unaryOp) Signal {
	result := make(Signal, len(s))
	for i := range result {
		result[i] = op(s[i])
	}
	return result
}

func (s Signal) applyBinaryOperator(another Signal, op binaryOp) Signal {
	result := make(Signal, len(s))
	for i := range result {
		result[i] = op(s[i], another[i])
	}
	return result
}
