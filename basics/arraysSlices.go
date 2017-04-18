package basics //import "github.com/workspace7/tourofgo/basics"

import (
	"fmt"
	"strings"
)

//ChantWithSlice is function to print chants from slice
func ChantWithSlice(chants []string) {
	fmt.Printf("Length: %d Capacity: %d  \n", len(chants), cap(chants))

	strChants := strings.Join(chants, "!")

	fmt.Printf("%v \n", strChants)

	chants = append(chants, "Hare Krishna", "Krishna Krishna Hare Hare", "Hare Rama", "Hare Rama", "Rama Rama Hare Hare!")

	//strChants = strings.Join(chants, "!")

	//fmt.Printf("After Append: %v \n", strChants)

	for i, s := range chants {
		fmt.Printf("chants[%d] = %s \t", i, s)
	}

	chants = chants[:0]
	fmt.Printf("\n Length: %d Capacity: %d  %v \n", len(chants), cap(chants), chants)

}

//ChantWithArray is a function to print the chants from passed array of strings
func ChantWithArray() []string {
	var chant [2]string

	chant[0] = "Hare"
	chant[1] = "Krishna"

	fmt.Printf("%s %s \n", chant[0], chant[1])

	return chant[0:1]
}

// func main() {
// 	chant_with_array()
// 	chants := make([]string, 2, 10)
// 	chants[0] = "Hare Krishna"
// 	chant_with_slice(chants)
// }
