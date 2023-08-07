package main

import (
	"fmt"
)

func main() {
	src := []int{-3, -1, 5, 7, 11, 15}
	n := len(src)
	var res int
	if src[0] >= 0 {
		fmt.Println(src[0] + src[1])
		return
	} else if src[n-1] <= 0 {
		fmt.Println(-(src[n-1] + src[n-2]))
		return
	}
	for i := range src {
		if src[i] >= 0 {

		}
	}
	fmt.Println(res)
}
