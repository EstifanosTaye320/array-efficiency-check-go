package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("The result for slice with preallocation:", testSlice(make([]int, 0, 1000000)))
	fmt.Println("The result for slice with no preallocation:", testSlice([]int{}))
}

func testSlice(slice []int) time.Duration {
	start:= time.Now()
	for i:=0; i<1000000; i++ {
		slice = append(slice, i)
	}
	return time.Since(start)
}