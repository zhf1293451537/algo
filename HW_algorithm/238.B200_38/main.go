package main

import "fmt"

func main() {
	src := [][]int{
		{2, 1},
		{3, 2},
		{4, 3},
		{4, 5},
	}
	mapV := map[int][]int{}
	end := 0
	for i := range src {
		if src[i][1] > end {
			end = src[i][1]
		}
		if src[i][0] > end {
			end = src[i][0]
		}
		if src[i][1] < src[i][0] {
			mapV[src[i][0]] = append(mapV[src[i][0]], src[i][1])
		}
	}
	ans := []int{}
	for i := 1; i <= end; i++ {
		if _, ok := mapV[i]; !ok {
			ans = append(ans, 0)
		} else {
			res := 0
			temp := []int{i}
		l1:
			for len(temp) > 0 {
				newTemp := []int{}
				for i := range temp {
					if _, ok := mapV[temp[i]]; !ok {
						break l1
					}
					res += len(mapV[temp[i]])
					for _, v := range mapV[temp[i]] {
						newTemp = append(newTemp, v)
					}
				}
				temp = newTemp
			}
			ans = append(ans, res)
		}
	}
	fmt.Println(ans)
}
