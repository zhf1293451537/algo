package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temp := scanner.Text()
	operations := strings.Split(temp, " ")
	snake := [][]int{}
	var n, m int
	fmt.Scan(&n, &m)
	src := make([][]string, n)
	for i := range src {
		src[i] = make([]string, m)
		for j := range src[i] {
			fmt.Scan(&src[i][j])
			if src[i][j] == "H" {
				snake = append(snake, []int{i, j})
			}
		}
	}
	fmt.Println(operations, snake, n, m)
	i, j := 0, 0
	for _, v := range operations {
		switch v {
		case "U":
			i = -1
			j = 0
		case "D":
			i = 1
			j = 0
		case "L":
			i = 0
			j = -1
		case "R":
			i = 0
			j = 1
		case "G":
			newi, newj := snake[0][0]+i, snake[0][1]+j
			fmt.Println(newi, newj)
			if newi < 0 || newi >= n || newj < 0 || newj >= m {
				fmt.Println(len(snake))
				return
			} else if src[newi][newj] == "F" {
				snake = append([][]int{{newi, newj}}, snake...)
				src[newi][newj] = "H"
			} else if src[newi][newj] == "E" {
				src[newi][newj] = "H"
				snake = append([][]int{{newi, newj}}, snake...)
				src[snake[len(snake)-1][0]][snake[len(snake)-1][1]] = "E"
				snake = snake[:len(snake)-1]
			} else {
				fmt.Println(len(snake))
				return
			}
		}

	}
	fmt.Println(len(snake))

}
