package slice

// FilterBy returns a new slice containing only the items that satisfy the predicate.
func FilterBy[TItem any](items []TItem, predicate func(item TItem) bool) []TItem {
	var filteredItems []TItem
	for _, item := range items {
		if predicate(item) {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}
