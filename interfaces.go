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
// Returns any error encountered during reading if happened or nil.
type BinaryReaderFrom interface {
	BinaryReadFrom(*BinaryReader) error
}

// BinaryWriterTo interface wraps the BinaryWriteTo method.
// Implementation method BinaryWriteTo writes implementors data into BinaryWriter
// until all marshalled or any error occurs.
// Returns any error encountered during writing if happened or nil.
type BinaryWriterTo interface {
	BinaryWriteTo(*BinaryWriter) error
}
