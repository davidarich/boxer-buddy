package domain

type Point struct {
	X int
	Y int
}

type Rectangle struct {
	Length int
	Width int
}

// Location and size of a Window
type Window struct {
	Point Point
	Size Rectangle
}
