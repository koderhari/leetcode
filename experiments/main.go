package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	aaa bool
	bbb int32
	ccc bool
	ddd bool
}

type FooBig struct {
	aaa bool
	bbb int64
	ccc bool
	ddd bool
}

type FooSmall struct {
	aaa bool
	ccc bool
	ddd bool
	bbb int32
}

func main() {
	fmt.Println(unsafe.Sizeof(Foo{}))
	fmt.Println(unsafe.Sizeof(FooBig{}))
	fmt.Println(unsafe.Sizeof(FooSmall{}))
}
