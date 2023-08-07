package main

import (
	"fmt"
	"sort"
)

func getResult(start int, n int, src []string) string {
	res := src[start]
	lmap := map[byte][]string{}
	for i := range src {
		if i == start {
			continue
		}
		if _, ok := lmap[src[i][0]]; !ok {
			lmap[src[i][0]] = []string{}
		}
		lmap[src[i][0]] = append(lmap[src[i][0]], src[i])

	}
	for _, v := range lmap {
		sort.Strings(v)

	}
	for {
		if _, ok := lmap[res[len(res)-1]]; !ok {
			break
		}
		temp := lmap[res[len(res)-1]]
		maxlen := 0
		stop := -1
		for i, v := range temp {
			if len(v) > maxlen {
				maxlen = len(v)
				stop = i
			}
		}
		maxstr := temp[stop]
		lmap[res[len(res)-1]] = append(temp[:stop], temp[stop+1:]...)
		res += maxstr
	}
	return res
}
func main() {
	fmt.Println(getResult(3, 6, []string{"word", "dd", "da", "dc", "dword", "d"}))
}
