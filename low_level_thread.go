package python3

/*
#include "Python.h"
*/
import "C"

// PyInterpreterState : https://docs.python.org/3/c-api/init.html#c.PyInterpreterState
type PyInterpreterState C.PyInterpreterState

func PyInterpreterState_New() *PyInterpreterState {
	return (*PyInterpreterState)(C.PyInterpreterState_New())
}

func PyInterpreterState_Clear(interp *PyInterpreterState) {
	C.PyInterpreterState_Clear((*C.PyInterpreterState)(interp))
}

func PyInterpreterState_Delete(interp *PyInterpreterState) {
	C.PyInterpreterState_Delete((*C.PyInterpreterState)(interp))
}

func PyThreadState_New(interp *PyInterpreterState) *PyThreadState {
	return (*PyThreadState)(C.PyThreadState_New((*C.PyInterpreterState)(interp)))
}

func PyThreadState_Clear(tstate *PyThreadState) {
	C.PyThreadState_Clear((*C.PyThreadState)(tstate))
}

func PyThreadState_Delete(tstate *PyThreadState) {
	C.PyThreadState_Delete((*C.PyThreadState)(tstate))
}

func PyInterpreterState_GetID(interp *PyInterpreterState) int64 {
	return int64(C.PyInterpreterState_GetID((*C.PyInterpreterState)(interp)))
}

func PyThreadState_GetDict() *PyObject {
	return togo(C.PyThreadState_GetDict())
}

func PyThreadState_SetAsyncExc(id uint64, exc *PyObject) int {
	return int(C.PyThreadState_SetAsyncExc(C.ulong(id), toc(exc)))
}

func PyEval_AcquireThread(tstate *PyThreadState) {
	C.PyEval_AcquireThread((*C.PyThreadState)(tstate))
}

func PyEval_ReleaseThread(tstate *PyThreadState) {
	C.PyEval_ReleaseThread((*C.PyThreadState)(tstate))
}
