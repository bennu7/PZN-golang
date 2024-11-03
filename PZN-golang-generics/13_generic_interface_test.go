package pzngolanggenerics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)

	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func (m *MyData[T]) SetValue(value T) {
	m.Value = value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}
	result := ChangeValue(&myData, "hii Ibnu")
	t.Log("result -> ", result)
	assert.Equal(t, "hii Ibnu", result)
	assert.Equal(t, "hii Ibnu", myData.Value)

	dataInt := MyData[int]{Value: 10}
	t.Log("dataInt -> ", dataInt)
	assert.Equal(t, 10, dataInt.Value)

	ChangeValue(&dataInt, 20)
	t.Log("dataInt -> ", dataInt)
	assert.Equal(t, 20, dataInt.Value)
}
