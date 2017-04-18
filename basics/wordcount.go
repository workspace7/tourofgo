package basics

import "strings"

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
