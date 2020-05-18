package binutils

import (
	"bytes"
	"encoding"
	"encoding/hex"
	"os"
	"path/filepath"
)

// Buffer type implements wrapper to easy marshalling & unmarshalling binary data.
// Defines some method to get info about stored data and marshaling/unmarshalling helpers.
//
// Binary Marshaling/Unmarshalling with buffer
//
// All write and read handlers takes additional argument stopIt and follow the rules:
// 1) If stopIt is error do nothing except stopIt as error
// 2) If stopIt is not nil, but not error generates new error
// It helps to make marshalling/unmarshalling chains like:
//
//   buffer := binaries.NewBuffer([]byte{})
//	 var err error
//	 _, err = buffer.WriteUint8(item1, err)
//	 _, err = buffer.WriteUint16(item2, err)
//	 _, err = buffer.WriteString(item3, err)
//	return buffer.Bytes(), err
//
// The same with unmarshalling, you can chain read operations using first error as stop indicator for following calls:
//
//  var err error
//	err = buffer.ReadUint8(&uin8Item, err)
//	err = buffer.ReadUint16(&uint16item, err)
//	err = buffer.ReadString(&stringItem, err)
//	return err
//
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

// NewEmptyBuffer is a shorthand to create new empty Buffer with binaries.NewBuffer([]byte{})
func NewEmptyBuffer() *Buffer {
	return &Buffer{buffer: bytes.NewBuffer(make([]byte, 0))}
}

// Len returns current length of buffer data in bytes
func (x *Buffer) Len() int {
	return x.buffer.Len()
}

// Bytes returns copy of current buffer data []bytes.
func (x *Buffer) Bytes() []byte {
	return append(make([]byte, 0), x.buffer.Bytes()...)
}

// WriteUint8 writes uint8 value into buffer as byte.
// Returns written bytes count and possible error.
func (x *Buffer) WriteUint8(data uint8, stopIt interface{}) (int, error) {
	if err, ok := stopIt.(error); ok {
		return 0, err
	} else if stopIt != nil {
		return 0, NewError("unexpected stopIt: %T %v", stopIt, stopIt)
	}
	return 1, x.buffer.WriteByte(data)
}

// ReadUint8 translates next byte from buffer into uint8 value and place it into target pointer.
// Returns nil or error
func (x *Buffer) ReadUint8(target *uint8, stop interface{}) error {
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if d, err := x.buffer.ReadByte(); err != nil {
		return err
	} else {
		*target = d
		return nil
	}
}

// WriteInt8 writes int8 value into buffer as byte.
// Returns written bytes count and possible error.
func (x *Buffer) WriteInt8(data int8, stopIt interface{}) (int, error) {
	if err, ok := stopIt.(error); ok {
		return 0, err
	} else if stopIt != nil {
		return 0, NewError("unexpected stopIt: %T %v", stopIt, stopIt)
	}
	return x.buffer.Write(Int8bytes(data))
}

