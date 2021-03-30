package signal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScale(t *testing.T) {
	t.Run("scale by 0", func(t *testing.T) {
		s := Signal{1, 2, 3, 4}
		assert.Equal(t, s.Scale(0), Signal{0, 0, 0, 0})
	})

	t.Run("scale by 1", func(t *testing.T) {
		s := Signal{1, 2, 3, 4}
		assert.Equal(t, s.Scale(1), s)
	})

	t.Run("double the signal", func(t *testing.T) {
		s := Signal{1, 2, 3, 4}
		assert.Equal(t, s.Scale(2), Signal{2, 4, 6, 8})
	})
}

func TestAdd(t *testing.T) {
	t.Run("add two small signals", func(t *testing.T) {
		a := Signal{1, 2, 3}
		b := Signal{2, 3, 4}

		assert.Equal(t, a.Add(b), Signal{3, 5, 7})
	})
}

func TestMultiply(t *testing.T) {
	t.Run("multiply by itself", func(t *testing.T) {
		s := Signal{2, 2, 2}
		assert.Equal(t, s.Multiply(s), Signal{4, 4, 4})
	})
}

func BenchmarkMultiply(b *testing.B) {
	s := Signal{1, 2, 3}
	for n := 0; n < b.N; n++ {
		s.Multiply(s)
	}
}

func BenchmarkAdd(b *testing.B) {
	s := Signal{1, 2, 3}
	for n := 0; n < b.N; n++ {
		s.Add(s)
	}
}

func BenchmarkScale(b *testing.B) {
	s := Signal{1, 2, 3}
	for n := 0; n < b.N; n++ {
		s.Scale(2)
	}
}
