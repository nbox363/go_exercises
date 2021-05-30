package main

import (
	"fmt"
	"os"
)

func ex1_1() {
	var s, sep string
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
