package vm

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

type Device interface {
	Read() (byte, error)
	Write(byte) error
	Test() bool
}

type InputDevice struct {
	reader *bufio.Reader
}

func NewInputDevice(reader io.Reader) Device {
	return &InputDevice{reader: bufio.NewReader(reader)}
}

func (d *InputDevice) Read() (byte, error) {
	return d.reader.ReadByte()
}

func (d *InputDevice) Write(value byte) error {
	return nil
}

func (d *InputDevice) Test() bool {
	return true
}

type OutputDevice struct {
	writer *bufio.Writer
}

func NewOutputDevice(writer io.Writer) Device {
	return &OutputDevice{writer: bufio.NewWriter(writer)}
}

func (d *OutputDevice) Read() (byte, error) {
	return 0, nil
}

func (d *OutputDevice) Write(value byte) error {
	defer d.writer.Flush()
	return d.writer.WriteByte(value)
}
func (d *OutputDevice) Test() bool {
	return true
}

// Maps to a file on disk named XX.dev, from the hex value of the byte of the device number
// TODO: cleanup files after use?
type FileDevice struct {
	filename string
}

func NewFileDevice(name byte) Device {
	filename := strconv.FormatUint(uint64(name), 16) + ".dev"
	return &FileDevice{filename: filename}
}

func (d *FileDevice) Read() (byte, error) {
	file, err := os.Open(d.filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return bufio.NewReader(file).ReadByte()
}

func (d *FileDevice) Write(value byte) error {
	return os.WriteFile(d.filename, []byte{value}, 0644)
}
func (d *FileDevice) Test() bool {
	info, err := os.Stat(d.filename)
	return err == nil && !info.IsDir()
}
