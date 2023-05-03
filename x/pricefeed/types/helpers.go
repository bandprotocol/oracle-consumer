package types

import "sort"

func CalculateGas(base, each, n uint64) uint64 {
	return base + each*n
}

func SortKeysUint64StringMap(uint64StringMap map[uint64][]string) []uint64 {
	keys := make([]uint64, 0)
	for k := range uint64StringMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}
