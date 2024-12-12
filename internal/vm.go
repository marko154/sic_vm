package vm

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type VM struct {
	Memory       Memory
	Registers    *Registers
	Devices      map[int32]Device // device number -> device
	Loader       *Loader
	ProgramName  string
	CodeAddress  int32
	StartAddress int32
	CodeLength   int32
}

func NewVM(reader io.Reader) *VM {
	vm := &VM{
		Loader:    NewLoader(reader),
		Devices:   make(map[int32]Device),
		Registers: NewRegisters(),
	}
	vm.SetDevice(0, NewInputDevice(os.Stdin))
	vm.SetDevice(1, NewOutputDevice(os.Stdout))
	vm.SetDevice(2, NewOutputDevice(os.Stderr))
	for i := int32(3); i <= MAX_DEVICES; i++ {
		vm.SetDevice(i, NewFileDevice(byte(i)))
	}
	return vm
}

func (vm *VM) Load() error {
	return vm.Loader.Load(vm)
}

// TODO: bounds check?
func (vm *VM) GetDevice(num int32) Device {
	device := vm.Devices[num]
	return device
}

func (vm *VM) SetDevice(num int32, device Device) {
	vm.Devices[num] = device
}

// !!! warning: sign extension when using operands !!! ?

func (vm *VM) Run() error {
	vm.Registers.PC = vm.StartAddress

	for {
		prevPC := vm.Registers.PC
		opcode, err := vm.fetch()
		if err != nil {
			return err
		}
		if err := vm.execute(opcode); err != nil {
			return err
		}
		if prevPC == vm.Registers.PC {
			break // HALT
		}
	}
	return nil
}

func (vm *VM) fetch() (byte, error) {
	value, err := vm.Memory.GetByte(vm.Registers.PC)
	if err != nil {
		return 0, err
	}
	vm.Registers.PC++
	return value, nil
}

func (vm *VM) execute(opcodeByte byte) error {
	opcode := opcodeByte & 0xFC
	log.Debugf("executing opcode %v\n", Opcode(opcode))
	ni := opcodeByte & 0x03
	if executed, err := vm.tryExecuteF1(opcode); executed || err != nil {
		return err
	}
	operand, err := vm.fetch()
	if err != nil {
		return err
	}
	if executed, err := vm.tryExecuteF2(opcode, operand); executed || err != nil {
		return err
	}
	return vm.tryExecuteTypeSICF3F4(opcode, ni, operand)
}

func (vm *VM) tryExecuteF1(opcode byte) (bool, error) {
	switch Opcode(opcode) {
	case FIX:
		vm.Registers.A = int32(vm.Registers.F)
	case FLOAT:
		vm.Registers.F = float64(vm.Registers.A)
	case HIO:
		return true, notImplementedError(HIO)
	case NORM:
		return true, notImplementedError(NORM) // what does this even do?
	case SIO:
		return true, notImplementedError(SIO)
	case TIO:
		return true, notImplementedError(TIO)
	default:
		return false, nil
	}
	return true, nil
}
func (vm *VM) tryExecuteF2(opcode, operands byte) (bool, error) {
	r1 := int(operands >> 4)
	r2 := int(operands & 0x0F)
	switch Opcode(opcode) {
	case ADDR:
		vm.Registers.SetReg(r2, vm.Registers.GetReg(r2)+vm.Registers.GetReg(r1))
	case CLEAR:
		vm.Registers.SetReg(r1, 0)
	case COMPR:
		r1Val := vm.Registers.GetReg(r1)
		r2Val := vm.Registers.GetReg(r2)
		cc := getConditionCodes(r1Val, r2Val)
		vm.Registers.SetCC(cc)
	case DIVR:
		r1Val := vm.Registers.GetReg(r1)
		if r1Val == 0 {
			return false, zeroDivisionError()
		}
		vm.Registers.SetReg(r2, vm.Registers.GetReg(r2)/r1Val)
	case MULR:
		vm.Registers.SetReg(r2, vm.Registers.GetReg(r2)*vm.Registers.GetReg(r1))
	case RMO:
		vm.Registers.SetReg(r2, vm.Registers.GetReg(r1))
	case SHIFTL:
		shiftAmount := vm.Registers.GetReg(r2)
		vm.Registers.SetReg(r1, vm.Registers.GetReg(r1)<<shiftAmount)
	case SHIFTR:
		shiftAmount := vm.Registers.GetReg(r2)
		vm.Registers.SetReg(r1, vm.Registers.GetReg(r1)>>shiftAmount)
	case SUBR:
		vm.Registers.SetReg(r2, vm.Registers.GetReg(r2)-vm.Registers.GetReg(r1))
	case TIXR:
		vm.Registers.X++
		r := vm.Registers.GetReg(r1)
		vm.Registers.SetCC(getConditionCodes(vm.Registers.X, r))
	case SVC:
		return false, notImplementedError(SVC)
	default:
		return false, nil
	}
	return true, nil
}

