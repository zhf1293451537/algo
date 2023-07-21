package main

import "fmt"

func getResult(s string) int {
	falg := true
	res := 0
	temp := 0
	for i, _ := range s {
		if s[i] == '-' {
			falg = false
		}
		if s[i] >= '0' && s[i] <= '9' {
			if falg {
				res += int(s[i] - '0')
			} else {
				temp = temp*10 + int(s[i]-'0')
			}
		} else {
			res += temp * -1
			temp = 0
		}
	}
	return res
}
func main() {
	fmt.Println(getResult("bb12-34aa"))
}
