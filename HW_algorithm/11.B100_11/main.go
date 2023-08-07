package main

import (
	"fmt"
)

func getResult(s string) string {
	res := []byte{}
	src := []int64{1, 2, 4}
	for i := range s {
		if i >= 3 {
			temp := src[i-1] + src[i-2] + src[i-3]
			src = append(src, temp)
		}
		offset := src[i]
		res = append(res, ((s[i]+byte(offset))-'a')%26+'a')
	}
	fmt.Println(src)
	return string(res)
}
func main() {
	fmt.Println(getResult("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"))
}
