package basics //import "github.com/workspace7/tourofgo/basics"

import "fmt"

//PointersCheck  see how pointers work
func PointersCheck() {
	i, j := 10.0, 20.0

	var p *float64 //  pointer of type int

	p = &i
	k := &j

	*k = *p + *k

	fmt.Printf("p %0.2f is of type %T", *p, p)
	fmt.Printf("k %0.2f is of type %T", *k, k)
}
