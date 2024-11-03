package pzngolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, int(100), Min[int](100, 200))

	assert.Equal(t, Age(100), Min[Age](Age(100), Age(200)))
}

func TestTypeInference(t *testing.T) {
	assert.Equal(t, int(100), Min(100, 200))
	assert.Equal(t, float64(100.0), Min(100.0, 200.0))
	assert.Equal(t, Age(100), Min(Age(100), Age(200)))
}
