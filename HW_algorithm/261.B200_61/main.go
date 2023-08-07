package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	names := strings.Split(name, ",")
	var abbr string
	fmt.Scan(&abbr)
	getResult1(names, abbr)
}
func getResult(names []string, abbr string) {
	ans := []string{}
	for _, name := range names {
		parts := strings.Split(name, " ")
		if len(parts) > len(abbr) {
			continue
		}
		res := dfs(parts, 0, abbr, 0)
		if res {
			ans = append(ans, name)
		}
	}
	fmt.Println(strings.Join(ans, ","))
}

func dfs(parts []string, index int, abbr string, start int) bool {
	if start >= len(abbr) {
		return index >= len(parts)
	}
	for i := index; i < len(parts); i++ {
		part := parts[i]
		for j := 0; j < len(part); j++ {
			if start < len(abbr) && part[j] == abbr[start] {
				start += 1
				res := dfs(parts, i+1, abbr, start)
				if res {
					return true
				}
			} else {
				return false
			}
		}
	}
	return false
}

func getResult1(names []string, abbr string) {
	ans := []string{}
	for _, name := range names {
		parts := strings.Split(name, " ")
		if len(parts) > len(abbr) {
			continue
		}
		res := dfs1(parts, 0, abbr, 0)
		if res {
			ans = append(ans, name)
		}
	}
	fmt.Println(strings.Join(ans, ","))
}
func dfs1(parts []string, index int, abbr string, start int) bool {
	if start >= len(abbr) {
		return index >= len(parts)
	}
	for i := index; i < len(parts); i++ {
		part := parts[i]
		for j := 0; i < len(part); j++ {
			if start < len(abbr) && abbr[start] == part[j] {
				start += 1
				res := dfs(parts, index+1, abbr, start)
				if res {
					return true
				}
			} else {
				return false
			}
		}
	}
	return false
}
