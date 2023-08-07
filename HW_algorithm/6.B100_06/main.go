package main

import (
	"fmt"
	"strings"
)

func getResult(k int, s string) string {
	if !strings.Contains(s, "-") {
		return "ERROR"
	}
	res := []string{}
	for i := range s {
		if s[i] == '-' {
			res = append(res, s[0:i])
			temp := check(strings.ReplaceAll(s[i+1:], "-", ""), k)
			res = append(res, temp...)
			break
		}
	}
	return strings.Join(res, "-")
}
func check(s string, k int) []string {
	a, b, start := 0, 0, 0
	res := []string{}
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			a++
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			b++
		}
		if (i+1)%k == 0 || i == len(s)-1 {
			fmt.Println(a, b, start, i)
			if a > b {
				res = append(res, strings.ToLower(s[start:i+1]))
			} else if a < b {
				res = append(res, strings.ToUpper(s[start:i+1]))
			} else {
				res = append(res, s[start:i+1])
			}
			start = i + 1
			a, b = 0, 0
		}
	}
	return res
}
func main() {
	fmt.Println(getResult(3, "12abc-abCABc-4aB@"))
}
