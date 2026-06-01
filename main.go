package main

import (
	"fmt"

	duplicatenums "github.com/gayathrisuresh150501/neetcode150/1_duplicate_nums"
)

func main() {
	result := duplicatenums.HasDuplicateNums([]int{1, 2, 3, 4, 4})
	fmt.Print("Has Duplicates? ", result)
}
