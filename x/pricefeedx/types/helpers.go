package types

func CalculateGas(base, each, n uint64) uint64 {
	return base + each*n
}
