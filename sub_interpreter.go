package python3

/*
#include "Python.h"
*/
import "C"

//Py_NewInterpreter : https://docs.python.org/3/c-api/init.html#c.Py_NewInterpreter
func Py_NewInterpreter() *PyThreadState {
	return (*PyThreadState)(C.Py_NewInterpreter())
}

//Py_EndInterpreter : https://docs.python.org/3/c-api/init.html#c.Py_EndInterpreter
func Py_EndInterpreter(tstate *PyThreadState) {
	C.Py_EndInterpreter((*C.PyThreadState)(tstate))
}
