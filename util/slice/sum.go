package slice

import "github.com/viebiz/clgt/types"

func Sum[TNumber types.Number](items []TNumber) TNumber {
	var total TNumber
	for _, item := range items {
		total += item
	}
	return total
}
