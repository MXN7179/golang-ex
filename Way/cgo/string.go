package cgo

import "C"
import "unsafe"

// #include <stdio.h>
// #include <stdlib.h>
// void print(char * s){
//	printf("%s\n", s);
// }
import "C"

func Print(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
	C.print(cs)
	C.printf("%s\n", cs)
}
