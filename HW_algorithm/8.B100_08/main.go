package main

import (
	"fmt"
	"strconv"
)

func getResult(k int, src [][]int) int {
	res := 0
	for _, v := range src {
		temp := []byte{}
		for i := range v {
			temp = append(temp, byte(v[i]+'0'))
		}
		maxlen := 0
		for i := 0; i < len(v); i++ {
			temp1, _ := strconv.ParseInt(string(temp[i:])+string(temp[:i]), 2, 10)
			maxlen = max(maxlen, int(temp1))
		}
		res += int(maxlen)
	}
	return res
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	fmt.Println(getResult(5, [][]int{
		[]int{1, 0, 0, 0, 1},
		[]int{0, 0, 0, 1, 1},
		[]int{0, 1, 0, 1, 0},
		[]int{1, 0, 0, 1, 1},
		[]int{1, 0, 1, 0, 1},
	}))
}
