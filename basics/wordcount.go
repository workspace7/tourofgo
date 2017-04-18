package basics //import "github.com/workspace7/tourofgo/basics"

import "strings"

//WordCount - counts the words in a string
func WordCount(s string) map[string]int {
	strs := strings.Fields(s)

	wcMap := make(map[string]int)

	for _, v := range strs {
		wcMap[v]++
	}

	return wcMap
}

// func main() {
// 	wc.Test(WordCount)
// }
