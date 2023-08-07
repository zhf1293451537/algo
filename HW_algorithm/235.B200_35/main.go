package main

import "fmt"

func main() {
	x, y := 6, 4
	n := 5
	relation := [][]int{
		{0, 2},
		{1, 2},
		{2, 2},
		{4, 1},
		{5, 1},
	}
	fmt.Println(x, y, n, relation)
	fmt.Println(2, 3)
	src := make([][]int, x)
	for i := range src {
		src[i] = make([]int, y)
	}
	for i := range relation {
		src[relation[i][0]][relation[i][1]] = 1
	}
	src[x-1][y-1] = 2
	dfs(src, 0, 0)
	fmt.Println(src)
}
func dfs(src [][]int, idxX, idxY int) bool {
	if idxX >= 6 || idxY >= 4 {
		return false
	}
	if src[idxX][idxY] == 1 {
		return false
	}
	if src[idxX][idxY] == -1 {
		return false
	}
	if src[idxX][idxY] == 2 {
		return true
	}
	if src[idxX][idxY] == 0 {
		east := dfs(src, idxX+1, idxY)
		narth := dfs(src, idxX, idxY+1)
		if east || narth {
			src[idxX][idxY] = 2
		} else {
			src[idxX][idxY] = -1
		}
	}
	return src[idxX][idxY] == 2
}
