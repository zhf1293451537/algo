package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getResult(s string, n int) {
	src := [][]int{}
	res := strings.Split(s, ",")
	for _, v := range res {
		if strings.Contains(v, "-") {
			temp := strings.Split(v, "-")
			l1, _ := strconv.Atoi(temp[0])
			l2, _ := strconv.Atoi(temp[1])
			src = append(src, []int{l1, l2})
		} else {
			temp, _ := strconv.Atoi(v)
			src = append(src, []int{temp})
		}
	}
	sort.Slice(src, func(i, j int) bool { return src[i][0] < src[j][0] })
	for i, v := range src {
		if len(v) == 2 && n >= v[0] && n <= v[1] {
			if n == v[0] {
				v[0] += 1
				if v[0] == v[1] {
					src[i] = src[i][:1]
				}
			} else if n == v[1] {
				v[1] -= 1
				if v[0] == v[1] {
					src[i] = src[i][:1]
				}
			} else {
				temp1 := []int{v[0], n - 1}
				temp2 := []int{n + 1, v[1]}
				if temp1[0] == temp1[1] {
					temp1 = temp1[:1]
				}
				if temp2[0] == temp2[1] {
					temp2 = temp2[:1]
				}
				src = append(src[:i], append([][]int{temp1, temp2}, src[i+1:]...)...)
			}
		} else if len(v) == 1 && n == v[0] {
			src = append(src[:i], src[i+1:]...)
		}
	}
	fmt.Println(src)
}
func main() {
	getResult("20-21,15,18,30,5-7", 6)
}
