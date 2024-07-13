package resp

import (
	"io"
	"strconv"
)

type Writer struct {
	writer io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w}
}

func (v Value) Marshal() []byte {
	switch v.typ {
	case "array":
		return v.marshalArray()
	case "bulk":
		return v.marshalBulk()
	case "error":
		return v.marshalError()
	case "null":
		return v.marshalNull()
	case "string":
		return v.marshalString()
	default:
		return []byte{}
	}
}

func (v Value) marshalArray() []byte {
	len := len(v.array)
	var bytes []byte
	bytes = append(bytes, ARRAY)
	bytes = append(bytes, []byte(strconv.Itoa(len))...)
	bytes = append(bytes, '\r', '\n')

	for i := 0; i < len; i++ {
		bytes = append(bytes, v.array[i].Marshal()...)
	}

	return bytes
}

func (v Value) marshalBulk() []byte {
	len := len(v.bulk)
	var bytes []byte
	bytes = append(bytes, BULK)
	bytes = append(bytes, []byte(strconv.Itoa(len))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, []byte(v.bulk)...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalNull() []byte {
	return []byte("$-1\r\n")
}

func (v Value) marshalError() []byte {
	var bytes []byte
	bytes = append(bytes, ERROR)
	bytes = append(bytes, []byte(v.str)...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshalString() []byte {
	var bytes []byte
	bytes = append(bytes, STRING)
	bytes = append(bytes, []byte(v.str)...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (w *Writer) Write(v Value) error {
	bytes := v.Marshal()

	_, err := w.writer.Write(bytes)

	if err != nil {
		return err
	}
	return nil
}
