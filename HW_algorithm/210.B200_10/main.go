package main

import (
	"fmt"
	"math"
)

func main() {
	var s string
	fmt.Scan(&s)
	src := []int{}
	for i := range s {
		if s[i] == ',' {
			continue
		}
		src = append(src, int(s[i]-'0'))
	}
	total := len(src)
	n := int(math.Sqrt(float64(total)))
	matrix := make([][]int, n)
	queue := [][]int{}
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = src[i*n+j]
			if matrix[i][j] == 1 {
				queue = append(queue, []int{i, j})
				total -= 1
			}
		}
	}
	if len(queue) == 0 || len(queue) == len(src) {
		fmt.Println(-1)
		return
	}
	offsets := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 && total > 0 {
		newQueue := [][]int{}
		for _, v := range queue {
			i, j := v[0], v[1]
			num := matrix[i][j]
			for _, u := range offsets {
				newI := i + u[0]
				newJ := j + u[1]
				if newI >= 0 && newI < n && newJ >= 0 && newJ < n && matrix[newI][newJ] == 0 {
					matrix[newI][newJ] = num + 1
					newQueue = append(newQueue, []int{newI, newJ})
					fmt.Println(newQueue)
					total -= 1
					if total == 0 {
						fmt.Println(num)
						return
					}
				}
			}
		}
		queue = newQueue
	}
}
