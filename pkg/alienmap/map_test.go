package alienmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitMap(t *testing.T) {
	_, err := helper()
	assert.NoError(t, err)
}

func TestPrintMap(t *testing.T) {
	m, err := helper()
	assert.NoError(t, err)
	m.PrintMap()
}

func TestRemoveCity(t *testing.T) {
	m, err := helper()
	assert.NoError(t, err)
	err = m.RemoveCity("Baz")
	assert.NoError(t, err)
}
func TestFirstWave(t *testing.T) {
	m, err := helper()
	assert.NoError(t, err)
	m.FirstWave(5)
}

func helper() (Map, error) {
	var m Map
	c := Config{
		Filepath:  "test.txt",
		NumAliens: 5,
		NumSteps:  10,
	}
	err := m.InitMap(&c)
	return m, err
}
