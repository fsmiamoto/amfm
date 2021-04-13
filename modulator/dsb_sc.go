package modulator

import "github.com/fsmiamoto/amfm/signal"

// DSB_SC represents a AM DSB-SC modulator
type DSB_SC struct {
}

func NewDSB_SC() *DSB_SC {
	return &DSB_SC{}
}

func (d *DSB_SC) Modulate(carrier signal.Signal, message signal.Signal) signal.Signal {
	if len(carrier) != len(message) {
		return nil
	}

	result := make([]float64, len(carrier))
	for i := range result {
		result[i] = carrier[i] * message[i]
	}

	return result
}
