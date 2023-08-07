package main

import "fmt"

func main() {
	var sum, n, left, right int
	res := []int{}
	fmt.Scan(&sum, &n)
	temp := sum / n
	fmt.Println(temp)
	halflen := 0
	if n%2 == 0 {
		halflen = n/2 - 1
		left = temp - halflen
		right = temp + halflen + 1
	} else {
		halflen = n / 2
		left = temp - halflen
		right = temp + halflen
	}
	if left < 0 {
		fmt.Println(-1)
		return
	}
	for i := left; i <= right; i++ {
		res = append(res, i)
	}
	fmt.Println(res)
}
