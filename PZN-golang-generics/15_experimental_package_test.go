package pzngolanggenerics

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"maps"
	"testing"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	}

	return second
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, int64(100), ExperimentalMin(int64(100), int64(200)))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 200.0))

	// using string
	assert.Equal(t, "A", ExperimentalMin("A", "B"))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"name": "John",
	}
	second := map[string]string{
		"name": "John",
	}

	assert.True(t, maps.Equal(first, second))
}
