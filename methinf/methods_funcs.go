package methinf

import "fmt"

//Dimension to hold length and breadth of a shape
type Dimension struct {
	l float64
	b float64
}

//AreaOfRect computes the area of rectangle
func (d Dimension) AreaOfRect() float64 {
	fmt.Printf("%v\n", d)
	d.l = d.l * d.l
	d.b = d.b * d.b
	return (d.l * d.b)
}

//ScaledArea computes the scaled area of rectangle
func (d *Dimension) ScaledArea(scale float64) {
	fmt.Printf("%v\n", d)
	d.l = d.l * scale
	d.b = d.b * scale
}

// func main() {
// 	d := Dimension{3, 4}
// 	fmt.Printf("Area :%0.2f", d.AreaOfRect())
// 	d.ScaledArea(10)
// 	fmt.Printf("Area :%0.2f", d.AreaOfRect())
// }
