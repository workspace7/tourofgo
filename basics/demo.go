package basics //import "github.com/workspace7/tourofgo/basics"

import (
	"bytes"
	"fmt"
)

func IntsToStrings(values []int) string {

	var buf bytes.Buffer

	buf.WriteByte('[')

	for i, v := range values {

		if i > 0 {
			buf.WriteString(", ")
		}

		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}

// func main() {
// 	fmt.Println(intsToStrings([]int{1, 2, 3, 4, 5}))
// }
