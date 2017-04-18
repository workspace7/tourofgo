package basics //import "github.com/workspace7/tourofgo/basics"
import "fmt"

//LoopFor - just a simple method to dmeo for loop
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
