package vm

import (
	"testing"
)

func TestByteOperations(t *testing.T) {
	mem := Memory{}
	mem.SetByte(0, 42)
	mem.SetByte(51354, 123)

	value, err := mem.GetByte(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != 42 {
		t.Errorf("Expected 42, got %d", value)
	}

	value, err = mem.GetByte(51354)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != 123 {
		t.Errorf("Expected 123, got %d", value)
	}

	err = mem.SetByte(-1, 123)
	if err == nil {
		t.Error("Expected an error when setting an invalid address")
	}
	_, err = mem.GetByte(MAX_ADDRESS + 10)
	if err == nil {
		t.Error("Expected an error when getting an invalid address")
	}
}

func TestWordOperations(t *testing.T) {
	mem := Memory{}
	mem.SetWord(0, 0x123456)
	mem.SetWord(12345, -1)
	value, err := mem.GetWord(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != 0x123456 {
		t.Errorf("Expected 0x123456, got %x", value)
	}
	value, err = mem.GetWord(12345)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != -1 {
		t.Errorf("Expected -1, got %v", value)
	}
}

func TestFloatOperation(t *testing.T) {
	mem := Memory{}
	mem.SetFloat(0, 3.14159)

	value, err := mem.GetFloat(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != 3.14159 {
		t.Errorf("Expected 3.14159, got %v", value)
	}
}
