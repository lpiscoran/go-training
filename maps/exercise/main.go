package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	words := strings.Fields(s)
	// fmt.Println(words)
	// fmt.Println(len(words))

	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	// fmt.Println(counts)
	return counts
	//

}

func main() {
	wc.Test(WordCount)
}
