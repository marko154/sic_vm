package vm

import (
	"bytes"
	"testing"
)

func TestLoad(t *testing.T) {
	// vm := NewVM()
	// vm.Load("layout.obj") // TODO: create a simple file to test if it loads correctly into memory

	// assert that memory
}

func TestGetEffectiveAddress(t *testing.T) {
	reader := bytes.NewReader([]byte{0x10, 0x00})
	vm := NewVM(reader)
	vm.Registers.B = 123
	vm.Registers.X = 0x100
	vm.Memory.SetByte(0, byte(LDA)|byte(IMMEDIATE))
	address, err := vm.getEffectiveAddress(0b11, 0x10)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if address != 0x1100 {
		t.Errorf("Expected 0x1100, got %x", address)
	}
}
