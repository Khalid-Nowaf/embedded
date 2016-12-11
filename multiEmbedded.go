package embedded

// NamedObj struct
type NamedObj struct {
	Name string
}

// Shape struct
type Shape struct {
	NamedObj  //inheritance
	Color     int32
	IsRegular bool
}

// Point struct
type Point struct {
	X, Y float64
}

// Rectangle struct
type Rectangle struct {
	NamedObj            //multiple inheritance
	Shape               //^^
	Center        Point //standard composition
	Width, Height float64
}
