package modulator

import "github.com/fsmiamoto/amfm/signal"

// DSB represents a AM-DSB modulator
type DSB struct {
	Sensitivity float64 // Given in 1/V
}

func NewDSB(sensitivity float64) *DSB {
	return &DSB{
		sensitivity,
	}
}

// Modulate applies AM-DSB modulation to the carrier with a given message signal
func (dsb *DSB) Modulate(carrier signal.Signal, message signal.Signal) signal.Signal {
	if len(carrier) != len(message) {
		return nil
	}

	result := make([]float64, len(carrier))
	for i := range result {
		result[i] = carrier[i] + carrier[i]*dsb.Sensitivity*message[i]
	}

	return result
}
