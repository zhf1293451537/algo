package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getResult(src []int) []int {
	temp := make([][]int, len(src))
	for i := range temp {
		temp[i] = make([]int, 2)
		temp[i][0], temp[i][1] = src[i], i
	}
	sort.Slice(temp, func(i, j int) bool { return temp[i][0] > temp[j][0] })
	res := make([]int, len(src))
	for i := range temp {
		res[temp[i][1]] = i
	}
	return res
}

func main() {
	var temp string
	fmt.Scan(&temp)
	temp1 := strings.Split(temp, ",")
	src := make([]int, len(temp1))
	for i := range src {
		dig, _ := strconv.Atoi(temp1[i])
		src[i] = dig
	}
	fmt.Println(getResult(src))
}
