package main

func main() {
	src := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	stack := [][]int{}
	for i := range src {
		for j := range src[i] {
			if src[i][j] == 1 {
				stack = append(stack, []int{i, j})
			}
		}
	}
	for {

	}
}
