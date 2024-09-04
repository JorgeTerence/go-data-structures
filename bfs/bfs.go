package bfs

import (
	"colors"
	"fmt"
	g "geometry"
)

func IsValid(m [][]int, p g.Point, start int) bool {
	return p.X >= 0 && p.X < len(m[0]) && p.Y >= 0 && p.Y < len(m) && m[p.Y][p.X] == start
}

func Neighbours(m [][]int, v g.Point, start int) []g.Point {
	indexes := []g.Point{g.P(v.X, v.Y-1), g.P(v.X, v.Y+1), g.P(v.X-1, v.Y), g.P(v.X+1, v.Y)}
	res := make([]g.Point, 0)

	for _, coord := range indexes {
		if IsValid(m, coord, start) {
			res = append(res, coord)
		}
	}

	return res
}

func PrintMatrix(m [][]int, visited []g.Point) {
	for y, row := range m {
		for x, v := range row {
			if g.Contains(visited, g.P(x, y)) {
				fmt.Print(colors.Red, v, colors.Reset)
			} else {
				fmt.Print("", v)
			}
		}
		fmt.Print("\n")
	}
}
