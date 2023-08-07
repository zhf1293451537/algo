package main

import "fmt"

func main() {
	res := []int{5}
	fmt.Println(res)
	fix(res)
	fmt.Println(res)
	fix1(res)
	fmt.Println(res)
	///////////////////////
	res1 := make([]int, 5)
	fmt.Println(res1)
	fix(res1)
	fmt.Println(res1)
	fix1(res1)
	fmt.Println(res1)
	// var n, target int
	// fmt.Scan(&n)
	// src := make([][]int, n)
	// for i := range src {
	// 	src[i] = make([]int, 2)
	// 	fmt.Scan(&src[i][0], &src[i][1])
	// }
	// fmt.Scan(&target)
	// tree := make(map[int][]int)
	// for _, v := range src {
	// 	tree[v[1]] = append(tree[v[1]], v[0])
	// }
	// if target == 0 {
	// 	fmt.Println("")
	// 	return
	// }
	// res := []int{}
	// dfs(tree, 0, target, &res)
	// sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	// fmt.Println(res)
}
func fix(res []int) {
	res[0] = 1
}
func fix1(res []int) {
	res = append(res, 10)
}

func dfs(tree map[int][]int, index, target int, res *[]int) {
	if _, ok := tree[index]; ok {
		children := tree[index]
		for _, child := range children {
			if child != target {
				*res = append(*res, child)
				dfs(tree, child, target, res)
			}
		}
	}
}
