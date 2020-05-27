package binutils

import (
	"bytes"
	"encoding"
	"os"
	"path/filepath"
)

// Buffer type implements wrapper to easy marshalling & unmarshalling binary data.
// Defines some method to get info about stored data and marshaling/unmarshalling helpers.
type Buffer struct {
	buffer *bytes.Buffer
}

// NewBuffer is a default constructor to create Buffer. It requires data argument to init underlying data.
// To make new empty buffer use:
//   buffer := binaries.NewBuffer([]byte{})
//
// To make new buffer with predefined data of []byte use:
//   buffer := binaries.NewBuffer(data)
func NewBuffer(data []byte) *Buffer {
	// create underlying buffer using detached slice
	return &Buffer{buffer: bytes.NewBuffer(append(make([]byte, 0), data...))}
}

// NewEmptyBuffer is a shorthand to create new empty Buffer with binaries.NewBuffer([]byte{}).
func NewEmptyBuffer() *Buffer {
	return &Buffer{buffer: bytes.NewBuffer(make([]byte, 0))}
}

// Len returns current length of buffer data in bytes.
func (x *Buffer) Len() int {
	return x.buffer.Len()
}

// Bytes returns copy of current buffer data []bytes.
func (x *Buffer) Bytes() []byte {
	return append(make([]byte, 0), x.buffer.Bytes()...)
}

// WriteUint8 writes uint8 value into buffer as byte.
// Returns written bytes count and possible error.
func (x *Buffer) WriteUint8(data uint8) (int, error) {
	return 1, x.buffer.WriteByte(data)
}

// ReadUint8 translates next byte from buffer into uint8 value and place it into target pointer.
// Returns nil or error.
func (x *Buffer) ReadUint8(target *uint8) error {
	if d, err := x.buffer.ReadByte(); err != nil {
		return err
	} else {
		*target = d
		return nil
	}
}

// WriteInt8 writes int8 value into buffer as byte.
// Returns written bytes count and possible error.
func (x *Buffer) WriteInt8(data int8) (int, error) {
	return x.buffer.Write(Int8bytes(data))
}

