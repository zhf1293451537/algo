package main

import (
	"fmt"
	"strings"
)

func getResult(s string, n int) string {
	stack := []byte{}
	result := []string{}
	for i, _ := range s {
		if s[i] == '_' && (len(stack) == 0 || stack[0] != '"') {
			result = append(result, string(stack))
			stack = []byte{}
		} else if s[i] == '"' && len(stack) != 0 {
			stack = append(stack, s[i])
			result = append(result, string(stack))
			stack = []byte{}
		} else {
			stack = append(stack, s[i])
		}
	}
	if len(stack) > 0 {
		result = append(result, string(stack))
	}
	ans := []string{}
	for _, v := range result {
		if v != "" {
			ans = append(ans, v)
		}
	}
	if n > len(ans) {
		return "ERROR"
	}
	ans[n] = "******"
	return strings.Join(ans, "_")
}
func main() {
	s := "\"password_\"_a123456789_timeout__100_\"\"_"
	fmt.Println(getResult(s, 1))
}
