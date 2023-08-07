package main

import (
	"fmt"
	"math"
)

func getResult(n, edge int, relations [][]int, src int) int {
	const inf = math.MaxInt64 / 2
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
		for j := range graph[i] {
			graph[i][j] = inf
		}
	}

	for _, v := range relations {
		graph[v[0]-1][v[1]-1] = 1
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[src-1] = 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		x := -1
		for y, u := range used {
			if !u && (x == -1 || dist[y] < dist[x]) {
				x = y
			}
		}
		used[x] = true
		for y, time := range graph[x] {
			dist[y] = min(dist[y], dist[x]+time)
		}
	}
	res := 0
	for i := range dist {
		if dist[i] == inf {
			return -1
		}
		res = max(dist[i], res)
	}
	return res * 2
}
func main() {
	fmt.Println(getResult(5, 7, [][]int{
		{1, 4},
		{2, 1},
		{2, 3},
		{2, 4},
		{3, 4},
		{3, 5},
		{4, 5},
	}, 2))
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func networkDelayTime(times [][]int, n, k int) (ans int) {
//     const inf = math.MaxInt64 / 2
//     g := make([][]int, n)
//     for i := range g {
//         g[i] = make([]int, n)
//         for j := range g[i] {
//             g[i][j] = inf
//         }
//     }
//     for _, t := range times {
//         x, y := t[0]-1, t[1]-1
//         g[x][y] = t[2]
//     }

//     dist := make([]int, n)
//     for i := range dist {
//         dist[i] = inf
//     }
//     dist[k-1] = 0
//     used := make([]bool, n)
//     for i := 0; i < n; i++ {
//         x := -1
//         for y, u := range used {
//             if !u && (x == -1 || dist[y] < dist[x]) {
//                 x = y
//             }
//         }
//         used[x] = true
//         for y, time := range g[x] {
//             dist[y] = min(dist[y], dist[x]+time)
//         }
//     }

//     for _, d := range dist {
//         if d == inf {
//             return -1
//         }
//         ans = max(ans, d)
//     }
//     return
// }

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }
