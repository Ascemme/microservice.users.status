package main

import (
	"fmt"
	"runtime"
)

func main() {
	s := runtime.NumCPU()
	fmt.Println(s)
}
