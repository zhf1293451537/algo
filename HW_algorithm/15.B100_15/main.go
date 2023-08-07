package main

import (
	"fmt"
	"sort"
)

func getResult(s string, k int) int {
	src := []byte(s)
	var res byte
	sort.Slice(src, func(i, j int) bool { return src[i] < src[j] })
	if k > len(src) {
		res = src[len(src)-1]
	} else {
		res = src[k-1]
	}
	for i := 0; i < len(s); i++ {
		if s[i] == res {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(getResult("fAdDAkBbBq", 4))
}
