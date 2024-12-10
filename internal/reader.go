package vm

import (
	"bufio"
	"io"
	"strconv"
)

type ByteReader interface {
	ReadByte() (byte, error)
	ReadString(n int) (string, error)
	ReadWord() (int, error)
}

type Reader struct {
	reader *bufio.Reader
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{reader: bufio.NewReader(reader)}
}

func (r *Reader) ReadByte() (byte, error) {
	buf := make([]byte, 2)
	_, err := r.reader.Read(buf)
	if err != nil {
		return 0, err
	}
	// this is wrong use strconv.parseuint(txt, 16, 8)
	parsedValue, err := strconv.ParseUint(string(buf), 16, 8)
	if err != nil {
		return 0, err
	}
	return byte(parsedValue), nil
}

func (r *Reader) ReadString(n int) (string, error) {
	buf := make([]byte, n)
	_, err := r.reader.Read(buf)
	return string(buf), err
}

func (r *Reader) ReadWord() (int, error) {
	b1, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	b2, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	b3, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	word := int(b1)<<16 | int(b2)<<8 | int(b3)
	return word, err
}
