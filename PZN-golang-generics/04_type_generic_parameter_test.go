package pzngolanggenerics

import (
	"fmt"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println("value -> ", param)
	return param
}

func TestSample(t *testing.T) {
	var result = Length[string]("Hello")
	fmt.Println("result -> ", result)

	var result2 = Length[int](10)
	fmt.Println("result2 -> ", result2)
}
