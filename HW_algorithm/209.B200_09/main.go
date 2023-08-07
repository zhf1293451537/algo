package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getResult() {

}
func main() {
	s := "3[k]2[mn]"
	stack := []byte{}
	idxs := []int{}
	nums := []int{}
	tmpRepeatCount := []byte{}

	for i := range s {
		if s[i] == '[' {
			repeatCount, _ := strconv.Atoi(string(tmpRepeatCount))
			nums = append(nums, repeatCount)
			tmpRepeatCount = []byte{}
			idxs = append(idxs, len(stack))
		} else if s[i] == ']' {
			start := idxs[len(idxs)-1]
			idxs = idxs[:len(idxs)-1]
			repeatCount := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			repeatStr := string(stack[start:])
			stack = stack[:start]
			final := strings.Repeat(repeatStr, repeatCount)
			stack = append(stack, final...)
		} else if s[i] >= '0' && s[i] <= '9' {
			tmpRepeatCount = append(tmpRepeatCount, s[i])
		} else {
			stack = append(stack, s[i])
		}
	}
	fmt.Println(string(stack))
}
