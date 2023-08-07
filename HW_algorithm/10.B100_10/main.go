package main

import (
	"fmt"
	"sort"
)

func getResult(s1, s2 string) string {
	set1 := make(map[byte]int)
	set2 := make(map[byte]int)
	res := []byte{}
	for i := range s1 {
		set1[s1[i]]++
	}
	for i := range s2 {
		set2[s2[i]]++
	}
	for i := range set1 {
		if _, ok := set2[i]; !ok {
			delete(set1, i)
		} else {
			res = append(res, i)
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return string(res)
}
func main() {
	fmt.Println(getResult("fach", "bbaaccedfg"))
}
