package basics

import "fmt"

func LoopFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// func main() {
// 	loop_for()
// }
