package main

import (
	"fmt"
	"sort"
	"strings"
)

func getResult(n int, s []string) []string {
	cache := map[string]bool{}
	ans := []string{}
	path := []string{}
	visit := make([]bool, n)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			temp := strings.Join(path, "")
			if !cache[temp] {
				ans = append(ans, temp)
				cache[temp] = true
			}
		}
		for i := 0; i < n; i++ {
			if !visit[i] {
				path = append(path, s[i])
				visit[i] = true
				dfs(idx + 1)
				visit[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func main() {
	s := []string{"a", "b", "a"}
	s1 := getResult(3, s)
	sort.Strings(s1)
	fmt.Println(s1)
}
