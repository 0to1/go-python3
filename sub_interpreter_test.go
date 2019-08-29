package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInterpreter(t *testing.T) {
	Py_Initialize()

	s := Py_NewInterpreter()

	assert.NotNil(t, s)

	//Py_EndInterpreter(s)

	// Py_Finalize()
}
