package methinf

import "fmt"

// Calculator type defines the method of the caculator
type Calculator interface {
	Add()
	Sub()
}

//IntNumbers a simple type to hold the slice of int number
type IntNumbers struct {
	numbers []int
}

//FloatNumbers a simple type to hold the slice of int number
type FloatNumbers struct {
	numbers []float64
}

//Add will add the integer numbers
func (ints IntNumbers) Add() {
	sum := 0
	for _, v := range ints.numbers {
		sum += v
	}
	fmt.Printf("\nSum of integers %v = %d", ints, sum)
}

//Sub will subtract the integer numbers
func (ints IntNumbers) Sub() {
	sum := 0
	for _, v := range ints.numbers {
		sum -= v
	}
	fmt.Printf("\nDiff of integers %v = %d", ints, sum)
}

//Add will add the float numbers
func (floats FloatNumbers) Add() {
	sum := 0.0
	for _, v := range floats.numbers {
		sum += v
	}
	fmt.Printf("\nSum of float numbers %v = %0.2f", floats, sum)
}

//Sub will add the float numbers
func (floats FloatNumbers) Sub() {
	sum := 0.0
	for _, v := range floats.numbers {
		sum -= v
	}
	fmt.Printf("\nDiff of float numbers %v = %0.2f", floats, sum)
}

func describe(calc Calculator) {
	fmt.Printf("(%v,%T)", calc, calc)
}

// func main() {

// 	calc := &IntNumbers{[]int{1, 2, 3, 4, 5}}
// 	calc.Add()
// 	calc.Sub()
// 	describe(calc)

// 	fcalc := &FloatNumbers{[]float64{1.2, 2.7, 3.3, 4.8, 5.0}}
// 	fcalc.Add()
// 	fcalc.Sub()
// 	describe(fcalc)

// }
