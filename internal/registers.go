package vm

import "fmt"

type Registers struct {
	A  int
	X  int
	L  int
	B  int
	S  int
	T  int
	F  float64
	PC int
	SW int
}

func NewRegisters() *Registers {
	return &Registers{}
}

func (registers *Registers) GetRegRef(idx int) *int {
	switch idx {
	case 0:
		return &registers.A
	case 1:
		return &registers.X
	case 2:
		return &registers.L
	case 3:
		return &registers.B
	case 4:
		return &registers.S
	case 5:
		return &registers.T
	case 6:
		panic("Float register not supported")
	case 8:
		return &registers.PC
	case 9:
		return &registers.SW
	default:
		panic(fmt.Sprintf("Invalid register index: %d", idx))
	}
}

func (registers *Registers) GetReg(idx int) int {
	regRef := registers.GetRegRef(idx)
	return *regRef
}

func (registers *Registers) SetReg(idx int, val int) {
	regRef := registers.GetRegRef(idx)
	*regRef = val
}
