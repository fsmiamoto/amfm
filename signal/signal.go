package signal

type Signal []float64

// Multiply multiplies two signals, sample by sample, returning the resulting signal
func (s Signal) Multiply(another Signal) Signal {
	return s.applyBinaryOperator(multiply, another)
}

// Add adds two signals, sample by sample, returning the resulting signal
func (s Signal) Add(another Signal) Signal {
	return s.applyBinaryOperator(add, another)
}

// Scale scales a signal by a scalar factor
func (s Signal) Scale(factor float64) Signal {
	return s.applyUnaryOperator(func(a float64) float64 {
		return a * factor
	})
}

func (s Signal) applyUnaryOperator(op unaryOp) Signal {
	result := make(Signal, len(s))
	for i := range result {
		result[i] = op(s[i])
	}
	return result
}

func (s Signal) applyBinaryOperator(op binaryOp, another Signal) Signal {
	result := make(Signal, len(s))
	for i := range result {
		result[i] = op(s[i], another[i])
	}
	return result
}
