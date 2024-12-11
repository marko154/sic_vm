package vm

import (
	"os"
	"testing"
)

func runVMTestProgram(filename string) (*VM, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	vm := NewVM(reader)
	if err := vm.Load(); err != nil {
		return vm, err
	}
	return vm, vm.Run()
}

func TestArithmetic(t *testing.T) {
	vm, err := runVMTestProgram("../programs/arith.obj")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// x=9, y=2
	sum, _ := vm.Memory.GetWord(0x3f)
	diff, _ := vm.Memory.GetWord(0x42)
	prod, _ := vm.Memory.GetWord(0x45)
	quot, _ := vm.Memory.GetWord(0x48)
	mod, _ := vm.Memory.GetWord(0x4b)

	if sum != 11 {
		t.Errorf("Expected 11, got %d", sum)
	}
	if diff != 7 {
		t.Errorf("Expected 7, got %d", diff)
	}
	if prod != 18 {
		t.Errorf("Expected 18, got %d", prod)
	}
	if quot != 4 {
		t.Errorf("Expected 4, got %d", quot)
	}
	if mod != 1 {
		t.Errorf("Expected 1, got %d", mod)
	}
}
