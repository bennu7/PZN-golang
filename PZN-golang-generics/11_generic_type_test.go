package pzngolanggenerics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		println(value)
	}
}

func TestGenericTypeIntBag(t *testing.T) {
	numberBag := Bag[int]{1, 2, 3, 4}

	numberBag = append(numberBag, 1)
	numberBag = append(numberBag, 2)
	numberBag = append(numberBag, 3)
	numberBag = append(numberBag, 4)

	PrintBag[int](numberBag)
	assert.Equal(t, 8, len(numberBag))
}

func TestGenericTypeStringBag(t *testing.T) {
	stringBag := Bag[string]{"A", "B", "C", "D"}

	stringBag = append(stringBag, "AB")

	PrintBag(stringBag)
	assert.Equal(t, 5, len(stringBag))
}
