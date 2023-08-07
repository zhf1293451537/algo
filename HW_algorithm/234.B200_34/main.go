package main

import (
	"fmt"
	"sort"
)

func main() {
	books := [][]int{
		{20, 17},
		{15, 16},
		{15, 17},
		{10, 10},
		{10, 9},
	}
	sort.Slice(books, func(i, j int) bool {
		if books[i][0] == books[j][0] {
			return books[i][1] > books[j][1]
		}
		return books[i][0] < books[j][0]
	})
	widths := []int{}
	for i := range books {
		widths = append(widths, books[i][1])
	}
	dp := []int{widths[0]}
	for i := 1; i < len(widths); i++ {
		if widths[i] > dp[len(dp)-1] {
			dp = append(dp, widths[i])
			continue
		}
		if widths[i] < dp[0] {
			dp[0] = widths[i]
			continue
		}
		idx := binarySearch(dp, widths[i])
		if idx < 0 {
			dp[-idx-1] = widths[i]
		}
	}
	fmt.Println(len(dp), dp)
}
func binarySearch(arr []int, key int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		midV := arr[mid]
		if key > midV {
			low = mid + 1
		} else if key < midV {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -(low + 1)
}
