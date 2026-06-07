package arrays

import (
	"fmt"
	"reflect"
)

func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	groups := make(map[[26]int][]string)

	for _, str := range strs {
		var freq [26]int

		for _, ch := range str {
			freq[ch-'a']++
		}

		groups[freq] = append(groups[freq], str)
	}

	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func RunTests_GroupAnagrams() {
	tests := []struct {
		name   string
		input  []string
		result [][]string
	}{
		{
			name:   "empty slice",
			input:  []string{},
			result: [][]string{},
		},
		{
			name:   "nil slice",
			input:  nil,
			result: [][]string{},
		},
		{
			name:   "single word",
			input:  []string{"abc"},
			result: [][]string{{"abc"}},
		},
		{
			name:  "multiple anagrams",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			result: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name:  "all anagrams",
			input: []string{"abc", "cab", "bca"},
			result: [][]string{
				{"abc", "cab", "bca"},
			},
		},
		{
			name:  "no anagrams",
			input: []string{"abc", "def", "ghi"},
			result: [][]string{
				{"abc"},
				{"def"},
				{"ghi"},
			},
		},
	}

	passed := 0
	failed := 0

	fmt.Println("=== Running GroupAnagrams Tests ===")

	for _, tc := range tests {
		got := GroupAnagrams(tc.input)
		expected := tc.result

		if reflect.DeepEqual(got, expected) {
			fmt.Printf("PASS: %s\n", tc.name)
			passed++
		} else {
			fmt.Printf(
				"FAIL: %s\nExpected: %v\nGot: %v\n",
				tc.name,
				expected,
				got,
			)
			failed++
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
