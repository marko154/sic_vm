package vm

// Opcode defines the type for opcodes
type Opcode byte

// SIC/XE Opcodes
const (
	// A ← (A) + (m..m+2)
	ADD Opcode = 0x18
	// F ← (F) + (m..m+5)
	ADDF Opcode = 0x58
	// r2 ← (r2) + (r1)
	ADDR Opcode = 0x90
	// A ← (A) & (m..m+2)
	AND Opcode = 0x40
	// r ← 0
	CLEAR Opcode = 0xB4
	// (A) : (m..m+2)
	COMP Opcode = 0x28
	// (F) : (m..m+5)
	COMPF Opcode = 0x88
	// (r1) : (r2)
	COMPR Opcode = 0xA0
	// A ← (A) / (m..m+2)
	DIV Opcode = 0x24
	// F ← (F) / (m..m+5)
	DIVF Opcode = 0x64
	// r2 ← (r2) / (r1)
	DIVR Opcode = 0x9C
	// A ← int(F)
	FIX Opcode = 0xC4
	// F ← float(A)
	FLOAT Opcode = 0xC0
	// haltio(A)
	HIO Opcode = 0xF4
	// PC ← m
	J Opcode = 0x3C
	// PC ← m if CC is =
	JEQ Opcode = 0x30
	// PC ← m if CC is >
	JGT Opcode = 0x34
	// PC ← m if CC is <
	JLT Opcode = 0x38
	// L ← (PC); PC ← m
	JSUB Opcode = 0x48
	// A ← (m..m+2)
	LDA Opcode = 0x00
	// B ← (m..m+2)
	LDB Opcode = 0x68
	// A.low ← (m)
	LDCH Opcode = 0x50
	// F ← (m..m+5)
	LDF Opcode = 0x70
	// L ← (m..m+2)
	LDL Opcode = 0x08
	// S ← (m..m+2)
	LDS Opcode = 0x6C
	// T ← (m..m+2)
	LDT Opcode = 0x74
	// X ← (m..m+2)
	LDX Opcode = 0x04
	// PS ← (m..m+2)
	LPS Opcode = 0xD0
	// A ← (A) * (m..m+2)
	MUL Opcode = 0x20
	// F ← (F) * (m..m+5)
	MULF Opcode = 0x60
	// r2 ← (r2) * (r1)
	MULR Opcode = 0x98
	// F ← n
	NORM Opcode = 0xC8
	// A ← (A) | (m..m+2)
	OR Opcode = 0x44
	// A.low ← readdev(m)
	RD Opcode = 0xD8
	// (r2) ← (r1)
	RMO Opcode = 0xAC
	// PC ← (L)
	RSUB Opcode = 0x4C
	// (r1) ← (r1) << n
	SHIFTL Opcode = 0xA4
	// (r1) ← (r1) >> n
	SHIFTR Opcode = 0xA8
	// startio(A, S)
	SIO Opcode = 0xF0
	// SSK operation
	SSK Opcode = 0xEC
	// m..m+2 ← (A)
	STA Opcode = 0x0C
	// m..m+2 ← (B)
	STB Opcode = 0x78
	// m ← (A.low)
	STCH Opcode = 0x54
	// m..m+5 ← (F)
	STF Opcode = 0x80
	// timer ← (m..m+2)
	STI Opcode = 0xD4
	// m..m+2 ← (L)
	STL Opcode = 0x14
	// m..m+2 ← (S)
	STS Opcode = 0x7C
	// m..m+2 ← (SW)
	STSW Opcode = 0xE8
	// m..m+2 ← (T)
	STT Opcode = 0x84
	// m..m+2 ← (X)
	STX Opcode = 0x10
	// A ← (A) - (m..m+2)
	SUB Opcode = 0x1C
	// F ← (F) - (m..m+5)
	SUBF Opcode = 0x5C
	// r2 ← (r2) - (r1)
	SUBR Opcode = 0x94
	// interrupt(n)
	SVC Opcode = 0xB0
	// testdev(m)
	TD Opcode = 0xE0
	// testio(A)
	TIO Opcode = 0xF8
	// X ← (X) + 1; (X) : (m..m+2)
	TIX Opcode = 0x2C
	// X ← (X) + 1; (X) : (r)
	TIXR Opcode = 0xB8
	// writedev(m, A.low)
	WD Opcode = 0xDC
)

func isStoreInstruction(opcode byte) bool {
	switch Opcode(opcode) {
	case STA, STB, STCH, STF, STL, STS, STSW, STT, STX:
		return true
	}
	return false
}

func isJumpInstruction(opcode byte) bool {
	switch Opcode(opcode) {
	case J, JEQ, JGT, JLT, JSUB:
		return true
	}
	return false
}

//go:generate stringer -type=Opcode
