package fft

import (
	"math/cmplx"

	"github.com/fsmiamoto/amfm/signal"
	dsp "github.com/mjibson/go-dsp/fft"
)

func ForSignal(s signal.Signal) []float64 {
	result := dsp.FFTReal(s)

	abs := make([]float64, len(result))

	for i := range abs {
		abs[i] = cmplx.Abs(result[i])
	}

	return abs
}
