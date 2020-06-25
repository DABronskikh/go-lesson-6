package transactions

import (
	"sort"
)

func Sort(transitions []int64) []int64 {
	sort.SliceStable(transitions, func(i, j int) bool {
		return transitions[i] > transitions[j]
	})
	return transitions
}
