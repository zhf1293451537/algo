package main

import (
	"fmt"
	"sort"
	"strings"
)

func getResult(s, pre string) string {
	str := strings.ReplaceAll(s, ",", " ")
	str = strings.ReplaceAll(str, "'", " ")
	str = strings.ReplaceAll(str, ".", " ")
	src := strings.Split(strings.TrimSpace(str), " ")
	res := []string{}
	for _, v := range src {
		if strings.HasPrefix(v, pre) {
			res = append(res, v)
		}
	}
	if len(res) > 0 {
		sort.Strings(res)
		return strings.Join(res, " ")
	}
	return pre

}
func main() {
	fmt.Println(getResult("The furehest distance in the world,don't knoe front ", "f"))
}
