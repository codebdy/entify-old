package utils

import (
	"fmt"
	"log"
	"runtime/debug"
)

func PrintErrorStack() {
	if x := recover(); x != nil {
		println(fmt.Sprintf("%T: %+v", x, x))
		log.Printf("%s\n", debug.Stack())
		panic(x)
	}
}
