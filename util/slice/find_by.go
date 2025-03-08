package slice

// FindBy returns the first item that satisfies the predicate.
func FindBy[TItem any](items []TItem, predicate func(item TItem) bool) (TItem, bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	var empty TItem
	return empty, false
}
