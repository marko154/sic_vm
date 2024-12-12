package vm

import (
	"fmt"
	"io"
)

type Loader struct {
	Reader ByteReader
}

func NewLoader(reader io.Reader) *Loader {
	return &Loader{Reader: NewReader(reader)}
}

// TODO: refactor this, Loader doesn't need the entire VM struct
func (l *Loader) Load(vm *VM) error {
	for {
		recordType, err := l.Reader.ReadString(1)
		if err == io.EOF {
			break
		}

		switch recordType {
		case "H":
			err = l.readHRecord(vm)
		case "T":
			err = l.readTRecord(vm)
		case "E":
			err = l.readERecord(vm)
		case "M":
			err = l.readMRecord()
		default:
			return fmt.Errorf("record type %v not implemented", recordType)
		}
		if err != nil {
			return err
		}
		// skip newline
		if _, err := l.Reader.ReadString(1); err != nil {
			return fmt.Errorf("failed to read newline: %v", err)
		}
	}

	return nil
}

func (l *Loader) readHRecord(vm *VM) error {
	value, err := l.Reader.ReadString(6)
	if err != nil {
		return err
	}
	vm.ProgramName = value
	address, err := l.Reader.ReadWord()
	if err != nil {
		return err
	}
	vm.StartAddress = address
	codeLength, err := l.Reader.ReadWord()
	if err != nil {
		return err
	}
	vm.CodeLength = codeLength
	return nil
}

func (l *Loader) readERecord(vm *VM) error {
	address, err := l.Reader.ReadWord()
	if err != nil {
		return err
	}
	vm.StartAddress = address
	return nil
}

func (l *Loader) readTRecord(vm *VM) error {
	address, err := l.Reader.ReadWord()
	if err != nil {
		return err
	}
	// length of the code in bytes (2 * n nibbles or characters in obj file)
	length, err := l.Reader.ReadByte()
	if err != nil {
		return err
	}
	for offset := int32(0); offset < int32(length); offset++ {
		value, err := l.Reader.ReadByte()
		if err != nil {
			return err
		}
		vm.Memory.SetByte(address+offset, value)
	}

	return nil
}

func (l *Loader) readMRecord() error {
	_, err := l.Reader.ReadWord()
	if err != nil {
		return err
	}
	_, err = l.Reader.ReadByte()
	if err != nil {
		return err
	}
	// TODO: peek at the next byte to determine if it's a plus or minus and process the rest
	return nil
}
