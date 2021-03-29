package signal

import "math"

func Sin(frequency float64, numOfSamples int) Signal {
	result := make(Signal, numOfSamples)
	for i := range result {
		t := float64(i) / float64(numOfSamples)
		result[i] = math.Sin(2 * frequency * math.Pi * t)
	}
	return result
}

func Cos(frequency float64, numOfSamples int) Signal {
	result := make(Signal, numOfSamples)

	for i := range result {
		t := float64(i) / float64(numOfSamples)
		result[i] = math.Cos(2 * frequency * math.Pi * t)
	}
	return result
}
