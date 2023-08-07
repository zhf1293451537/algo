package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 1000
	binary := strconv.FormatInt(int64(n), 2)
	src := []string{}
	for len(binary) > 7 {
		src = append(src, binary[len(binary)-7:])
		binary = binary[:len(binary)-7]
		if len(binary) > 0 {
			src[len(src)-1] = "1" + src[len(src)-1]
		}
	}
	res := ""
	src = append(src, binary)
	for i := range src {
		temp, _ := strconv.ParseInt(src[i], 2, 64)
		final := strconv.FormatInt(temp, 16)
		if len(final) == 1 {
			final = "0" + final
		}
		res += strings.ToUpper(final)
	}
	fmt.Println(res)
}
