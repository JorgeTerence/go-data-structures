package main

import (
	"fmt"
	"os"
	"os/exec"

	"bfs"
	g "geometry"
)

const (
	upArrow    = 'w'
	downArrow  = 's'
	leftArrow  = 'a'
	rightArrow = 'd'
)

func main() {
	m := [][]int{
		{1, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{0, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 1, 1},
	}

	start := g.P(0, 0)

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for i := 1; ; i++ {
		bfs.PrintMatrix(m, []g.Point{start})
		dir := ""
		fmt.Scanf("%s", &dir)

		if len(dir) == 0 {
			break
		}

		if dir[0] == upArrow {
			start.Y -= 1
		} else if dir[0] == downArrow {
			start.Y += 1
		} else if dir[0] == leftArrow {
			start.X -= 1
		} else if dir[0] == rightArrow {
			start.X += 1
		} else {
			break
		}

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	visited := make([]g.Point, 0)
	queue := []g.Point{start}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if !g.Contains(visited, v) {
			visited = append(visited, v)
			for _, w := range bfs.Neighbours(m, v, m[start.Y][start.X]) {
				if !g.Contains(visited, w) {
					queue = append(queue, w)
				}
			}
		}

	}

	bfs.PrintMatrix(m, visited)
}