// ReadInt8 translates next byte from buffer into int8 value and place it into target pointer.
// Returns nil or error.
func (x *Buffer) ReadInt8(target *int8) error {
	if v, err := Int8(x.buffer.Next(Int8size)); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteUint16 writes uint16 value into buffer using big-endian bytes order.
// Returns written bytes count and possible error.
func (x *Buffer) WriteUint16(data uint16) (int, error) {
	return x.buffer.Write(Uint16bytes(data))
}

// ReadUint16 translates next 2 bytes from buffer into uint16 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadUint16(target *uint16) error {
	d := AllocateBytes(Uint16size)
	if consumedBytes, err := x.buffer.Read(d); err != nil {
		return err
	} else if Uint16size != consumedBytes {
		return NewError("Expected %d bytes, only %d consumed", Uint16size, consumedBytes)
	} else if v, err := Uint16(d); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteInt16 writes int16 value into buffer using big-endian bytes order.
// Returns written bytes count and possible error.
func (x *Buffer) WriteInt16(data int16) (int, error) {
	return x.buffer.Write(Int16bytes(data))
}

// ReadInt16 translates next 2 bytes from buffer into int16 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadInt16(target *int16) error {
	if v, err := Int16(x.buffer.Next(Int16size)); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteUint32 writes uint32 value into buffer using big-endian bytes order.
// Returns written bytes count and possible error.
func (x *Buffer) WriteUint32(data uint32) (int, error) {
	return x.buffer.Write(Uint32bytes(data))
}

// ReadUint32 translates next 4 bytes from buffer into uint32 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadUint32(target *uint32) error {
	d := AllocateBytes(Uint32size)
	if consumedBytes, err := x.buffer.Read(d); err != nil {
		return err
	} else if Uint32size != consumedBytes {
		return NewError("Expected %d bytes, only %d consumed", Uint32size, consumedBytes)
	} else if v, err := Uint32(d); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteInt32 writes int32 value into buffer using big-endian bytes order.
// Returns written bytes count and possible error.
func (x *Buffer) WriteInt32(data int32) (int, error) {
	return x.buffer.Write(Int32bytes(data))
}

// ReadInt32 translates next 4 bytes from buffer into int32 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadInt32(target *int32) error {
	if v, err := Int32(x.buffer.Next(Unt32size)); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteUint64 writes uint64 value into buffer using big-endian bytes order.
// Returns written bytes count and possible error.
func (x *Buffer) WriteUint64(data uint64) (int, error) {
	return x.buffer.Write(Uint64bytes(data))
}

// ReadUint64 translates next 4 bytes from buffer into uint64 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadUint64(target *uint64) error {
	d := AllocateBytes(Uint64size)
	if consumedBytes, err := x.buffer.Read(d); err != nil {
		return err
	} else if Uint64size != consumedBytes {
		return NewError("Expected %d bytes, only %d consumed", Uint64size, consumedBytes)
	} else if v, err := Uint64(d); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteInt64 writes int64 value into buffer using big-endian bytes order.
// Returns written bytes count and possible error.
func (x *Buffer) WriteInt64(data int64) (int, error) {
	return x.buffer.Write(Int64bytes(data))
}

// ReadInt64 translates next 4 bytes from buffer into int64 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadInt64(target *int64) error {
	if v, err := Int64(x.buffer.Next(Int64size)); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteString adds binary representation of string as zero-terminated string.
func (x *Buffer) WriteString(data string) (int, error) {
	return x.buffer.Write(StringBytes(data))
}

// ReadString reads zero-terminated string from buffer.
func (x *Buffer) ReadString(target *string) error {
	if line, err := x.buffer.ReadBytes(0); err != nil {
		return err
	} else if v, err := String(line); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteBytes adds data from byte slice into buffer.
// Returns written bytes count and nil or possible error.
func (x *Buffer) WriteBytes(data []byte) (int, error) {
	return x.buffer.Write(data)
}

// ReadBytes takes required amount of bytes from buffer into target byte slice pointer.
// Returns nil or possible error
func (x *Buffer) ReadBytes(target *[]byte, numBytes int) error {
	if x.buffer.Len() < numBytes {
		return NewError("buffer len %d less then required %d", x.buffer.Len(), numBytes)
	} else {
		d := x.buffer.Next(numBytes)
		*target = append(*target, d...)
		return nil
	}
}

// WriteObject add encoding.BinaryMarshaler binary data into buffer.
// Returns written bytes count and possible error.
func (x *Buffer) WriteObject(data encoding.BinaryMarshaler) (int, error) {
	d, err := data.MarshalBinary()
	if err != nil {
		return 0, err
	}

	return x.buffer.Write(d)
}

// ReadObjectBytes provides expected bytes count into encoding.BinaryUnmarshaler implementations UnmarshalBinary method.
// It uses same interface as another read methods in buffer.
// Returns nil or possible error.
func (x *Buffer) ReadObjectBytes(data encoding.BinaryUnmarshaler, bytes int) error {
	var objectBytes []byte

	if x.buffer.Len() < bytes {
		return NewError("required %d bytes, buffer len %d", bytes, x.buffer.Len())
	} else if err := x.ReadBytes(&objectBytes, bytes); err != nil {
		return WrapError(err, "cant read %d bytes", bytes)
	} else if err := data.UnmarshalBinary(objectBytes); err != nil {
		return WrapError(err, "cant unmarshal %T", data)
	}

	return nil
}

// ReadObject allows BufferUnmarshaler instances to take its bytes themselves. Returns nil or possible error.
func (x *Buffer) ReadObject(data BufferUnmarshaler) error {
	if err := data.UnmarshalFromBuffer(x); err != nil {
		return err
	}

	return nil
}

// MarshalBinary implementing binary.BinaryMarshaler for Buffer.
// Simply returns copy of underlying data with always nil error.
func (x *Buffer) MarshalBinary() (data []byte, err error) {
	return x.Bytes(), nil
}

// UnmarshalBinary implements binary.BinaryUnmarshaler.
// Silently replaces underlying data with new data.
func (x *Buffer) UnmarshalBinary(data []byte) error {
	x.buffer = bytes.NewBuffer(data)
	return nil
}

// LoadFromFilePath loads additional bytes from file.
// Bytes will appended to the end of current data.
// NOTE: If buffer is not empty, it will not overwritten but extended with file data.
func (x *Buffer) LoadFromFilePath(filePath string) (int, error) {
	absFileName, err := filepath.Abs(filePath)
	if err != nil {
		return 0, WrapError(err, "cant detect absolute file path")
	}
	// stat file
	state, err := os.Stat(absFileName)
	if err != nil {
		return 0, WrapError(err, "cant stat file path")
	}

	reader, err := os.Open(absFileName)
	if err != nil {
		return 0, WrapError(err, "cant open file for reading")
	}

	bytesRed, err := x.buffer.ReadFrom(reader)
	if err != nil {
		return 0, WrapError(err, "cant read file")
	}

	if bytesRed != state.Size() {
		return int(bytesRed), NewError("red only %d bytes of %d", bytesRed, state.Size())
	}

	return int(bytesRed), nil
}

// SaveIntoFilePath unloads buffer data into binary file.
// Target file will be created even if buffer is empty.
func (x *Buffer) SaveIntoFilePath(filePath string) (int, error) {
	if absFileName, err := filepath.Abs(filePath); err != nil {
		return 0, err
	} else if writer, err := os.Create(absFileName); err != nil {
		return 0, err
	} else {
		bytesWritten, err := x.buffer.WriteTo(writer)
		return int(bytesWritten), err
	}
}
