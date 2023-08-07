package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	s := "hAkDAjByBq"
	temp := []byte(s)
	n := 4
	sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })
	target := temp[n-1]
	res := bytes.Index([]byte(s), []byte{target})
	fmt.Println(res)
}
