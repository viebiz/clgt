package slice

func Map[TItem, TResult any](items []TItem, mapper func(item TItem) TResult) []TResult {
	var mappedItems []TResult
	for _, item := range items {
		mappedItems = append(mappedItems, mapper(item))
	}
	return mappedItems
}
