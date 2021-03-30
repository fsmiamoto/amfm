package signal

import (
	"fmt"
	"os"
	"strings"
)

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

// CumulativeSum returns the cumulative sum of a signal as another signal
func (s Signal) CumulativeSum() Signal {
	result := make(Signal, len(s))
	acc := 0.0
	for i := range result {
		acc += s[i]
		result[i] = acc
	}
	return result
}

// WriteToFile saves the values of the signal to a human-readable formatted file.
// TODO: Add benchmark, is it better to write directly to the file or
// use a builder and then write everything 'at once'?
func (s Signal) WriteToFile(filename string) error {
	file, err := os.Create(filename)

	builder := strings.Builder{}

	for i := range s {
		builder.WriteString(fmt.Sprintf("%.4f\n", s[i]))
	}

	_, err = file.WriteString(builder.String())
	return err
}
