package binutils

type untilStopByteReader interface {
	// ReadBytes until the first occurrence of stopByte in the input,
	// returning a slice containing the data up to and including the delimiter.
	// If ReadBytes encounters an error before finding a delimiter,
	// it returns the data read before the error and the error itself (often io.EOF).
	// ReadBytes returns err != nil if and only if the returned data does not end in
	// stopByte.
	ReadBytes(stopByte byte) ([]byte, error)
}

// BinaryReaderFrom interface wraps the BinaryReadFrom method.
// Implementation method BinaryReadFrom reads implementors data from BinaryReader
// until its data restored or any error encountered.
// The return value n is the number of bytes taken from reader.
// Any error except io.EOF encountered during the read is also returned.
type BinaryReaderFrom interface {
	BinaryReadFrom(*BinaryReader) (n int64, err error)
}

// BinaryWriterTo interface wraps the BinaryWriteTo method.
// Implementation method BinaryWriteTo writes implementors data into BinaryWriter
// until all marshalled or any error occurs.
// Returns any error encountered during writing if happened or nil.
type BinaryWriterTo interface {
	BinaryWriteTo(*BinaryWriter) error
}

// BinaryUint8 requires implementation could translate its value to uint8 type.
// Used in unified BinaryWriter.WriteObject method.
type BinaryUint8 interface{ Uint8() uint8 }

// BinaryUint16 requires implementation could translate its value to uint16 type.
// Used in unified BinaryWriter.WriteObject method.
type BinaryUint16 interface{ Uint16() uint16 }

// BinaryUint32 requires implementation could translate its value to uint32 type.
// Used in unified BinaryWriter.WriteObject method.
type BinaryUint32 interface{ Uint32() uint32 }

// BinaryUint64 requires implementation could translate its value to uint64 type.
// Used in unified BinaryWriter.WriteObject method.
type BinaryUint64 interface{ Uint64() uint64 }

// BinaryInt8 requires implementation could translate its value to int8 type.
// Used in unified BinaryWriter.WriteObject method.
type BinaryInt8 interface{ Int8() int8 }

// BinaryInt16 requires implementation could translate its value to int16 type.
// Used in unified BinaryWriter.WriteObject method.
type BinaryInt16 interface{ Int16() int16 }

// BinaryInt32 requires implementation could translate its value to int32 type.
// Used in unified BinaryWriter.WriteObject method.
type BinaryInt32 interface{ Int32() int32 }

// BinaryInt64 requires implementation could translate its value to int64 type.
// Used in unified BinaryWriter.WriteObject method.
type BinaryInt64 interface{ Int64() int64 }

// BinaryRune requires implementation could translate its value to rune type.
// Used in unified BinaryWriter.WriteObject method.
type BinaryRune interface{ Rune() rune }

// BinaryString requires implementation could translate its value to string suitable to serialize as binary.
// Used in unified BinaryWriter.WriteObject method.
type BinaryString interface{ BinaryString() string }
