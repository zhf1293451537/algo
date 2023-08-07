package main

import (
	"fmt"
	"sort"
)

type instance struct {
	count  int
	letter byte
}

func getResult(s string, k int) int {
	flag := s[0]
	count := 0
	src := make(map[byte]int)
	for i := range s {
		if s[i] == flag {
			count++
		} else {
			if _, ok := src[flag]; ok {
				src[flag] = max(src[flag], count)
			} else {
				src[flag] = count
			}
			count = 1
			flag = s[i]
		}
	}
	if _, ok := src[flag]; ok {
		src[flag] = max(src[flag], count)
	} else {
		src[flag] = count
	}
	if k > len(src) {
		return -1
	}
	res := []instance{}
	for i, v := range src {
		res = append(res, instance{v, i})
	}
	sort.Slice(res, func(i, j int) bool { return res[i].count > res[i].count })
	return res[k-1].count
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	fmt.Println(getResult("ABC", 2))
}
