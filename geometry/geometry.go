package geometry

type Point struct {
	X int
	Y int
}

func P(x, y int) Point {
    return Point{X: x, Y: y}
}

func Contains(arr []Point, v Point) bool {
	for _, a := range arr {
        if a == v {
            return true
        }
    }

    return false
}