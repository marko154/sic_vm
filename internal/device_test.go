package vm

import (
	"os"
	"testing"
)

// TODO: use afero to mock filesystem instead of actually creating files

func TestFileDevice_RoundTrip(t *testing.T) {
	device := NewFileDevice(0xAB)
	testByte := byte(42)
	defer os.Remove("ab.dev")

	err := device.Write(testByte)
	if err != nil {
		t.Errorf("Failed to write to device: %v", err)
	}

	readByte, err := device.Read()
	if err != nil {
		t.Errorf("Failed to read from device: %v", err)
	}

	if readByte != testByte {
		t.Errorf("Read byte %v does not match written byte %v", readByte, testByte)
	}

}

func TestFileDevice_NonExistentFile(t *testing.T) {
	device := NewFileDevice(0xCD)

	_, err := device.Read()
	if err == nil {
		t.Error("Expected error reading from non-existent file")
	}
}

func TestFileDevice_Test(t *testing.T) {
	device := NewFileDevice(0xEF)
	defer os.Remove("ef.dev")

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
