package vm

import "fmt"

func notImplementedError(opcode Opcode) error {
	return fmt.Errorf("opcode %s not implemented", opcode.String())
}

func invalidOpcodeError(opcode byte) error {
	return fmt.Errorf("invalid opcode %d", opcode)
}

func zeroDivisionError() error {
	return fmt.Errorf("division by zero")
}
