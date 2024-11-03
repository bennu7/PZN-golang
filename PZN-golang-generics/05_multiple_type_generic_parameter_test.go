package pzngolanggenerics

import "testing"

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	return param1, param2
}

func TestMultipleTypeParameter(t *testing.T) {
	var result, resultSecond = MultipleParameter[string, int]("Hello", 10)
	t.Log(result)
	t.Log(resultSecond)

	var result2, resultSecond2 = MultipleParameter[int, string](10, "Hello")
	t.Log(result2)
	t.Log(resultSecond2)
}
