package basics //import "github.com/workspace7/tourofgo/basics"

//Pic is used to build  the image pixels
func Pic(dx, dy int) [][]uint8 {

	arr_dy := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		arr_dx := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			arr_dx[y] = uint8(x ^ y)
		}
		arr_dy[y] = arr_dx
	}

	return arr_dy
}

// func main() {
// 	pic.Show(Pic)
// }
