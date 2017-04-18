package methinf // import "github.com/workspace7/tourofgo/methinf"

import (
	"fmt"
	"math"
)

//ErrNegativeSqrt is used ot hold error object
type ErrNegativeSqrt float64

//Error defines how the err object is built
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Can't sqrt %0.2f as its negative", float64(e))
}

//Sqrt - will do a mathematical square root
func Sqrt(x float64) (float64, error) {
	var sqrt float64
	var e error

	if x < 0 {
		e = ErrNegativeSqrt(x)
	} else {
		sqrt = math.Sqrt(x)
	}

	return sqrt, e
}

// func main() {
// 	result, err := Sqrt(2)

// 	if err == nil {
// 		fmt.Printf("\nSquare Root of 2 is %0.2f", result)
// 	} else {
// 		fmt.Printf("\nError %v", err)
// 	}

// 	result, err = Sqrt(-2)

// 	if err == nil {
// 		fmt.Printf("\nSquare Root of 2 is %0.2f", result)
// 	} else {
// 		fmt.Printf("\nError %v", err)
// 	}
// }
