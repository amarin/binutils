package binutils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	// ErrWriter identifies any writer errors.
	ErrWriter = errors.New("writer")

	// ErrWriterWrite identifies writer failed during write.
	ErrWriterWrite = fmt.Errorf("%w: write", ErrWriter)
)

// BinaryWriter implements binary writing for various data types into file writer.
type BinaryWriter struct {
	writer io.Writer
}

// CreateFile creates file and wrap file writer into BinaryWriter.
// Target file will created.
func CreateFile(filePath string) (*BinaryWriter, error) {
	w := &BinaryWriter{}

	if absFileName, err := filepath.Abs(filePath); err != nil {
		return nil, fmt.Errorf("%v: resolve path: %w", ErrWriter, err)
	} else if w.writer, err = os.Create(absFileName); err != nil {
		return nil, fmt.Errorf("%v: create file: %w", ErrWriter, err)
	} else {
		return w, nil
	}
}

// NewBinaryWriter wraps existing io.Writer instance into BinaryWriter.
func NewBinaryWriter(writer io.Writer) *BinaryWriter {
	return &BinaryWriter{writer: writer}
}

// Close closes underlying writer if it implements io.Closer.
// Returns error if underlying writer is not implements io.Closer.
func (w BinaryWriter) Close() error {
	closer, ok := w.writer.(io.Closer)
	if !ok {
		return fmt.Errorf("%w: %T is not io.Closer", ErrClose, w.writer)
	}

	return closer.Close()
}

// Write simply writes into underlying writer.
// Implements io.Writer.
func (w BinaryWriter) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

// WriteUint8 writes uint8 value into writer as bytes.
func (w BinaryWriter) WriteUint8(data uint8) error {
	bytesWritten, err := w.writer.Write(Uint8bytes(data))

	switch {
	case err != nil:
		return fmt.Errorf("%v: uint8: %w", ErrWriterWrite, err)
	case bytesWritten != Uint8size:
		return fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, Uint8size, bytesWritten)
	}

	return nil
}

// WriteUint16 writes uint16 value into writer as bytes.
func (w BinaryWriter) WriteUint16(data uint16) error {
	bytesWritten, err := w.writer.Write(Uint16bytes(data))

	switch {
	case err != nil:
		return fmt.Errorf("%v: uint16: %w", ErrWriterWrite, err)
	case bytesWritten != Uint16size:
		return fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, Uint16size, bytesWritten)
	}

	return nil
}

// WriteUint32 writes uint16 value into writer as bytes.
func (w BinaryWriter) WriteUint32(data uint32) error {
	bytesWritten, err := w.writer.Write(Uint32bytes(data))

	switch {
	case err != nil:
		return fmt.Errorf("%v: uint32: %w", ErrWriterWrite, err)
	case bytesWritten != Uint32size:
		return fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, Uint32size, bytesWritten)
	}

	return nil
}

// WriteUint64 writes uint64 value into writer as bytes.
func (w BinaryWriter) WriteUint64(data uint64) error {
	bytesWritten, err := w.writer.Write(Uint64bytes(data))

	switch {
	case err != nil:
		return fmt.Errorf("%v: uint64: %w", ErrWriterWrite, err)
	case bytesWritten != Uint64size:
		return fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, Uint64size, bytesWritten)
	}

	return nil
}

// WriteUint uint value into writer as bytes.
func (w BinaryWriter) WriteUint(data uint) error {
	return w.WriteUint64(uint64(data))
}

// WriteInt8 writes int8 value into writer as byte.
func (w BinaryWriter) WriteInt8(data int8) error {
	bytesWritten, err := w.writer.Write(Int8bytes(data))

	switch {
	case err != nil:
		return fmt.Errorf("%v: int8: %w", ErrWriterWrite, err)
	case bytesWritten != Int8size:
		return fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, Int8size, bytesWritten)
	}

	return nil
}

// WriteInt16 writes int16 value into writer as bytes.
func (w BinaryWriter) WriteInt16(data int16) error {
	bytesWritten, err := w.writer.Write(Int16bytes(data))

	switch {
	case err != nil:
		return fmt.Errorf("%v: int16: %w", ErrWriterWrite, err)
	case bytesWritten != Int16size:
		return fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, Int16size, bytesWritten)
	}

	return nil
}

// WriteInt32 writes int32 value into writer as bytes.
func (w BinaryWriter) WriteInt32(data int32) error {
	bytesWritten, err := w.writer.Write(Int32bytes(data))

	switch {
	case err != nil:
		return fmt.Errorf("%v: int32: %w", ErrWriterWrite, err)
	case bytesWritten != Int32size:
		return fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, Int32size, bytesWritten)
	}

	return nil
}

// WriteInt64 writes int64 value into writer as bytes.
func (w BinaryWriter) WriteInt64(data int64) error {
	bytesWritten, err := w.writer.Write(Int64bytes(data))

	switch {
	case err != nil:
		return fmt.Errorf("%v: int64: %w", ErrWriterWrite, err)
	case bytesWritten != Int64size:
		return fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, Int64size, bytesWritten)
	}

	return nil
}

// WriteInt int value into writer as bytes.
func (w BinaryWriter) WriteInt(data int) error {
	return w.WriteInt64(int64(data))
}

// WriteStringZ writes string bytes into underlying writer as Zero-terminated string.
func (w BinaryWriter) WriteStringZ(data string) error {
	stringBytes := StringBytes(data)
	bytesWritten, err := w.writer.Write(stringBytes)

	switch {
	case err != nil:
		return fmt.Errorf("%v: stringZ: %w", ErrWriterWrite, err)
	case bytesWritten != len(stringBytes):
		return fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, len(stringBytes), bytesWritten)
	}

	return nil
}

// WriteBytes writes byte string into underlying writer.
// Returns error if written bytes count mismatch specified byte string length or any underlying error if occurs.
func (w BinaryWriter) WriteBytes(data []byte) error {
	bytesWritten, err := w.writer.Write(data)

	switch {
	case err != nil:
		return fmt.Errorf("%v: %w", ErrWriterWrite, err)
	case bytesWritten != len(data):
		return fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, len(data), bytesWritten)
	}

	return nil
}

// WriteHex adds byte string defined by hex string into writer.
func (w BinaryWriter) WriteHex(hexString string) error {
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return fmt.Errorf("%v: %v: hex: %w", ErrWriter, ErrDecodeTo, err)
	}

	return w.WriteBytes(data)
}

// WriteObject writes object data into underlying writer.
// Returns written bytes count and possible error.
func (w BinaryWriter) WriteObject(data interface{}) error {
	switch binaryReaderFrom := data.(type) {
	case BinaryWriterTo:
		if _, err := binaryReaderFrom.BinaryWriteTo(w); err != nil {
			return fmt.Errorf("%v: %w", ErrWriterWrite, err)
		}

		return nil

	case io.WriterTo:
		if _, err := binaryReaderFrom.WriteTo(w); err != nil {
			return fmt.Errorf("%v: %w", ErrWriterWrite, err)
		}

		return nil

	default:
		return fmt.Errorf("%w: should implement io.WriterTo or binutils.BinaryWriterTo", ErrWriterWrite)
	}
}
