package modulator

import (
	"errors"

	"github.com/fsmiamoto/amfm/signal"
)

type Modulator interface {
	Modulate(carrier signal.Signal, message signal.Signal) signal.Signal
}

type AM struct {
	Sensitivity float64 // Given in 1/V
}

func (am *AM) Modulate(carrier signal.Signal, message signal.Signal) (signal.Signal, error) {
	if len(carrier) != len(message) {
		return nil, errors.New("length of carrier and message must match")
	}

	result := make([]float64, len(carrier))
	for i := range result {
		result[i] = carrier[i] + carrier[i]*am.Sensitivity*message[i]
	}

	return result, nil
}
