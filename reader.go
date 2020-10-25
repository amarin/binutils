package binutils

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// BinaryWriter implements binary writing for various data types into file writer.
type BinaryReader struct {
	source io.Reader
}

// OpenFile opens specified file path and returns BinaryReader wrapping it.
// Target filePath must exists and readable before opening.
func OpenFile(filePath string) (*BinaryReader, error) {
	if absFileName, err := filepath.Abs(filePath); err != nil {
		return nil, err
	} else if source, err := os.Open(absFileName); err != nil {
		return nil, err
	} else {
		return &BinaryReader{source: source}, nil
	}
}

// NewBinaryReader wraps existing io.Reader into BinaryReader.
func NewBinaryReader(source io.Reader) *BinaryReader {
	return &BinaryReader{source: source}
}

// Close closes underlying reader. Implements io.Closer.
// Returns error if underlying reader not  implements io.Closer.
func (r BinaryReader) Close() error {
	closer, ok := r.source.(io.Closer)
	if !ok {
		return fmt.Errorf("%w: %T is not io.Closer", ErrClose, r.source)
	}

	return closer.Close()
}

// Read reads up to to len(p) bytes into p. It returns the number of bytes read (0 <= n <= len(p))
// and any error encountered, just calling underlying io.Reader Read method.
// Implements io.Reader itself.
func (r BinaryReader) Read(p []byte) (n int, err error) {
	return r.source.Read(p)
}

// ReadBytesCount reads exactly specified amount of bytes.
// Returns read bytes or error if insufficient bytes count ready to read or any underlying buffer error encountered.
func (r BinaryReader) ReadBytesCount(amount int) ([]byte, error) {
	buffer := make([]byte, amount)
	bytesRead, err := r.Read(buffer)

	switch {
	case err != nil:
		return buffer, err
	case bytesRead != amount:
		return buffer, fmt.Errorf("%w: expected %v read %v", io.EOF, amount, bytesRead)
	default:
		return buffer, nil
	}
}

// ReadUint8 reads uint8 value from underlying reader.
// Returns uint8 value and any error encountered.
func (r BinaryReader) ReadUint8() (uint8, error) {
	byteBuffer := AllocateBytes(Uint8size)
	bytesRead, err := r.source.Read(byteBuffer)

	switch {
	case err != nil:
		return 0, err
	case bytesRead != Uint8size:
		return 0, fmt.Errorf("%w: expected %v read %v", io.ErrUnexpectedEOF, Uint8size, bytesRead)
	default:
		return Uint8(byteBuffer)
	}
}

// ReadUint16 reads uint16 value from underlying reader.
// Returns uint16 value and any error encountered.
func (r BinaryReader) ReadUint16() (uint16, error) {
	byteBuffer := AllocateBytes(Uint16size)
	bytesRead, err := r.source.Read(byteBuffer)

	switch {
	case err != nil:
		return 0, err
	case bytesRead != Uint16size:
		return 0, fmt.Errorf("%w: expected %v read %v", io.ErrUnexpectedEOF, Uint16size, bytesRead)
	default:
		return Uint16(byteBuffer)
	}
}

// ReadUint32 reads uint32 value from underlying reader.
// Returns uint32 value and any error encountered.
func (r BinaryReader) ReadUint32() (uint32, error) {
	byteBuffer := AllocateBytes(Uint32size)
	bytesRead, err := r.source.Read(byteBuffer)

	switch {
	case err != nil:
		return 0, err
	case bytesRead != Uint32size:
		return 0, fmt.Errorf("%w: expected %v read %v", io.ErrUnexpectedEOF, Uint32size, bytesRead)
	default:
		return Uint32(byteBuffer)
	}
}

// ReadUint64 reads uint64 value from underlying reader.
// Returns uint64 value and any error encountered.
func (r BinaryReader) ReadUint64() (uint64, error) {
	byteBuffer := AllocateBytes(Uint64size)
	bytesRead, err := r.source.Read(byteBuffer)

	switch {
	case err != nil:
		return 0, err
	case bytesRead != Uint64size:
		return 0, fmt.Errorf("%w: expected %v read %v", io.ErrUnexpectedEOF, Uint64size, bytesRead)
	default:
		return Uint64(byteBuffer)
	}
}

// ReadUint reads uint value from underlying reader.
// Returns uint value and any error encountered.
func (r BinaryReader) ReadUint() (uint, error) {
	uint64result, err := r.ReadUint64()
	return uint(uint64result), err
}

