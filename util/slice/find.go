package slice

import "slices"

// Find returns the first item that satisfies the predicate.
func Find[TItem any](items []TItem, predicate func(item TItem) bool) (TItem, bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	slices.Concat()
	var empty TItem
	return empty, false
}
