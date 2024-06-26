/*
Unless explicitly stated otherwise all files in this repository are licensed
under the MIT License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

//go:generate go run script/variadic.go

/*
#include "Python.h"
#include "frameobject.h"
*/
import "C"

//togo converts a *C.PyObject to a *PyObject
func togo(cobject *C.PyObject) *PyObject {
	return (*PyObject)(cobject)
}

func toc(object *PyObject) *C.PyObject {
	return (*C.PyObject)(object)
}

//framectogo converts a *C.PyFrameObject to a *PyFrameObject
func framectogo(cobject *C.PyFrameObject) *PyFrameObject {
	return (*PyFrameObject)(cobject)
}

func framegotoc(object *PyFrameObject) *C.PyFrameObject {
	return (*C.PyFrameObject)(object)
}
