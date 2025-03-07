package slice

// Concat returns a new slice containing all the items from the input slices.
func Concat[TItem any](slices ...[]TItem) []TItem {
	var totalLen int
	for _, slice := range slices {
		totalLen += len(slice)
	}
	concatenated := make([]TItem, 0, totalLen)
	for _, slice := range slices {
		concatenated = append(concatenated, slice...)
	}
	return concatenated
}
