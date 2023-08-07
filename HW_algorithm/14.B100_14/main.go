package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getResult(s1, s2 string) string {
	src := strings.Split(s2, " ")
	res := []string{}
	for i := 0; i < len(src); {
		start, stop := i+1, i+2
		temp := src[stop] + src[start]
		length, _ := strconv.ParseInt(temp, 16, 10)
		if src[i] == s1 {
			res = append(res, src[i+3:i+3+int(length)]...)
			break
		}
		i = i + 3 + int(length)
	}
	return strings.Join(res, " ")
}
func main() {
	fmt.Println(getResult("90", "32 01 00 AE 90 02 00 01 02 30 03 00 AB 32 31 31 02 00 32 33 33 01 00 CC"))
}