// ReadUint8 reads int8 value from underlying reader.
// Returns int8 value and any error encountered.
func (r BinaryReader) ReadInt8() (int8, error) {
	byteBuffer := AllocateBytes(Int8size)
	bytesRead, err := r.source.Read(byteBuffer)

	switch {
	case err != nil:
		return 0, err
	case bytesRead != Int8size:
		return 0, fmt.Errorf("%w: expected %v read %v", io.ErrUnexpectedEOF, Int8size, bytesRead)
	default:
		return Int8(byteBuffer)
	}
}

// ReadUint16 reads int16 value from underlying reader.
// Returns int16 value and any error encountered.
func (r BinaryReader) ReadInt16() (int16, error) {
	byteBuffer := AllocateBytes(Int16size)
	bytesRead, err := r.source.Read(byteBuffer)

	switch {
	case err != nil:
		return 0, err
	case bytesRead != Int16size:
		return 0, fmt.Errorf("%w: expected %v read %v", io.ErrUnexpectedEOF, Int16size, bytesRead)
	default:
		return Int16(byteBuffer)
	}
}

// ReadUint32 reads int32 value from underlying reader.
// Returns int32 value and any error encountered.
func (r BinaryReader) ReadInt32() (int32, error) {
	byteBuffer := AllocateBytes(Int32size)
	bytesRead, err := r.source.Read(byteBuffer)

	switch {
	case err != nil:
		return 0, err
	case bytesRead != Int32size:
		return 0, fmt.Errorf("%w: expected %v read %v", io.ErrUnexpectedEOF, Int32size, bytesRead)
	default:
		return Int32(byteBuffer)
	}
}

// ReadUint64 reads int64 value from underlying reader.
// Returns int64 value and any error encountered.
func (r BinaryReader) ReadInt64() (int64, error) {
	byteBuffer := AllocateBytes(Int64size)
	bytesRead, err := r.source.Read(byteBuffer)

	switch {
	case err != nil:
		return 0, err
	case bytesRead != Int64size:
		return 0, fmt.Errorf("%w: expected %v read %v", io.ErrUnexpectedEOF, Int64size, bytesRead)
	default:
		return Int64(byteBuffer)
	}
}

// ReadInt reads int value from underlying reader.
// Returns int value and any error encountered.
func (r BinaryReader) ReadInt() (int, error) {
	int64result, err := r.ReadInt64()
	return int(int64result), err
}

// ReadBytes reads bytes sequence until the first occurrence of stop byte in the input.
// Returns a bytes slice containing the data up to and including the delimiter.
// If ReadBytes encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
func (r BinaryReader) ReadBytes(stop byte) ([]byte, error) {
	alreadyImplemented, ok := r.source.(untilStopByteReader)
	if ok {
		return alreadyImplemented.ReadBytes(stop)
	}
	// underlying reader does not implement read bytes until stop,
	// so read byte-by-byte and compare next ones until stop byte found or any read error happened.
	buffer := make([]byte, 0)

	for {
		currentByte, err := r.ReadUint()
		if err != nil {
			return buffer, err
		}

		buffer = append(buffer, byte(currentByte))

		if byte(currentByte) == stop {
			return buffer, nil
		}
	}
}

// ReadStringZ reads zero-terminated string from underlying reader.
func (r BinaryReader) ReadStringZ() (string, error) {
	line, err := r.ReadBytes(0)
	if err != nil {
		return "", fmt.Errorf("%w: read: %v", ErrRequired0T, err)
	}
	// read zero-terminated line successfully.
	return string(line), nil
}

// ReadHex reads exactly specified amount of bytes and return hex representation string for received bytes.
// Returns underlying reader errors encountered.
func (r BinaryReader) ReadHex(amount int) (string, error) {
	data, err := r.ReadBytesCount(amount)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(data), nil
}

// WriteObject add encoding.BinaryMarshaler binary data into buffer.
// Returns written bytes count and possible error.
func (r BinaryReader) ReadObject(target interface{}) error {
	switch binaryReaderFrom := target.(type) {
	case BinaryReaderFrom:
		_, err := binaryReaderFrom.BinaryReadFrom(r)
		if err != nil {
			return err
		}

		return nil

	case io.ReaderFrom:
		_, err := binaryReaderFrom.ReadFrom(r)
		if err != nil {
			return err
		}

		return nil

	default:
		return fmt.Errorf("%w: should implement io.ReaderFrom or binutils.BinaryReaderFrom", ErrRead)
	}
}
