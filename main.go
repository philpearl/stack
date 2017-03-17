package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	f1(0xdeadbeef)
}

func f1(val int) {
	f2(0xabad1dea)
}

func f2(val int) {
	f3(0xbaddcafe)
}

func f3(val int) {
	ptr := uintptr(unsafe.Pointer(&val))

	mem := *(*[20]uintptr)(unsafe.Pointer(&val))
	for i, x := range mem {
		fmt.Printf("%X: %X\n", ptr+uintptr(i*8), x)
	}

	showFunc(mem[2])
	showFunc(mem[5])
	showFunc(mem[8])
	showFunc(mem[19])
}

func showFunc(at uintptr) {
	if f := runtime.FuncForPC(at); f != nil {
		file, line := f.FileLine(at)
		fmt.Printf("%X is %s %s %d\n", at, f.Name(), file, line)
	}
}
