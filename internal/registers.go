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

// https://sic-xe.github.io/chapters/sic.html
type Registers struct {
	A  int32 // accumulator
	X  int32 // index register
	L  int32 // linkage register (jumps)
	B  int32 // base register
	S  int32
	T  int32
	F  float64
	PC int32
	SW int32 // status word
}

func NewRegisters() *Registers {
	return &Registers{}
}

func (registers *Registers) GetRegRef(idx int) *int32 {
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

func (registers *Registers) GetReg(idx int) int32 {
	regRef := registers.GetRegRef(idx)
	return *regRef
}

func (registers *Registers) SetReg(idx int, val int32) {
	regRef := registers.GetRegRef(idx)
	*regRef = val
}

func (registers *Registers) SetCC(cc ConditionCode) {
	registers.SW &= 0x00FFFF3F
	registers.SW |= int32(cc) << 6
}

func (registers *Registers) GetCC() ConditionCode {
	return ConditionCode((registers.SW >> 6) & 0x03)
}

func getConditionCodes(r1Val, r2Val int32) ConditionCode {
	if r1Val == r2Val {
		return CC_EQ
	} else if r1Val < r2Val {
		return CC_LT
	} else if r1Val > r2Val {
		return CC_GT
	}
	return CC_UNDEFINED
}
