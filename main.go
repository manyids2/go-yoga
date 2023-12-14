package main

/*
#cgo CFLAGS: -I${SRCDIR}/yoga
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/yoga
#cgo LDFLAGS: -L${SRCDIR}/yoga

#include <stdio.h>
#include <stdlib.h>
#include <yoga/Yoga.h>
*/
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println(int(C.YGAlignFlexStart))
}
