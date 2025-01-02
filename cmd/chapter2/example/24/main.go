package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("number of CPUs:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}
