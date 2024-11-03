package pzngolanggenerics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// inline type constraint
func FindMin[T interface{ int | int64 | float64 }](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMinInlineTypeConstraint(t *testing.T) {
	assert.Equal(t, 100, FindMin(100, 200))
	assert.Equal(t, int64(100), FindMin(int64(100), int64(200)))
	assert.Equal(t, 100.0, FindMin(100.0, 200.0))
}

// generic type parameter
func GetFirst[T []E, E any](slice T) E {
	return slice[0]
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"John",
		"Doe",
		"Ibnu",
	}
	first := GetFirst[[]string, string](names)
	assert.Equal(t, "John", first)

	assert.Equal(t, 1, GetFirst([]int{1, 2, 3}))
	assert.Equal(t, "A", GetFirst([]string{"A", "B", "C"}))
	assert.Equal(t, 1.0, GetFirst([]float64{1.0, 2.0, 3.0}))
}
