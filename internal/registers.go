package vm

import "fmt"

// Condition codes
type ConditionCode int

const (
	CC_EQ        ConditionCode = 0b00
	CC_LT        ConditionCode = 0b01
	CC_GT        ConditionCode = 0b10
	CC_UNDEFINED ConditionCode = 0b11
)

type AddressingMode int

const (
	IMMEDIATE AddressingMode = 0b01
	DIRECT    AddressingMode = 0b10
	INDIRECT  AddressingMode = 0b11
	SIC       AddressingMode = 0b00
)

// https://sic-xe.github.io/chapters/sic.html
type Registers struct {
	A  int // accumulator
	X  int // index register
	L  int // linkage register (jumps)
	B  int // base register
	S  int
	T  int
	F  float64
	PC int
	SW int // status word
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
		// TODO: return error instead of panic, for better error handling in the visualization
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

func (registers *Registers) SetCC(cc ConditionCode) {
	registers.SW &= 0x00FFFF3F
	registers.SW |= int(cc) << 6
}

func (registers *Registers) GetCC() ConditionCode {
	return ConditionCode((registers.SW >> 6) & 0x03)
}

func getConditionCodes(r1Val, r2Val int) ConditionCode {
	if r1Val == r2Val {
		return CC_EQ
	} else if r1Val < r2Val {
		return CC_LT
	} else if r1Val > r2Val {
		return CC_GT
	}
	return CC_UNDEFINED
}
