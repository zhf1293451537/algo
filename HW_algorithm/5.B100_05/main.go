package main

import "fmt"

func getResult(s string) int {
	maxlen := -1
	hasletter := false
	l, r := 0, 0
	letteridx := []int{}
	for r < len(s) {
		if isLetter(s[r]) {
			hasletter = true
			letteridx = append(letteridx, r)
		}
		if len(letteridx) > 1 {
			l = letteridx[0] + 1
			letteridx = letteridx[1:]
		}
		if r == l {
			r++
			continue
		}
		maxlen = max(maxlen, r-l+1)
		r++
	}
	if !hasletter {
		return -1
	}
	return maxlen
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func isLetter(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}
func main() {
	fmt.Println(getResult("abcdef"))
}
