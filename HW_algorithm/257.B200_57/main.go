package main

import (
	"fmt"
	"math"
)

func main() {
	n := 6
	src := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 5},
		{3, 6},
	}
	city := []int{}
	minV := math.MaxInt
	for i := 1; i < n+1; i++ {
		//init parent
		parent := make([]int, n+1)
		for d := range parent {
			parent[d] = d
		}
		//
		for _, v := range src {
			//排除当前节点
			if v[0] == i || v[1] == i {
				continue
			}
			union(parent, v[0], v[1])
		}
		count := map[int]int{}
		for x := 1; x < n+1; x++ {
			f := find(parent, x)
			count[f]++
		}
		fmt.Println(count)
		maxlen := 0
		for _, n := range count {
			if n > maxlen {
				maxlen = n
			}
		}
		if maxlen < minV {
			minV = maxlen
			city = []int{i}
		} else if maxlen == minV {
			city = append(city, i)
		}

	}
	fmt.Println(city)
}
func union(parent []int, x, y int) {
	rootX := find(parent, x)
	rootY := find(parent, y)
	if rootX == rootY {
		return
	}
	parent[rootX] = rootY
}
func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}