// load, arithmetic, device instructions resolve the address the same way
// store instructions have one level of indirection less
// jump instructions have one level of indirection less
func (vm *VM) tryExecuteTypeSICF3F4(opcode, ni, operand byte) error {
	effectiveAddress, err := vm.getEffectiveAddress(ni, operand)
	if err != nil {
		return err
	}
	operandValue, err := vm.getOperandValue(opcode, ni, effectiveAddress)
	if err != nil {
		return err
	}

	switch Opcode(opcode) {
	// arithmetic/logic/simple instructions
	case ADD:
		vm.Registers.A += operandValue
	case AND:
		vm.Registers.A &= operandValue
	case DIV:
		if operandValue == 0 {
			return zeroDivisionError()
		}
		vm.Registers.A /= operandValue
	case MUL:
		vm.Registers.A *= operandValue
	case SUB:
		vm.Registers.A -= operandValue
	case OR:
		vm.Registers.A |= operandValue
	case TIX:
		vm.Registers.X++
		vm.Registers.SetCC(getConditionCodes(vm.Registers.X, operandValue))
	case COMP:
		cc := getConditionCodes(vm.Registers.A, operandValue)
		vm.Registers.SetCC(cc)
	// jump instructions - only immediate addressing supported
	case J:
		vm.Registers.PC = effectiveAddress
	case JEQ:
		if vm.Registers.GetCC() == CC_EQ {
			vm.Registers.PC = effectiveAddress
		}
	case JGT:
		if vm.Registers.GetCC() == CC_GT {
			vm.Registers.PC = effectiveAddress
		}
	case JLT:
		if vm.Registers.GetCC() == CC_LT {
			vm.Registers.PC = effectiveAddress
		}
	case JSUB:
		vm.Registers.L = vm.Registers.PC
		vm.Registers.PC = effectiveAddress
	case RSUB:
		vm.Registers.PC = vm.Registers.L
	// load instructions
	case LDA:
		vm.Registers.A = operandValue
	case LDB:
		vm.Registers.B = operandValue
	case LDCH:
		vm.Registers.A = (vm.Registers.A & (-256)) | (operandValue & 0xFF)
	case LDL:
		vm.Registers.L = operandValue
	case LDS:
		vm.Registers.S = operandValue
	case LDT:
		vm.Registers.T = operandValue
	case LDX:
		vm.Registers.X = operandValue
	// store instructions
	case STA:
		vm.Memory.SetWord(operandValue, vm.Registers.A)
	case STB:
		vm.Memory.SetWord(operandValue, vm.Registers.B)
	case STCH:
		vm.Memory.SetByte(operandValue, byte(vm.Registers.A)) // TODO: fix
	case STL:
		vm.Memory.SetWord(operandValue, vm.Registers.L)
	case STS:
		vm.Memory.SetWord(operandValue, vm.Registers.S)
	case STSW:
		vm.Memory.SetWord(operandValue, vm.Registers.SW)
	case STT:
		vm.Memory.SetWord(operandValue, vm.Registers.T)
	case STX:
		vm.Memory.SetWord(operandValue, vm.Registers.X)
	// device instructions
	case RD:
		value, err := vm.GetDevice(operandValue).Read()
		if err != nil {
			return err
		}
		vm.Registers.A = (vm.Registers.A & (-256)) | int32(value&0xFF)
	case TD:
		vm.GetDevice(operandValue).Test()
	case WD:
		device := vm.GetDevice(operandValue)
		if err := device.Write(byte(vm.Registers.A)); err != nil {
			return err
		}

	// floating point instructions
	case ADDF:
		return notImplementedError(ADDF)
	case DIVF:
		return notImplementedError(DIVF)
	case LDF:
		return notImplementedError(LDF)
	case MULF:
		return notImplementedError(MULF)
	case STF:
		return notImplementedError(STF)
	case SUBF:
		return notImplementedError(SUBF)
	case COMPF:
		return notImplementedError(ADDF)
	// I have no idea what these are
	case LPS:
		return notImplementedError(LPS)
	case SSK:
		return notImplementedError(SSK)
	case STI:
		return notImplementedError(STI)
	default:
		return invalidOpcodeError(opcode)
	}
	return nil
}

