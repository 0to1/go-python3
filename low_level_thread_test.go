package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInterpreterState(t *testing.T) {
	Py_Initialize()

	s := PyInterpreterState_New()

	assert.NotNil(t, s)

	PyInterpreterState_Clear(s)

	PyInterpreterState_Delete(s)

	Py_Finalize()
}
