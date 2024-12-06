package vm

import (
	"bufio"
	"io"
)

type VM struct {
	Memory       Memory
	Registers    Registers
	Devices      map[int]Device // device number -> device
	Reader       ByteReader
	ProgramName  string
	CodeAddress  int
	StartAddress int
	CodeLength   int
}

func NewVM() *VM {
	return &VM{}
}

// !!! warning: sign extension when using operands !!!

// TODO: separate the loader from the VM
func (vm *VM) Load(reader *bufio.Reader) error {
	vm.Reader = NewReader(reader)
	for {
		recordType, err := reader.ReadByte()
		if err == io.EOF {
			break
		}

		switch recordType {
		case 'H':
			err = vm.readHRecord()
		case 'T':
			err = vm.readTRecord()
		case 'E':
			err = vm.readERecord()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (vm *VM) readHRecord() error {
	value, err := vm.Reader.ReadString(6)
	if err != nil {
		return err
	}
	vm.ProgramName = value
	address, err := vm.Reader.ReadWord()
	if err != nil {
		return err
	}
	vm.StartAddress = address
	codeLength, err := vm.Reader.ReadWord()
	if err != nil {
		return err
	}
	vm.CodeLength = codeLength
	return nil
}

func (vm *VM) readERecord() error {
	address, err := vm.Reader.ReadWord()
	if err != nil {
		return err
	}
	vm.StartAddress = address
	return nil
}

func (vm *VM) readTRecord() error {
	address, err := vm.Reader.ReadWord()
	if err != nil {
		return err
	}
	// length of the code in bytes (2 * n nibbles or characters in obj file)
	length, err := vm.Reader.ReadByte()
	if err != nil {
		return err
	}
	for offset := 0; offset < int(length); offset++ {
		value, err := vm.Reader.ReadByte()
		if err != nil {
			return err
		}
		vm.Memory.SetByte(address+offset, value)
	}

	return nil
}

func (vm *VM) Run() {
}
