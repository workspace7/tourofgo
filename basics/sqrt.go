package basics //import "github.com/workspace7/tourofgo/basics"

//Sqrt - the square root of a float
func Sqrt(x float64) float64 {
	z := float64(1)

	for i := 0; i < 10; i++ {
		//z = z - (z*z-x)/2*z -- this does not work why ???
		z = z - (float64(z*z)-float64(x))/float64(2*z)
	}
	return float64(z)
}

// func main() {
// 	fmt.Printf("Ans: %0.2f \n", Sqrt(4))
// 	fmt.Printf("math.Sqrt: %0.2f", math.Sqrt(4))
// }