func (vm *VM) getEffectiveAddress(ni, operand byte) (int32, error) {
	third, err := vm.fetch()
	if err != nil {
		return 0, err
	}
	mask := byte(0x0F)
	if AddressingMode(ni) == SIC {
		mask = 0x7F
	}
	offset := int32((operand&mask))<<8 + int32(third)
	addrMode := getAddressCalcMode((operand >> 4) & 0x0F)
	// only extend sign when address is pc relative (who's fucking idea was this?)
	if addrMode.P {
		offset = extendSign(offset, 12)
	}
	if addrMode.E {
		fourth, err := vm.fetch()
		if err != nil {
			return 0, err
		}
		offset = (offset << 8) + int32(fourth)
	}
	var effectiveAddress int32 = 0
	if addrMode.B {
		effectiveAddress += vm.Registers.B
	}
	if addrMode.P {
		effectiveAddress += vm.Registers.PC
	}
	if addrMode.X {
		effectiveAddress += vm.Registers.X
	}
	effectiveAddress += offset
	return effectiveAddress, nil
}

type AddressCalculationMode struct {
	X bool // index addressing
	B bool // base addressing
	P bool // PC relative
	E bool // extended format
}

func getAddressCalcMode(xbpe byte) AddressCalculationMode {
	return AddressCalculationMode{
		X: xbpe&0b1000 != 0,
		B: xbpe&0b0100 != 0,
		P: xbpe&0b0010 != 0,
		E: xbpe&0b0001 != 0,
	}
}

// TODO: refactor
func (vm *VM) getOperandValue(opcode, ni byte, effectiveAddress int32) (int32, error) {
	indirectionLevel := getIndirectionLevel(opcode, ni)
	log.Debugf("indirection level %v\n", indirectionLevel)
	return vm.resolveAddress(effectiveAddress, indirectionLevel)
}

func getIndirectionLevel(opcode, ni byte) int {
	level := 0
	switch AddressingMode(ni) {
	case IMMEDIATE:
		level = 0
	case DIRECT, SIC:
		level = 1
	case INDIRECT:
		level = 2
	}
	if isStoreInstruction(opcode) || isJumpInstruction(opcode) {
		level--
	}
	return level
}

func (vm *VM) resolveAddress(effectiveAddress int32, indirectionLevel int) (int32, error) {
	for i := 0; i < indirectionLevel; i++ {
		value, err := vm.Memory.GetWord(effectiveAddress)
		if err != nil {
			return 0, err
		}
		effectiveAddress = value
	}
	return effectiveAddress, nil
}
