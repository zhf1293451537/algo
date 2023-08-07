package main

import (
	"fmt"
	"strings"
)

func getResult(s string, start, stop int) string {
	src := strings.Split(s, " ")
	if start < 0 {
		start = 0
	}
	if stop > len(src) {
		stop = len(src) - 1
	}
	if stop < start || stop < 0 {
		return s
	}
	stack := []string{}
	for i := len(src) - 1; i >= 0; i-- {
		if i >= start && i <= stop {
			stack = append(stack, src[i])
		}
	}
	res := append(append(src[:start], stack...), src[stop+1:]...)
	return strings.Join(res, " ")
}
func main() {
	fmt.Println(getResult("I am a develop", -1, -2))
}
