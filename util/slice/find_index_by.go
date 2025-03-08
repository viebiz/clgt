package slice

// FindIndexBy returns the index of the first item that satisfies the predicate.
func FindIndexBy[TItem any](items []TItem, predicate func(item TItem) bool) int {
	for i, item := range items {
		if predicate(item) {
			return i
		}
	}
	return -1
}
