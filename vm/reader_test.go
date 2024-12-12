package vm

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadByte(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("23")) // 0010 0011
	reader := NewReader(r)
	value, err := reader.ReadByte()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != 35 {
		t.Errorf("Expected 35, got %d", value)
	}
}

func TestReadString(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("some text here to test"))
	reader := NewReader(r)
	value, err := reader.ReadString(4)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != "some" {
		t.Errorf("Expected 'some', got %s", value)
	}
}

func TestReadWord(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("0102a8")) // 0000 00010000001010101000
	reader := NewReader(r)
	value, err := reader.ReadWord()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != 66216 {
		t.Errorf("Expected 66216, got %d", value)
	}
}
