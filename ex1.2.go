package main

import (
	"fmt"
	"os"
)

func ex1_2() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
