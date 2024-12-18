package vm

import (
	"bufio"
	"io"
	"strconv"
)

type ByteReader interface {
	ReadByte() (byte, error)
	ReadString(n int) (string, error)
	ReadWord() (int32, error)
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

func (r *Reader) ReadWord() (int32, error) {
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

	word := int32(b1)<<16 | int32(b2)<<8 | int32(b3)
	return extendSign(word, 24), nil
	// TODO: should we extend sign here? i hate this fucking thing
	// return word, err
}
