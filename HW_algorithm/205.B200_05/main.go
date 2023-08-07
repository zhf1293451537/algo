package main

import (
	"fmt"
	"math"
)

func getResult(n, src, dst int, relations [][]int) int {
	const inf = math.MaxInt64 / 2
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
		for j := range graph[i] {
			graph[i][j] = inf
		}
	}
	for i := range relations {
		u, v, w := relations[i][0], relations[i][1], relations[i][2]
		graph[u-1][v-1] = w
	}
	distance := make([]int, n)
	for i := range distance {
		distance[i] = inf
	}
	distance[src-1] = 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		x := -1
		for y, u := range used {
			if !u && (x == -1 || distance[y] < distance[x]) {
				x = y
			}
		}
		used[x] = true
		for u, time := range graph[x] {
			distance[u] = min(distance[u], distance[x]+time)
		}
	}
	if distance[dst-1] == inf {
		return -1
	}
	return distance[dst-1]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	var n, m, src, dst int
	fmt.Scan(&n, &m)
	relations := make([][]int, m)
	for i := range relations {
		relations[i] = make([]int, 3)
		fmt.Scan(&relations[i][0], &relations[i][1], &relations[i][2])
	}
	fmt.Scan(&src, &dst)
	fmt.Println(getResult(n, src, dst, relations))
}
