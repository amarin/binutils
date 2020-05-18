package binutils

// Utility functions to translate native types into sequence of bytes and vise versa.
import (
	"encoding/binary"
)

// AllocateBytes creates a byte slice of required size
func AllocateBytes(size int) []byte {
	return make([]byte, size)
}

// Uint8 translates next byte from buffer into uint8 value.
// Returns error if insufficient bytes in buffer
func Uint8(data []byte) (uint8, error) {
	if len(data) != Uint8size {
		return 0, NewError("Expected %d bytes", Uint8size)
	}
	return data[0], nil
}

// Int8 translates next byte from buffer into int8 value.
// Returns error if insufficient bytes in buffer
func Int8(data []byte) (int8, error) {
	if len(data) != Int8size {
		return 0, NewError("Expected %d bytes", Int8size)
	}
	return int8(data[0]), nil
}

// Uint16 translates next 2 bytes from buffer into uint16 value using big-endian bytes order.
// Returns error if insufficient bytes in buffer
func Uint16(data []byte) (uint16, error) {
	if len(data) != Uint16size {
		return 0, NewError("Expected %d bytes", Uint16size)
	}
	return binary.BigEndian.Uint16(data), nil
}

// Int16 translates next 2 bytes from buffer into int16 value using big-endian bytes order.
// Returns error if insufficient bytes in buffer
func Int16(data []byte) (int16, error) {
	if len(data) != Int16size {
		return 0, NewError("Expected %d bytes", Int16size)
	}
	return int16(binary.BigEndian.Uint16(data)), nil
}

// Uint32 translates next 4 bytes from buffer into uint32 value using big-endian bytes order.
// Returns error if insufficient bytes in buffer
func Uint32(data []byte) (uint32, error) {
	if len(data) != Uint32size {
		return 0, NewError("Expected %d bytes", Uint32size)
	}
	return binary.BigEndian.Uint32(data), nil
}

// Int32 translates next 4 bytes from buffer into int32 value using big-endian bytes order.
// Returns error if insufficient bytes in buffer
func Int32(data []byte) (int32, error) {
	if len(data) != Unt32size {
		return 0, NewError("Expected %d bytes", Unt32size)
	}
	return int32(binary.BigEndian.Uint32(data)), nil
}

// Uint64 translates next 8 bytes from buffer into uint64 value using big-endian bytes order.
// Returns error if insufficient bytes in buffer
func Uint64(data []byte) (uint64, error) {
	if len(data) != Uint64size {
		return 0, NewError("Expected %d bytes", Uint64size)
	}
	return binary.BigEndian.Uint64(data), nil
}

// Int64 translates next 8 bytes from buffer into int64 value using big-endian bytes order.
// Returns error if insufficient bytes in buffer
func Int64(data []byte) (int64, error) {
	if len(data) != Int64size {
		return 0, NewError("Expected %d bytes", Int64size)
	}
	return int64(binary.BigEndian.Uint64(data)), nil
}

// Uint8bytes adds uint8 data to buffer
func Uint8bytes(data uint8) []byte {
	return []byte{data}
}

// Int8bytes adds int8 data to buffer
func Int8bytes(data int8) []byte {
	return []byte{uint8(data)}
}

// Uint16bytes adds uint16 data to buffer using big-endian bytes order
func Uint16bytes(data uint16) []byte {
	d := AllocateBytes(2)
	binary.BigEndian.PutUint16(d, data)
	return d
}

// Int16bytes adds int16 data to buffer using big-endian bytes order
func Int16bytes(data int16) []byte {
	d := AllocateBytes(Int16size)
	binary.BigEndian.PutUint16(d, uint16(data))
	return d
}

// Uint32bytes adds uint32 data to buffer using big-endian bytes order
func Uint32bytes(data uint32) []byte {
	d := AllocateBytes(4)
	binary.BigEndian.PutUint32(d, data)
	return d
}

// Int32bytes adds int32 data to buffer using big-endian bytes order
func Int32bytes(data int32) []byte {
	d := AllocateBytes(Uint32size)
	binary.BigEndian.PutUint32(d, uint32(data))
	return d
}

// Uint64bytes adds uint64 data to buffer using big-endian bytes order
func Uint64bytes(data uint64) []byte {
	d := AllocateBytes(Uint64size)
	binary.BigEndian.PutUint64(d, data)
	return d
}

// Int64bytes adds uint64 data to buffer using big-endian bytes order
func Int64bytes(data int64) []byte {
	d := AllocateBytes(Uint64size)
	binary.BigEndian.PutUint64(d, uint64(data))
	return d
}

// StringBytes makes a zero-terminated string []byte sequence
func StringBytes(s string) []byte {
	return append([]byte(s), 0)
}

// String reads a zero-terminated string from []byte sequence
// Returns error if last byte is not 0
func String(data []byte) (string, error) {
	if len(data) == 0 {
		return "", NewError("required at least 1 byte for 0-terminated string")
	} else if data[len(data)-1] != 0 {
		return "", NewError("required 0-terminated string")
	} else {
		return string(data[:len(data)-1]), nil
	}
}
