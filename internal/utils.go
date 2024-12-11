package vm

func extendSign(value int32, bits int) int32 {
	mask := int32(1 << (bits - 1))
	if mask&value != 0 {
		return value | (^(mask - 1))
	}
	return value
}
