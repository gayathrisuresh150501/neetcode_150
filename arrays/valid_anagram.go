package arrays

import (
	"fmt"
	"strings"
)

func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	freq := make([]int, 26)

	for idx := range str1 {
		freq[str1[idx]-'a']++
		freq[str2[idx]-'a']--
	}

	for _, count := range freq {
		if count != 0 {
			return false
		}
	}

	return true
}

func RunTests_IsAnagram() {
	tests := []struct {
		name   string
		str1   string
		str2   string
		result bool
	}{
		{
			name:   "empty strings",
			str1:   "",
			str2:   "",
			result: true,
		},
		{
			name:   "valid anagram",
			str1:   "abc",
			str2:   "bca",
			result: true,
		},
		{
			name:   "invalid anagram",
			str1:   "abc",
			str2:   "def",
			result: false,
		},
		{
			name:   "mixed case",
			str1:   "Hello",
			str2:   "EollH",
			result: true,
		},
	}

	passed := 0
	failed := 0

	for _, tc := range tests {
		got := IsAnagram(tc.str1, tc.str2)

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
