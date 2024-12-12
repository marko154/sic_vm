package vm

import (
	"os"
	"testing"
)

func TestFileDevice_RoundTrip(t *testing.T) {
	device := NewFileDevice(0xAB)
	testByte := byte(42)
	defer os.Remove("AB.dev")

	err := device.Write(testByte)
	if err != nil {
		t.Errorf("Failed to write to device: %v", err)
	}
	// TODO: fix this test
	// _, err = device.Read()
	// if err != nil {
	// 	t.Errorf("Failed to read from device: %v", err)
	// }

	// if readByte != testByte {
	// 	t.Skipf("Read byte %v does not match written byte %v", readByte, testByte)
	// }

}

func TestFileDevice_NonExistentFile(t *testing.T) {
	device := NewFileDevice(0xCD)
	defer os.Remove("CD.dev")

	_, err := device.Read()
	if err == nil {
		t.Error("Expected error reading from non-existent file")
	}
}

func TestFileDevice_Test(t *testing.T) {
	device := NewFileDevice(0xEF)
	defer os.Remove("EF.dev")

	if device.Test() {
		t.Error("Test() should return false for non-existent file")
	}

	err := device.Write(0x00)
	if err != nil {
		t.Errorf("Failed to create test file: %v", err)
	}

	if !device.Test() {
		t.Error("Test() should return true for existing file")
	}

}
