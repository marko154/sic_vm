package vm

func extendSign(value int, bits int) int {
	mask := 1 << (bits - 1)
	if mask&value != 0 {
		return value | (^(mask - 1))
	}
	return value
}
