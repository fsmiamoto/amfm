package modulator

import (
	"github.com/fsmiamoto/amfm/signal"
)

// Modulator is an interface for a signal modulator
type Modulator interface {
	Modulate(carrier signal.Signal, message signal.Signal) signal.Signal
}
