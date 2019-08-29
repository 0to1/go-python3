package python3

/*
#include "Python.h"
*/
import "C"

// PyInterpreterState : https://docs.python.org/3/c-api/init.html#c.PyInterpreterState
type PyInterpreterState C.PyInterpreterState

// PyInterpreterState_New : https://docs.python.org/3/c-api/init.html#c.PyInterpreterState_New
func PyInterpreterState_New() *PyInterpreterState {
	return (*PyInterpreterState)(C.PyInterpreterState_New())
}
// PyInterpreterState_Clear : https://docs.python.org/3/c-api/init.html#c.PyInterpreterState_Clear
func PyInterpreterState_Clear(interp *PyInterpreterState) {
	C.PyInterpreterState_Clear((*C.PyInterpreterState)(interp))
}
// PyInterpreterState_Delete : https://docs.python.org/3/c-api/init.html#c.PyInterpreterState_Delete
func PyInterpreterState_Delete(interp *PyInterpreterState) {
	C.PyInterpreterState_Delete((*C.PyInterpreterState)(interp))
}
// PyThreadState_New : https://docs.python.org/3/c-api/init.html#c.PyThreadState_New
func PyThreadState_New(interp *PyInterpreterState) *PyThreadState {
	return (*PyThreadState)(C.PyThreadState_New((*C.PyInterpreterState)(interp)))
}
// PyThreadState_Clear : https://docs.python.org/3/c-api/init.html#c.PyThreadState_Clear
func PyThreadState_Clear(tstate *PyThreadState) {
	C.PyThreadState_Clear((*C.PyThreadState)(tstate))
}
// PyThreadState_Delete : https://docs.python.org/3/c-api/init.html#c.PyThreadState_Delete
func PyThreadState_Delete(tstate *PyThreadState) {
	C.PyThreadState_Delete((*C.PyThreadState)(tstate))
}
// PyInterpreterState_GetID : https://docs.python.org/3/c-api/init.html#c.PyInterpreterState_GetID
func PyInterpreterState_GetID(interp *PyInterpreterState) int64 {
	return int64(C.PyInterpreterState_GetID((*C.PyInterpreterState)(interp)))
}
// PyThreadState_GetDict : https://docs.python.org/3/c-api/init.html#c.PyThreadState_GetDict
func PyThreadState_GetDict() *PyObject {
	return togo(C.PyThreadState_GetDict())
}
// PyThreadState_SetAsyncExc : https://docs.python.org/3/c-api/init.html#c.PyThreadState_SetAsyncExc
func PyThreadState_SetAsyncExc(id uint64, exc *PyObject) int {
	return int(C.PyThreadState_SetAsyncExc(C.ulong(id), toc(exc)))
}
// PyEval_AcquireThread : https://docs.python.org/3/c-api/init.html#c.PyEval_AcquireThread
func PyEval_AcquireThread(tstate *PyThreadState) {
	C.PyEval_AcquireThread((*C.PyThreadState)(tstate))
}
// PyEval_ReleaseThread : https://docs.python.org/3/c-api/init.html#c.PyEval_ReleaseThread
func PyEval_ReleaseThread(tstate *PyThreadState) {
	C.PyEval_ReleaseThread((*C.PyThreadState)(tstate))
}
