package signal

type Signal []float64

// Multiply multiplies two signals, sample by sample, returning the resulting signal
func (s Signal) Multiply(another Signal) Signal {
	return s.applyBinaryOperator(another, func(a, b float64) float64 {
		return a * b
	})
}

// Add adds two signals, sample by sample, returning the resulting signal
func (s Signal) Add(another Signal) Signal {
	return s.applyBinaryOperator(another, func(a, b float64) float64 {
		return a + b
	})
}

// Scale scales a signal by a scalar factor
func (s Signal) Scale(factor float64) Signal {
	return s.applyUnaryOperator(func(a float64) float64 {
		return a * factor
	})
}
