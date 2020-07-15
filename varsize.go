package binutils

import (
	"fmt"
	"math"
)

const (
	// Use8bit defines size ID for list indexes having no more than 256 values.
	Use8bit BitsPerIndex = iota // uint8
	// Use16bit defines size ID for list indexes having no more than 65536 values.
	Use16bit // uint16
	// Use32bit defines size ID for list indexes having no more than 4294967296 values.
	Use32bit // uint32
	// Use64bit defines size ID for list indexes having no more than 1,8 x 10**19 values.
	Use64bit // uint64
	// Well-known values to marshal/unmarshal BitsPerIndex as binary:
	// bitsPerByte defines byte size in bits.
	bitsPerByte = 8
	// BytesInUint8 defines bytes amount required to store uint16 values.
	BytesInUint8 = 1
	// BytesInUint16 defines bytes amount required to store uint16 values.
	BytesInUint16 = 2
	// BytesInUint32 defines bytes amount required to store uint32 values.
	BytesInUint32 = 4
	// BytesInUint64 defines bytes amount required to store uint32 values.
	BytesInUint64 = 8
	// UsingUint8Indexes defines byte value to store as using uint8 indicator.
	UsingUint8Indexes = uint8(BytesInUint8 * bitsPerByte) // uint8
	// UsingUint16Indexes defines byte value to store as using uint16 indicator.
	UsingUint16Indexes = uint8(BytesInUint16 * bitsPerByte) // uint16
	// UsingUint32Indexes defines byte value to store as using uint32 indicator.
	UsingUint32Indexes = uint8(BytesInUint32 * bitsPerByte) // uint32
	// UsingUint64Indexes defines byte value to store as using uint64 indicator.
	UsingUint64Indexes = uint8(BytesInUint64 * bitsPerByte) // uint64
)

// BitsPerIndex defines index values size ID for numbered lists.
// Allows to detect and use minimal required standard type for list indexes when marshalling & unmarshalling.
// Only uint8 (byte), uint16 (2 bytes), uint32 (4 bytes) and uint64 (8 bytes) supported.
type BitsPerIndex int

// UnmarshalFromBuffer takes byte value from buffer and translates it into correct BitsPerIndex constant.
// Implements BufferUnmarshaler.
func (b *BitsPerIndex) UnmarshalFromBuffer(buffer *Buffer) error {
	b0 := new(uint8)
	if err := buffer.ReadUint8(b0); err != nil {
		return err
	}

	switch *b0 {
	case UsingUint8Indexes:
		*b = Use8bit
	case UsingUint16Indexes:
		*b = Use16bit
	case UsingUint32Indexes:
		*b = Use32bit
	case UsingUint64Indexes:
		*b = Use64bit
	default:
		return fmt.Errorf("%w: %d", ErrSizing, b0)
	}

	return nil
}

// MarshalBinary makes a byte representation of known type BitsPerIndex int constant.
// Returns 1 byte with 8, 16, 32 or 64 value or error if unknown BitsPerIndex marshalled.
func (b BitsPerIndex) MarshalBinary() (data []byte, err error) {
	switch b {
	case Use8bit:
		return []byte{UsingUint8Indexes}, nil
	case Use16bit:
		return []byte{UsingUint16Indexes}, nil
	case Use32bit:
		return []byte{UsingUint32Indexes}, nil
	case Use64bit:
		return []byte{UsingUint64Indexes}, nil
	default:
		return []byte{}, fmt.Errorf("%w: %d", ErrSizing, b)
	}
}

// UnmarshalBinary restores BitsPerIndex value from byte sequence.
// Requires exactly 1 byte with predefined values 8, 16, 32 or 64.
// Returns error if unexpected value supplied or data length not equals 1.
func (b *BitsPerIndex) UnmarshalBinary(data []byte) error {
	if len(data) > 1 {
		return fmt.Errorf("%w: got %d", ErrExpected1, len(data))
	}

	switch data[0] {
	case UsingUint8Indexes:
		*b = Use8bit
	case UsingUint16Indexes:
		*b = Use16bit
	case UsingUint32Indexes:
		*b = Use32bit
	case UsingUint64Indexes:
		*b = Use64bit
	default:
		return fmt.Errorf("%w: %d", ErrSizing, data[0])
	}

	return nil
}

// CalculateUseBitsPerIndex returns required index value size for requested slice length.
// If reserve not set, calculates size using all possible values.
// If reserve is true, calculates size using size bounds reserving 1 value
// for store nil as max possible value of this size or store nil as 0 to use 1-based index.
func CalculateUseBitsPerIndex(sliceLen int, reserveNil bool) (BitsPerIndex, error) {
	reserveValueForNil := 0
	if reserveNil {
		reserveValueForNil = 1
	}

	switch {
	case sliceLen < 0:
		return Use8bit, fmt.Errorf("%w: %d", ErrNegativeLen, sliceLen)
	case uint64(sliceLen) > math.MaxUint64-uint64(reserveValueForNil):
		return Use64bit, fmt.Errorf("%w: uint64 size < %d", ErrOverflowBy, sliceLen)
	case sliceLen > math.MaxUint32-reserveValueForNil:
		return Use64bit, nil
	case sliceLen > math.MaxUint16-reserveValueForNil:
		return Use32bit, nil
	case sliceLen > math.MaxUint8-reserveValueForNil:
		return Use16bit, nil
	default:
		return Use8bit, nil
	}
}

// WriteUint64ToBufferUsingBits writes uint64 value into buffer using only required bits.
// Returns written bytes count and error if any.
// Error not nil if value exceeds requested bits width max value.
func WriteUint64ToBufferUsingBits(value uint64, buffer *Buffer, usingBits BitsPerIndex) (int, error) {
	switch usingBits {
	case Use8bit:
		if value > math.MaxUint8 {
			return 0, fmt.Errorf("%w: uint8 < %d", ErrOverflowBy, value)
		}

		return buffer.WriteUint8(uint8(value))
	case Use16bit:
		if value > math.MaxUint16 {
			return 0, fmt.Errorf("%w: uint16 < %d", ErrOverflowBy, value)
		}

		return buffer.WriteUint16(uint16(value))
	case Use32bit:
		if value > math.MaxUint32 {
			return 0, fmt.Errorf("%w: uint32 < %d", ErrOverflowBy, value)
		}

		return buffer.WriteUint32(uint32(value))
	case Use64bit:
		return buffer.WriteUint64(value)
	default:
		return 0, fmt.Errorf("%w: %d", ErrSizing, usingBits)
	}
}

// ReadUint64FromBufferUsingBits reads value of requested bits wide from buffer into target uint64 pointer.
// Error will not not nil if unexpected bit width specified or read from buffer failed.
func ReadUint64FromBufferUsingBits(target *uint64, buffer *Buffer, usingBits BitsPerIndex) error {
	switch usingBits {
	case Use8bit:
		var value uint8
		if err := buffer.ReadUint8(&value); err != nil {
			return fmt.Errorf("%w: %v", ErrBuffer, err)
		}

		*target = uint64(value)
	case Use16bit:
		var value uint16
		if err := buffer.ReadUint16(&value); err != nil {
			return fmt.Errorf("%w: %v", ErrBuffer, err)
		}

		*target = uint64(value)
	case Use32bit:
		var value uint32
		if err := buffer.ReadUint32(&value); err != nil {
			return fmt.Errorf("%w: %v", ErrBuffer, err)
		}

		*target = uint64(value)
	case Use64bit:
		return buffer.ReadUint64(target)
	default:
		return fmt.Errorf("%w: %d", ErrSizing, usingBits)
	}

	return nil
}
