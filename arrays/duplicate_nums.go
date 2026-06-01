package arrays

import "fmt"

func HasDuplicateNums[T comparable](nums []T) bool {
	hashSet := make(map[T]struct{}, len(nums))

	for _, v := range nums {
		if _, ok := hashSet[v]; ok {
			return true
		}

		hashSet[v] = struct{}{}
	}

	return false
}

func RunTests_HasDuplicateNums() {
	tests := []struct {
		name   string
		input  []int
		result bool
	}{
		{
			"empty slice",
			[]int{},
			false,
		},
		{
			"nil slice",
			nil,
			false,
		},
		{
			"unique values",
			[]int{1, 2, 3, 4},
			false,
		},
		{
			"duplicate values",
			[]int{1, 2, 3, 3, 4},
			true,
		},
	}

	passed := 0
	failed := 0

	fmt.Println("=== Running HasDuplicateNums Tests ===")

	for _, tc := range tests {
		got := HasDuplicateNums(tc.input)

		if got != tc.result {
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
