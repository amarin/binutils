package binutils

// BufferUnmarshaler defines interface to objects able to unmarshal itself from binary buffer.
// Such objects should read only own bytes from buffer leaving extra bytes intact for others.
type BufferUnmarshaler interface {
	UnmarshalFromBuffer(*Buffer) error
}

type untilStopByteReader interface {
	// ReadBytes until the first occurrence of stopByte in the input,
	// returning a slice containing the data up to and including the delimiter.
	// If ReadBytes encounters an error before finding a delimiter,
	// it returns the data read before the error and the error itself (often io.EOF).
	// ReadBytes returns err != nil if and only if the returned data does not end in
	// stopByte.
	ReadBytes(stopByte byte) ([]byte, error)
}

// BinaryReaderFrom is the interface that wraps the BinaryReadFrom method.
//
// BinaryReadFrom reads implementation data from binary until its data restored or any error encountered.
// The return value n is the number of bytes taken from reader.
// Any error except io.EOF encountered during the read is also returned.
type BinaryReaderFrom interface {
	BinaryReadFrom(*BinaryReader) (n int64, err error)
}

// BinaryWriterTo is the interface that wraps the BinaryWriteTo method.
//
// BinaryWriteTo writes implementors data into writer until marshalled or any error occurs.
// The return value n is the number of bytes
// written. Any error encountered during writing is also returned.
type BinaryWriterTo interface {
	BinaryWriteTo(*BinaryWriter) (n int64, err error)
}
