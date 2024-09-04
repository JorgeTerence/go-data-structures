package main

import (
	"bfs"
	"fmt"
	g "geometry"
	"os"
)

func main() {
	m := [][]int{
		{1, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{0, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 1, 1},
	}

	bfs.PrintMatrix(m, []g.Point{})
	x, y := -1, -1

	fmt.Println("Digite as coordenadas. Ex: 0 2")
	n, err := fmt.Scanf("%d %d", &x, &y)

	if n < 2 || err != nil {
		fmt.Println(x)
		fmt.Println(y)
		os.Exit(1)
	}

	visited := make([]g.Point, 0)
	queue := []g.Point{g.P(x, y)}

	bfs.PrintMatrix(m, queue)

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if !g.Contains(visited, v) {
			visited = append(visited, v)
			for _, w := range bfs.Neighbours(m, v, m[y][x]) {
				if !g.Contains(visited, w) {
					queue = append(queue, w)
				}
			}
		}

	}

	fmt.Println("---------")
	bfs.PrintMatrix(m, visited)
}