// ReadInt8 translates next byte from buffer into int8 value and place it into target pointer.
// Returns nil or error
func (x *Buffer) ReadInt8(target *int8, stop interface{}) error {
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if v, err := Int8(x.buffer.Next(Int8size)); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteUint16 writes uint16 value into buffer using big-endian bytes order.
// Returns written bytes count and possible error.
func (x *Buffer) WriteUint16(data uint16, stopIt interface{}) (int, error) {
	if err, ok := stopIt.(error); ok {
		return 0, err
	} else if stopIt != nil {
		return 0, NewError("unexpected stopIt: %T %v", stopIt, stopIt)
	} else {
		return x.buffer.Write(Uint16bytes(data))
	}

}

// ReadUint16 translates next 2 bytes from buffer into uint16 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadUint16(target *uint16, stop interface{}) error {
	d := AllocateBytes(Uint16size)
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if consumedBytes, err := x.buffer.Read(d); err != nil {
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
func (x *Buffer) WriteInt16(data int16, stopIt interface{}) (int, error) {
	if err, ok := stopIt.(error); ok {
		return 0, err
	} else if stopIt != nil {
		return 0, NewError("unexpected stopIt: %T %v", stopIt, stopIt)
	}
	return x.buffer.Write(Int16bytes(data))
}

// ReadInt16 translates next 2 bytes from buffer into int16 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadInt16(target *int16, stop interface{}) error {
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if v, err := Int16(x.buffer.Next(Int16size)); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteUint32 writes uint32 value into buffer using big-endian bytes order.
// Returns written bytes count and possible error.
func (x *Buffer) WriteUint32(data uint32, stopIt interface{}) (int, error) {
	if err, ok := stopIt.(error); ok {
		return 0, err
	} else if stopIt != nil {
		return 0, NewError("unexpected stopIt: %T %v", stopIt, stopIt)
	} else {
		return x.buffer.Write(Uint32bytes(data))
	}
}

// ReadUint32 translates next 4 bytes from buffer into uint32 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadUint32(target *uint32, stop interface{}) error {
	d := AllocateBytes(Uint32size)
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if consumedBytes, err := x.buffer.Read(d); err != nil {
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
func (x *Buffer) WriteInt32(data int32, stopIt interface{}) (int, error) {
	if err, ok := stopIt.(error); ok {
		return 0, err
	} else if stopIt != nil {
		return 0, NewError("unexpected stopIt: %T %v", stopIt, stopIt)
	}
	return x.buffer.Write(Int32bytes(data))
}

// ReadInt32 translates next 4 bytes from buffer into int32 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadInt32(target *int32, stop interface{}) error {
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if v, err := Int32(x.buffer.Next(Unt32size)); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteUint64 writes uint64 value into buffer using big-endian bytes order.
// Returns written bytes count and possible error.
func (x *Buffer) WriteUint64(data uint64, stopIt interface{}) (int, error) {
	if err, ok := stopIt.(error); ok {
		return 0, err
	} else if stopIt != nil {
		return 0, NewError("unexpected stopIt: %T %v", stopIt, stopIt)
	} else {
		return x.buffer.Write(Uint64bytes(data))
	}

}

// ReadUint64 translates next 4 bytes from buffer into uint64 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadUint64(target *uint64, stop interface{}) error {
	d := AllocateBytes(Uint64size)
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if consumedBytes, err := x.buffer.Read(d); err != nil {
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
func (x *Buffer) WriteInt64(data int64, stopIt interface{}) (int, error) {
	if err, ok := stopIt.(error); ok {
		return 0, err
	} else if stopIt != nil {
		return 0, NewError("unexpected stopIt: %T %v", stopIt, stopIt)
	}
	return x.buffer.Write(Int64bytes(data))
}

// ReadInt64 translates next 4 bytes from buffer into int64 value and place it into target pointer.
// It uses big-endian byte order.
// Returns nil or error
func (x *Buffer) ReadInt64(target *int64, stop interface{}) error {
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if v, err := Int64(x.buffer.Next(Int64size)); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteString adds binary representation of string as zero-terminated string.
func (x *Buffer) WriteString(data string, stop interface{}) (int, error) {
	if err, ok := stop.(error); ok {
		return 0, err
	} else if stop != nil {
		return 0, NewError("unexpected stop: %T %v", stop, stop)
	}
	return x.buffer.Write(StringBytes(data))
}

// ReadString reads zero-terminated string from buffer.
func (x *Buffer) ReadString(target *string, stop interface{}) error {
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if line, err := x.buffer.ReadBytes(0); err != nil {
		return err
	} else if v, err := String(line); err != nil {
		return err
	} else {
		*target = v
		return nil
	}
}

// WriteBytes adds data from byte slice into buffer.
// Returns written bytes count and nil or possible error
func (x *Buffer) WriteBytes(data []byte, stop interface{}) (int, error) {
	if err, ok := stop.(error); ok {
		return 0, err
	} else if stop != nil {
		return 0, NewError("unexpected stop: %T %v", stop, stop)
	}
	return x.buffer.Write(data)
}

// ReadBytes takes required amount of bytes from buffer into target byte slice pointer.
// Returns nil or possible error
func (x *Buffer) ReadBytes(target *[]byte, numBytes int, stop interface{}) error {
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if x.buffer.Len() < numBytes {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else {
		d := x.buffer.Next(numBytes)
		*target = append(*target, d...)
		return nil
	}
}

// WriteObject takes encoding.BinaryMarshaler and fill underlying data buffer with provided bytes.
// Returns written bytes count and possible error
func (x *Buffer) WriteObject(data encoding.BinaryMarshaler, stop interface{}) (int, error) {
	if err, ok := stop.(error); ok {
		return 0, err
	} else if stop != nil {
		return 0, NewError("unexpected stop: %T %v", stop, stop)
	}
	d, err := data.MarshalBinary()
	if err != nil {
		return 0, err
	}
	return x.buffer.Write(d)

}

// ReadObject provides expected bytes count into encoding.BinaryUnmarshaler implementations UnmarshalBinary method.
// It uses same interface as another read methods in buffer.
// Returns nil or possible error
func (x *Buffer) ReadObject(data encoding.BinaryUnmarshaler, bytes int, stop interface{}) error {
	var objectBytes []byte
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if x.buffer.Len() < bytes {
		return NewError(
			"required %d bytes, have only %d: %v",
			bytes, x.buffer.Len(), hex.EncodeToString(x.buffer.Bytes()),
		)
	} else if err := x.ReadBytes(&objectBytes, bytes, nil); err != nil {
		return err
	} else if err := data.UnmarshalBinary(objectBytes); err != nil {
		return err
	} else {
		return nil
	}
}

// UnmarshalObject passes buffer itself as argument to BufferUnmarshaler instance. Returns nil or possible error.
func (x *Buffer) UnmarshalObject(data BufferUnmarshaler, stop interface{}) error {
	if err, ok := stop.(error); ok {
		return err
	} else if stop != nil {
		return NewError("unexpected stop: %T %v", stop, stop)
	} else if err := data.UnmarshalFromBuffer(x); err != nil {
		return err
	} else {
		return nil
	}
}

// MarshalBinary implementing binary.BinaryMarshaler for buffer itself.
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

// WriteFromFile loads additional bytes from file.
// Bytes will appended to the end of current data.
// NOTE: If buffer is not empty, it will not overwritten but extended with file data
func (x *Buffer) WriteFromFile(filePath string) (int, error) {
	if absFileName, err := filepath.Abs(filePath); err != nil {
		return 0, err
	} else if state, err := os.Stat(absFileName); err != nil {
		return 0, err
	} else if reader, err := os.Open(absFileName); err != nil {
		return 0, err
	} else if bytesRed, err := x.buffer.ReadFrom(reader); err != nil {
		return 0, err
	} else if bytesRed != state.Size() {
		return int(bytesRed), NewError("red only %d bytes of %d", bytesRed, state.Size())
	} else {
		return int(bytesRed), nil
	}
}

// ReadIntoFile unloads buffer data into binary file.
// Target file will be created even if buffer is empty.
func (x *Buffer) ReadIntoFile(filePath string) (int, error) {
	if absFileName, err := filepath.Abs(filePath); err != nil {
		return 0, err
	} else if writer, err := os.Create(absFileName); err != nil {
		return 0, err
	} else {
		bytesWritten, err := x.buffer.WriteTo(writer)
		return int(bytesWritten), err
	}
}
