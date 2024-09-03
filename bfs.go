package main

import (
	"fmt"
	"os"
)

var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

func isValid(m [][]int, p Point, start int) bool {
	return p.x >= 0 && p.x < len(m[0]) && p.y >= 0 && p.y < len(m) && m[p.y][p.x] == start
}

func neighbours(m [][]int, v Point, start int) []Point {
	indexes := []Point{{v.x, v.y - 1}, {v.x, v.y + 1}, {v.x - 1, v.y}, {v.x + 1, v.y}}
	res := make([]Point, 0)

	for _, coord := range indexes {
		if isValid(m, coord, start) {	
			res = append(res, coord)
		}
	}

	return res
}

func printMatrix(m [][]int, visited []Point) {
	for y, row := range m {
		for x, v := range row {
			if Contains(visited, Point{x, y}) {
				fmt.Print(Red, v, Reset)
			} else {
				fmt.Print("", v)
			}
		}
		fmt.Println("")
	}
}

func main() {
	m := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 1, 1},
		{0, 1, 1, 1, 1},
		{1, 1, 0, 0, 1},
		{1, 0, 1, 0, 0},
	}

	printMatrix(m, []Point{})
	x, y := -1, -1

	fmt.Print("Coords: ")
	n, err := fmt.Scanf("%d %d", &x, &y)

	if n < 2 || err != nil {
		fmt.Println(x)
		fmt.Println(y)
		os.Exit(1)
	}

	fmt.Println(m[y][x])

	visited := make([]Point, 0)
	queue := []Point{{x, y}}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if !Contains(visited, v) {
			visited = append(visited, v)
			for _, w := range neighbours(m, v, m[y][x]) {
				if !Contains(visited, w) {
					queue = append(queue, w)
				}
			}
		}

		printMatrix(m, visited)
		fmt.Println("---------")
	}

	printMatrix(m, visited)
}
