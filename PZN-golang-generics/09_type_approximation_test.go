package pzngolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

// * dengan  menambahkan ~ di depan int, maka type Age akan dianggap sebagai type Number dengan type Approximation
type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func TestTypeApprox(t *testing.T) {

	assert.Equal(t, int(100), Min[int](100, 200))

	//di anggap tidak kompatibel karena type Age bukan merupakan type Number di dalam Min[T Number], meski di type Age terdapat int
	//assert.Equal(t, Age(100), Min[Age](Age(100), Age(200)))
}
