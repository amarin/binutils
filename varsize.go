package binutils

import (
	"math"
)

// BitsPerIndex defines index values size ID for numbered lists.
// Allows to detect and use minimal required standard type for list indexes when marshalling & unmarshalling.
// Only uint8 (byte), uint16 (2 bytes), uint32 (4 bytes) and uint64 (8 bytes) supported.
type BitsPerIndex int

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
		return []byte{}, NewError("unexpected bits per index value `%v`", b)
	}
} // nolint:gofmt

// UnmarshalBinary restores BitsPerIndex value from byte sequence.
// Requires exactly 1 byte with predefined values 8, 16, 32 or 64.
// Returns error if unexpected value supplied or data length not equals 1.
func (b *BitsPerIndex) UnmarshalBinary(data []byte) error {
	if len(data) > 1 {
		return NewError("expected 1 byte, not %d", len(data))
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
		return NewError("unexpected size byte %d", data[0])
	}

	return nil
}

const (
	// Predefined BitsPerIndex values

	// Use8bit defines size ID for list indexes having no more than 256 values.
	Use8bit BitsPerIndex = iota // uint8
	// Use16bit defines size ID for list indexes having no more than 65536 values.
	Use16bit // uint16
	// Use32bit defines size ID for list indexes having no more than 4294967296 values.
	Use32bit // uint32
	// Use64bit defines size ID for list indexes having no more than 1,8 x 10**19 values.
	Use64bit // uint64

	// Well-known values to marshal/unmarshal BitsPerIndex as binary:

	// UsingUint8Indexes defines byte value to store as using uint8 indicator.
	UsingUint8Indexes = uint8(8) // nolint:gomnd    // uint8
	// UsingUint16Indexes defines byte value to store as using uint16 indicator.
	UsingUint16Indexes = uint8(16) // nolint:gomnd    // uint16
	// UsingUint32Indexes defines byte value to store as using uint32 indicator.
	UsingUint32Indexes = uint8(32) // nolint:gomnd    // uint32
	// UsingUint64Indexes defines byte value to store as using uint64 indicator.
	UsingUint64Indexes = uint8(64) // nolint:gomnd    // uint64
)

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
		return Use8bit, NewError("cant detect required size by negative len %d", sliceLen)
	case uint64(sliceLen) > math.MaxUint64-uint64(reserveValueForNil):
		return Use64bit, NewError("len %d too big even for uint64", sliceLen)
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
