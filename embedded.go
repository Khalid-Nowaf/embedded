package embedded

import "fmt"

// Base struct
type Base struct {
	A string
	B int
	C float32
}

// Xyz method
func (b Base) Xyz() {
	fmt.Println("xyz, a is:", b.A)
}

// Display method
func (b Base) Display() {
	fmt.Println("base, a is:", b.A)
}

// Derived struct
type Derived struct {
	Base         // embedding
	A    float64 //-SHADOWS
}

// Display method  -SHADOWED
func (d Derived) Display() {
	fmt.Println("derived a is:", d.A)
}
