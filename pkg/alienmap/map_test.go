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
