package vm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const MAX_DEVICES = 255

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
type FileDevice struct {
	filename string
	file     *os.File
	reader   *bufio.Reader
}

func NewFileDevice(name byte) Device {
	fname := strings.ToUpper(strconv.FormatUint(uint64(name), 16))
	return &FileDevice{filename: fname + ".dev"}
}

func (d *FileDevice) initFile() error {
	var err error
	if d.file == nil {
		d.file, err = os.OpenFile(d.filename, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		d.reader = bufio.NewReader(d.file)
	}
	return nil
}

func (d *FileDevice) Read() (byte, error) {
	err := d.initFile()
	if err != nil {
		return 0, err
	}
	return d.reader.ReadByte()
}

func (d *FileDevice) Write(value byte) error {
	err := d.initFile()
	if err != nil {
		return err
	}
	if _, err := d.file.Write([]byte{value}); err != nil {
		return fmt.Errorf("failed to write to file %s: %v", d.filename, err)
	}
	return d.file.Sync()
}

func (d *FileDevice) Test() bool {
	info, err := os.Stat(d.filename)
	return err == nil && !info.IsDir()
}

func (d *FileDevice) Close() error {
	if d.file != nil {
		err := d.file.Close()
		return err
	}
	return nil
}
