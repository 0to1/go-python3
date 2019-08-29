package python3

/*
#include "Python.h"
*/
import "C"

//PyInterpreterState_Head : https://docs.python.org/3/c-api/init.html#c.PyInterpreterState_Head
func PyInterpreterState_Head() *PyInterpreterState {
	return (*PyInterpreterState)(C.PyInterpreterState_Head())
}

//PyInterpreterState_Main : https://docs.python.org/3/c-api/init.html#c.PyInterpreterState_Main
func PyInterpreterState_Main() *PyInterpreterState {
	return (*PyInterpreterState)(C.PyInterpreterState_Main())
}

//PyInterpreterState_Next : https://docs.python.org/3/c-api/init.html#c.PyInterpreterState_Next
func PyInterpreterState_Next(interp *PyInterpreterState) *PyInterpreterState {
	return (*PyInterpreterState)(C.PyInterpreterState_Next((*C.PyInterpreterState)(interp)))
}

//PyInterpreterState_ThreadHead : https://docs.python.org/3/c-api/init.html#c.PyInterpreterState_ThreadHead
func PyInterpreterState_ThreadHead(interp *PyInterpreterState) *PyThreadState {
	return (*PyThreadState)(C.PyInterpreterState_ThreadHead((*C.PyInterpreterState)(interp)))
}

//PyThreadState_Next : https://docs.python.org/3/c-api/init.html#c.PyThreadState_Next
func PyThreadState_Next(tstate *PyThreadState) *PyThreadState {
	return (*PyThreadState)(C.PyThreadState_Next((*C.PyThreadState)(tstate)))
}
