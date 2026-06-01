package duplicatenums

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
