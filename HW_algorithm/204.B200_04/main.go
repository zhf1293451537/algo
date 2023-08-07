package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func getResult(n int, seatOrLeave []int) int {
	seat := []int{}
	ans := -1
	for _, v := range seatOrLeave {
		onSeat := len(seat)
		if v == 1 {
			if onSeat == 0 {
				ans = 0
				seat = append(seat, ans)
			} else if onSeat == 1 {
				ans = n - 1
				seat = append(seat, ans)
			} else {
				maxDis := 0
				start := -1
				for i := 1; i < len(seat); i++ {
					p := seat[i-1]
					c := seat[i]
					if c == p+1 {
						continue
					}
					dis := math.Ceil(float64(c-p-1) / 2.0)
					fmt.Println(dis)
					if int(dis) > maxDis {
						maxDis = int(dis)
						start = p
					}
				}
				brand_dis := n - 1 - seat[len(seat)-1]
				if brand_dis > maxDis {
					maxDis = brand_dis
					start = seat[len(seat)-1]
				}
				ans = start + maxDis
				if ans != -1 {
					seat = append(seat, ans)
					sort.Slice(seat, func(i, j int) bool { return seat[i] < seat[j] })
				}
			}
		} else {
			for m, n := range seat {
				if -v == n {
					seat = append(seat[:m], seat[m+1:]...)
				}
			}
		}
		fmt.Println(seat)
	}
	return ans
}
func main() {
	var n int
	fmt.Scan(&n)
	var str string
	fmt.Scan(&str)
	temp := strings.Split(str[1:len(str)-1], ",")
	seatOrLeave := []int{}
	for _, v := range temp {
		i, _ := strconv.Atoi(v)
		seatOrLeave = append(seatOrLeave, i)
	}
	fmt.Println(n, seatOrLeave)
	fmt.Println(getResult(n, seatOrLeave))
}
