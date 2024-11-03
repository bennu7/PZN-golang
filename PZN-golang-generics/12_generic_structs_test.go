package pzngolanggenerics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	first  T
	second T
}

func TestData(t *testing.T) {
	dataInt := Data[int]{10, 20}
	t.Log("dataInt -> ", dataInt)
	assert.Equal(t, 10, dataInt.first)

	dataString := Data[string]{
		first:  "A",
		second: "B",
	}

	t.Log("dataString -> ", dataString)
	assert.Equal(t, "A", dataString.first)
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.first = first

	return d.first
}

func TestGenericMethod(t *testing.T) {

	dataInt := Data[int]{10, 20}
	t.Log("dataInt -> ", dataInt)
	assert.Equal(t, 10, dataInt.first)

	sayHello := dataInt.SayHello("John")
	t.Log("sayHello -> ", sayHello)
	assert.Equal(t, "Hello John", sayHello)

	valChangeFirst := dataInt.ChangeFirst(125)
	t.Log("valChangeFirst -> ", valChangeFirst)
	assert.Equal(t, 125, valChangeFirst)
}
