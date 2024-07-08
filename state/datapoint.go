package state

type Series struct {
	Name       string
	DataPoints []Point
}

type Point struct {
	TS    int64
	Value float64
}
