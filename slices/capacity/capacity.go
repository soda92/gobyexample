package main

import (
	"fmt"
)

func Extend(slice0 []int, element int) []int {
	n := len(slice0)
	slice := slice0[0 : n+1]
	slice[n] = element
    if cap(slice) <= len(slice) {
        fmt.Println("slice is full!")
        return slice0
    }
	return slice
}

func main() {
	var iBuffer [10]int
	slice := iBuffer[0:0]
	for i := 0; i < 20; i++ {
		slice = Extend(slice, i)
		fmt.Println(slice)
	}
}
