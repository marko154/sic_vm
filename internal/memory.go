package vm

import "fmt"

type Memory [MAX_ADDRESS]byte

const MAX_ADDRESS = 1 << 20

func (mem *Memory) GetByte(address int) (byte, error) {
	if address < 0 || address >= MAX_ADDRESS {
		return 0, fmt.Errorf("invalid address: %d", address)
	}
	return mem[address], nil
}

func (mem *Memory) SetByte(address int, value byte) error {
	return nil
}
