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
	_, exists := m.cities["Baz"]
	assert.Equal(t, exists, false)
}

func TestAddCity(t *testing.T) {
	m, err := helper()
	assert.NoError(t, err)
	newCity := City{"Paris"}
	err = m.AddCity(&newCity)
	assert.NoError(t, err)
	_, exists := m.cities["Paris"]
	assert.Equal(t, exists, true)
}

func TestHasNeighbors(t *testing.T) {
	m, err := helper()
	assert.NoError(t, err)
	city := m.cities["Foo"]
	assert.Equal(t, m.hasNeighbors(city), true)
}

func TestMakeAliens(t *testing.T) {
	assert.Equal(t, len(makeAliens(1, 1000)), 1000)
}

func TestFirstWave(t *testing.T) {
	m, err := helper()
	assert.NoError(t, err)
	err = m.FirstWave(5)
	assert.NoError(t, err)
}

func TestSimulate(t *testing.T) {
	m, err := helper()
	assert.NoError(t, err)
	err = m.Simulate(5, 10)
	assert.NoError(t, err)
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
