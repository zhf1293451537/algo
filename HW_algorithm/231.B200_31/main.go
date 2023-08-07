package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	str := "2615371"
	n := 4
	minIdx := -1
	minV := 10
	for i := 0; i < n+1; i++ {
		if int(str[i]-'0') < minV {
			minV = int(str[i] - '0')
			minIdx = i
		}
	}
	res := []byte{str[minIdx]}
	ans := math.MaxInt
	dfs(str, minIdx+1, &res, len(str)-n, &ans)
	fmt.Println(ans)
}
func dfs(str string, index int, res *[]byte, length int, ans *int) {
	if len(*res) == length {
		temp, _ := strconv.Atoi(string(*res))
		if temp < *ans {
			*ans = temp
		}
		return
	}
	for i := index; i < len(str); i++ {
		*res = append(*res, str[i])
		dfs(str, i+1, res, length, ans)
		*res = (*res)[:len(*res)-1]
	}
}
