package vm

import (
	"bytes"
	"os"
	"testing"
)

func loadVMTestProgram(filename string) (*VM, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	vm := NewVM(reader)
	if err := vm.Load(); err != nil {
		return vm, err
	}
	return vm, nil
}

func runVMTestProgram(filename string) (*VM, error) {
	vm, err := loadVMTestProgram(filename)
	if err != nil {
		return vm, err
	}
	return vm, vm.Run()
}

func TestArithmetic(t *testing.T) {
	vm, err := runVMTestProgram("../programs/arith/arith.obj")
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

func TestRegisterArithmetic(t *testing.T) {
	vm, err := runVMTestProgram("../programs/arithr/arithr.obj")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// x=9, y=2
	sum, _ := vm.Memory.GetWord(0x35)
	diff, _ := vm.Memory.GetWord(0x38)
	prod, _ := vm.Memory.GetWord(0x3b)
	quot, _ := vm.Memory.GetWord(0x3e)
	mod, _ := vm.Memory.GetWord(0x41)

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

func TestFactorial(t *testing.T) {
	vm, err := runVMTestProgram("../programs/fakulteta/fakulteta.obj")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
		return
	}
	// fact(6) = 720
	value := vm.Registers.A
	if value != 720 {
		t.Errorf("Expected 720, got %d", value)
	}
}

func TestEcho(t *testing.T) {
	vm, _ := loadVMTestProgram("../programs/echo/echo.obj")
	buffer := new(bytes.Buffer)
	vm.SetDevice(1, NewOutputDevice(buffer))
	err := vm.Run()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
		return
	}
	output := buffer.String()
	expected := "12345\nhello world!\n"
	if output != expected {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}
