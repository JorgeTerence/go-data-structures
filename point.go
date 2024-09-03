package main

type Point struct {
	x int
	y int
}

func Contains(arr []Point, v Point) bool {
	for _, a := range arr {
        if a == v {
            return true
        }
    }

    return false
}