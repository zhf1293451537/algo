package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getResult(s string) string {
	valid, _ := regexp.Match("[a-z0-9]+", []byte(s))
	if !valid {
		return "ERROR"
	}
	res := []byte{}
	stack := []byte{}

	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' && len(stack) > 0 {
			if len(res) > 0 && s[i] == res[len(res)-1] {
				return "ERROR"
			}
			temp, _ := strconv.Atoi(string(stack))
			res = append(res, []byte(strings.Repeat(string(s[i]), temp))...)
			stack = []byte{}
			continue
		}
		if s[i] >= '0' && s[i] <= '2' {
			return "ERROR"
		} else if s[i] > '2' && s[i] <= '9' {
			stack = append(stack, s[i])
			continue
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			res = append(res, s[i])
		} else {
			return "ERROR"
		}

	}
	return string(res)
}
func main() {

	fmt.Println(getResult("dd4c"))
}
