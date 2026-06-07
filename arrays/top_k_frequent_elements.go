package arrays

import (
	"container/heap"
	"fmt"
	"reflect"
	"slices"
)

type Item struct {
	Num  int
	Freq int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Freq < h[j].Freq }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func TopKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}

	// O(N)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// O(M log K), where M = number of unique elements
	h := &MinHeap{}
	heap.Init(h)

	for num, count := range freq {
		heap.Push(h, Item{
			Num:  num,
			Freq: count,
		})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Extract result
	result := make([]int, h.Len())

	for i := len(result) - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(Item).Num
	}

	return result
}

func RunTests_TopKFrequent() {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "empty input",
			nums:     []int{},
			k:        0,
			expected: []int{},
		},
		{
			name:     "single element",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "leetcode example",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "all unique",
			nums:     []int{1, 2, 3, 4},
			k:        2,
			expected: []int{2, 4}, // any 2 would technically be valid
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -1, -1, 2, 2, 3},
			k:        2,
			expected: []int{-1, 2},
		},
		{
			name:     "k equals unique count",
			nums:     []int{1, 1, 2, 2, 3},
			k:        3,
			expected: []int{1, 2, 3},
		},
	}

	passed := 0
	failed := 0

	fmt.Println("=== Running TopKFrequent Tests ===")

	for _, tc := range tests {
		got := TopKFrequent(tc.nums, tc.k)

		slices.Sort(got)
		slices.Sort(tc.expected)

		if reflect.DeepEqual(got, tc.expected) {
			fmt.Printf("PASS: %s\n", tc.name)
			passed++
		} else {
			fmt.Printf(
				"FAIL: %s\nExpected: %v\nGot: %v\n",
				tc.name,
				tc.expected,
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
