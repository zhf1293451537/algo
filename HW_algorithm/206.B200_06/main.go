package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getResult(src []int) int {
	sort.Slice(src, func(i, j int) bool { return src[i] > src[j] })
	length := len(src)
	sum := 0
	for i := range src {
		sum += src[i]
	}
	for length > 1 {
		if canPartition(src, sum, length) {
			return length
		}
		length -= 1
	}
	return -1
}
func canPartition(src []int, sum, length int) bool {
	if sum%length != 0 {
		return false
	}
	subSum := sum / length
	if subSum < src[0] {
		return false
	}
	for len(src) > 0 && src[0] == subSum {
		src = src[1:]
		length -= 1
	}
	buckets := make([]int, length)
	return partition(src, 0, buckets, subSum)
}
func partition(src []int, index int, buckets []int, subSum int) bool {
	fmt.Println(src, buckets)
	if len(src) == index {
		return true
	}
	sel := src[index]
	for i := range buckets {
		if i > 0 && buckets[i] == buckets[i-1] {
			continue
		}
		fmt.Println(sel, buckets)
		if sel+buckets[i] <= subSum {
			buckets[i] += sel
			if partition(src, index+1, buckets, subSum) {
				return true
			}
			buckets[i] -= sel
		}
	}
	return false
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	temp := strings.Split(str, " ")
	src := make([]int, len(temp))
	for i := range temp {
		dig, _ := strconv.Atoi(temp[i])
		src[i] = dig
	}
	fmt.Println(getResult(src))
}
