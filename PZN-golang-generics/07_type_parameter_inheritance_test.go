package pzngolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return "Manager " + m.Name
}

func (m *MyManager) GetManagerName() string {
	return "My Manager " + m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (vp *MyVicePresident) GetName() string {
	return "Vice President " + vp.Name
}

func (vp *MyVicePresident) GetVicePresidentName() string {
	return "GetVicePresidentName " + vp.Name
}

func TestTypeParameter(t *testing.T) {
	managerName := &MyManager{
		Name: "John",
	}
	assert.Equal(t, "Manager John", GetName[Manager](managerName))

	vicePresidentName := &MyVicePresident{
		Name: "Doe",
	}
	assert.Equal(t, "Vice President Doe", GetName[VicePresident](vicePresidentName))
}
