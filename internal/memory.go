package vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type AddressingMode int

const (
	IMMEDIATE AddressingMode = 0b01
	DIRECT    AddressingMode = 0b11
	INDIRECT  AddressingMode = 0b10
	SIC       AddressingMode = 0b00
)

type Memory [MAX_ADDRESS]byte

const MAX_ADDRESS = 1 << 20

func (mem *Memory) ValidateAddress(address int) error {
	if address < 0 || address >= MAX_ADDRESS {
		return fmt.Errorf("invalid address: %d", address)
	}
	return nil
}

func (mem *Memory) GetByte(address int) (byte, error) {
	if err := mem.ValidateAddress(address); err != nil {
		return 0, err
	}
	return mem[address], nil
}

func (mem *Memory) SetByte(address int, value byte) error {
	if err := mem.ValidateAddress(address); err != nil {
		return err
	}
	mem[address] = value
	return nil
}

func (mem *Memory) GetWord(address int) (int, error) {
	if err := mem.ValidateAddress(address + 2); err != nil {
		return 0, err
	}
	return int(mem[address])<<16 | int(mem[address+1])<<8 | int(mem[address+2]), nil
}

func (mem *Memory) SetWord(address, value int) error {
	if err := mem.ValidateAddress(address + 2); err != nil {
		return err
	}
	mem[address] = byte(value >> 16)
	mem[address+1] = byte(value >> 8)
	mem[address+2] = byte(value)
	return nil
}

// GetFloat reads a 48-bit floating point number from the memory. The floating point value is in the last 4 bytes of the 6 byte section
func (mem *Memory) GetFloat(address int) (float32, error) {
	if err := mem.ValidateAddress(address + 5); err != nil {
		return 0, err
	}
	var value float32
	b := mem[address+2 : address+6]
	buf := bytes.NewReader(b)
	// not sure if the endianess is correct
	if err := binary.Read(buf, binary.NativeEndian, &value); err != nil {
		return 0, err
	}
	return value, nil
}

func (mem *Memory) SetFloat(address int, value float32) error {
	if err := mem.ValidateAddress(address + 5); err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.NativeEndian, &value); err != nil {
		return err
	}
	copy(mem[address+2:], buf.Bytes())
	return nil
}
