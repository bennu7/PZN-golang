package pzngolanggenerics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//type Number interface {
//	int | int8 | int16 | int32 | int64 | float32 | float64
//}

func Min[T Number](first T, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestTypeSet(t *testing.T) {
	var result = Min[int](10, 20)
	t.Log("result -> ", result)

	var resultInt64 = Min[int64](100, 200)
	t.Log("result -> ", resultInt64)

	var result2 = Min[float32](10.5, 20.5)
	t.Log("result2 -> ", result2)

	assert.Equal(t, 10, result)
	assert.Equal(t, int64(100), resultInt64)
	assert.Equal(t, float32(10.5), result2)
}
