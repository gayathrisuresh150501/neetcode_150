package arrays

import (
	"fmt"
	"reflect"
)

func TwoSum[T comparable](nums []T) []int {
	hashSet := make(map[T]int, len(nums))

	for idx, num := range nums {
		if i, ok := hashSet[num]; ok {
			return []int{idx, i}
		}

		hashSet[num] = idx
	}

	return nil
}

func RunTests_TwoSum() {
	tests := []struct {
		name   string
		input  []int
		target int
		result []int
	}{
		{
			"empty slice",
			[]int{},
			3,
			[]int{},
		},
		{
			"nil slice",
			nil,
			3,
			[]int{},
		},
		{
			"with sum values",
			[]int{1, 2, 3, 4},
			6,
			[]int{1, 3},
		},
		{
			"no sum values",
			[]int{1, 2, 3, 3, 4},
			11,
			nil,
		},
		{
			"input with negative values",
			[]int{-1, 2, -2, 4},
			-3,
			[]int{0, 2},
		},
	}

	passed := 0
	failed := 0

	fmt.Println("=== Running HasDuplicateNums Tests ===")

	for _, tc := range tests {
		got := TwoSum(tc.input)

		if reflect.DeepEqual(got, tc.result) {
			fmt.Printf("%s: expected %v, got %v\n",
				tc.name, tc.result, got)
			failed += 1
		} else {
			fmt.Printf("PASS: %s\n", tc.name)
			passed += 1
		}
	}

	fmt.Println("\n=== Test Summary ===")
	fmt.Printf("Total : %d\n", len(tests))
	fmt.Printf("Passed: %d\n", passed)
	fmt.Printf("Failed: %d\n", failed)

	if failed == 0 {
		fmt.Println("All tests passed!")
	}
}
