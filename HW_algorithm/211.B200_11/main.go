package main

import "fmt"

func Find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = Find(parent, parent[x])
	}
	return parent[x]
}
func Union(parent []int, x, y int) {
	rootX, rootY := Find(parent, x), Find(parent, y)
	if rootX == rootY {
		return
	}
	parent[rootX] = rootY
}
func main() {
	s1, s2 := "14345", "rybbr"
	n := len(s1)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	for i := range s1 {
		for j := i + 1; j < n; j++ {
			if s1[i] == s1[j] || s2[i] == s2[j] {
				Union(parent, i, j)
			}
		}
	}
	res := map[int]int{}
	for i := range s1 {
		temp := Find(parent, i)
		res[temp]++
	}
	ans := 0
	for i := range res {
		ans = max(ans, res[i])
	}
	fmt.Println(ans)
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
