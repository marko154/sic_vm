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

type Memory struct {
	memory  [MAX_ADDRESS + 1]byte
	Changes []int32
}

const MAX_ADDRESS = (1 << 20) - 1

func (mem *Memory) ValidateAddress(address int32) error {
	if address < 0 || address > MAX_ADDRESS {
		return fmt.Errorf("invalid address: %d", address)
	}
	return nil
}

func (mem *Memory) GetByte(address int32) (byte, error) {
	if err := mem.ValidateAddress(address); err != nil {
		return 0, err
	}
	return mem.memory[address], nil
}

func (mem *Memory) SetByte(address int32, value byte) error {
	if err := mem.ValidateAddress(address); err != nil {
		return err
	}
	mem.memory[address] = value
	mem.Changes = []int32{address}
	return nil
}

func (mem *Memory) GetWord(address int32) (int32, error) {
	if err := mem.ValidateAddress(address + 2); err != nil {
		return 0, err
	}
	value := int32(mem.memory[address])<<16 | int32(mem.memory[address+1])<<8 | int32(mem.memory[address+2])
	// ??? is sign extension needed anywhere else ???
	return extendSign(value, 24), nil
	// return value, nil
}

func (mem *Memory) SetWord(address, value int32) error {
	if err := mem.ValidateAddress(address + 2); err != nil {
		return err
	}
	mem.memory[address] = byte(value >> 16)
	mem.memory[address+1] = byte(value >> 8)
	mem.memory[address+2] = byte(value)
	mem.Changes = []int32{address, address + 1, address + 2}
	return nil
}

// GetFloat reads a 48-bit floating point number from the memory. The floating point value is in the last 4 bytes of the 6 byte section
func (mem *Memory) GetFloat(address int32) (float32, error) {
	if err := mem.ValidateAddress(address + 5); err != nil {
		return 0, err
	}
	var value float32
	b := mem.memory[address+2 : address+6]
	buf := bytes.NewReader(b)
	// not sure if the endianess is correct
	if err := binary.Read(buf, binary.NativeEndian, &value); err != nil {
		return 0, err
	}
	return value, nil
}

func (mem *Memory) SetFloat(address int32, value float32) error {
	if err := mem.ValidateAddress(address + 5); err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.NativeEndian, &value); err != nil {
		return err
	}
	copy(mem.memory[address+2:], buf.Bytes())
	mem.Changes = []int32{address, address + 1, address + 2, address + 3, address + 4, address + 5}
	return nil
}
