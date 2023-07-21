package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getResult(s string) string {
	res := strings.Split(s, "#")
	if len(res) != 4 {
		return "invalid IP"
	}
	ans := ""
	for i, v := range res {
		temp, err := strconv.Atoi(v)
		if err != nil {
			return "invalid IP"
		}
		if (i == 0 && (temp > 128 || temp < 1)) || len(v) == 0 {
			return "invalid IP"
		}
		if (i != 0 && (temp > 255 || temp < 0)) || len(v) == 0 {
			return "invalid IP"
		}
		ss := strconv.FormatInt(int64(temp), 16)
		if len(ss) == 1 {
			ans += "0" + ss
		} else {
			ans += ss
		}
	}
	final, _ := strconv.ParseInt(ans, 16, 64)
	return fmt.Sprintf("%d", final)
}
func main() {
	fmt.Println(getResult("255#255##255"))
}
