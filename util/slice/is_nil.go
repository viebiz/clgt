package slice

func IsNil[TItem any](items []TItem) bool {
	return items == nil
}
