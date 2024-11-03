package pzngolanggenerics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func IsSame[T comparable](Value1, Value2 T) bool {
	return Value1 == Value2
}

func TestIsSame(t *testing.T) {
	var result = IsSame[int](10, 10)
	t.Log(result)

	var result2 = IsSame[string]("Hello", "Hello")
	t.Log(result2)

	var result3 = IsSame[int](10, 20)
	t.Log(result3)

	var result4 = IsSame[string]("Hello", "World")
	t.Log(result4)

	assert.True(t, result)
}
