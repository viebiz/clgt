package slice

func IsEmpty[TItem any](items []TItem) bool {
	return len(items) == 0
}
