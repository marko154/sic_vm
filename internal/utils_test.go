package vm

import "testing"

func TestExtendSign(t *testing.T) {
	value := extendSign(4095, 12)

	if value != -1 {
		t.Errorf("expected %v, got %v", -1, value)
	}

	value = extendSign(123, 12)
	if value != 123 {
		t.Errorf("expected %v, got %v", 123, value)
	}
}
