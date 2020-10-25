package binutils

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// BinaryWriter implements binary writing for various data types into file writer.
type BinaryWriter struct {
	writer io.Writer
}

// CreateBinaryWriter creates file and wrap file writer into BinaryWriter.
// Target file will created.
func CreateBinaryWriter(filePath string) (*BinaryWriter, error) {
	if absFileName, err := filepath.Abs(filePath); err != nil {
		return nil, err
	} else if writer, err := os.Create(absFileName); err != nil {
		return nil, err
	} else {
		return &BinaryWriter{writer: writer}, nil
	}
}

// NewBinaryWriter wraps existing io.Writer into BinaryWriter.
func NewBinaryWriter(writer io.Writer) *BinaryWriter {
	return &BinaryWriter{writer: writer}
}

// Close closes underlying writer. Implements io.Closer.
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

// WriteUint16 writes uint8 value into writer as bytes.
func (w BinaryWriter) WriteUint8(data uint8) error {
	bytesWritten, err := w.writer.Write(Uint8bytes(data))

	switch {
	case err != nil:
		return err
	case bytesWritten != Uint8size:
		return fmt.Errorf("%w: expected %v written %v", io.ErrShortWrite, Uint8size, bytesWritten)
	}

	return nil
}

// WriteUint16 writes uint16 value into writer as bytes.
func (w BinaryWriter) WriteUint16(data uint16) error {
	bytesWritten, err := w.writer.Write(Uint16bytes(data))

	switch {
	case err != nil:
		return err
	case bytesWritten != Uint16size:
		return fmt.Errorf("%w: expected %v written %v", io.ErrShortWrite, Uint16size, bytesWritten)
	}

	return nil
}

// WriteUint32 writes uint16 value into writer as bytes.
func (w BinaryWriter) WriteUint32(data uint32) error {
	bytesWritten, err := w.writer.Write(Uint32bytes(data))

	switch {
	case err != nil:
		return err
	case bytesWritten != Uint32size:
		return fmt.Errorf("%w: expected %v written %v", io.ErrShortWrite, Uint32size, bytesWritten)
	}

	return nil
}

// WriteUint32 writes uint16 value into writer as bytes.
func (w BinaryWriter) WriteUint64(data uint64) error {
	bytesWritten, err := w.writer.Write(Uint64bytes(data))

	switch {
	case err != nil:
		return err
	case bytesWritten != Uint64size:
		return fmt.Errorf("%w: expected %v written %v", io.ErrShortWrite, Uint64size, bytesWritten)
	}

	return nil
}

// WriteInt uint value into writer as bytes.
func (w BinaryWriter) WriteUint(data uint) error {
	return w.WriteUint64(uint64(data))
}

// WriteInt8 writes int8 value into writer as byte.
func (w BinaryWriter) WriteInt8(data int8) error {
	bytesWritten, err := w.writer.Write(Int8bytes(data))

	switch {
	case err != nil:
		return err
	case bytesWritten != Int8size:
		return fmt.Errorf("%w: expected %v written %v", io.ErrShortWrite, Int8size, bytesWritten)
	}

	return nil
}

// WriteInt16 writes int16 value into writer as bytes.
func (w BinaryWriter) WriteInt16(data int16) error {
	bytesWritten, err := w.writer.Write(Int16bytes(data))

	switch {
	case err != nil:
		return err
	case bytesWritten != Int16size:
		return fmt.Errorf("%w: expected %v written %v", io.ErrShortWrite, Int16size, bytesWritten)
	}

	return nil
}

// WriteInt32writes int32 value into writer as bytes.
func (w BinaryWriter) WriteInt32(data int32) error {
	bytesWritten, err := w.writer.Write(Int32bytes(data))

	switch {
	case err != nil:
		return err
	case bytesWritten != Int32size:
		return fmt.Errorf("%w: expected %v written %v", io.ErrShortWrite, Int32size, bytesWritten)
	}

	return nil
}

// WriteInt64writes int64 value into writer as bytes.
func (w BinaryWriter) WriteInt64(data int64) error {
	bytesWritten, err := w.writer.Write(Int64bytes(data))

	switch {
	case err != nil:
		return err
	case bytesWritten != Int64size:
		return fmt.Errorf("%w: expected %v written %v", io.ErrShortWrite, Int64size, bytesWritten)
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
		return err
	case bytesWritten != len(stringBytes):
		return fmt.Errorf("%w: expected %v written %v", io.ErrShortWrite, len(stringBytes), bytesWritten)
	}

	return nil
}

// WriteStringZ writes byte string into underlying writer.
func (w BinaryWriter) WriteBytes(data []byte) error {
	bytesWritten, err := w.writer.Write(data)

	switch {
	case err != nil:
		return err
	case bytesWritten != len(data):
		return fmt.Errorf("%w: expected %v written %v", io.ErrShortWrite, len(data), bytesWritten)
	}

	return nil
}

// WriteHex adds byte string defined by hex string into writer.
func (w BinaryWriter) WriteHex(hexString string) error {
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return fmt.Errorf("%w: hex: %v", ErrDecodeTo, err)
	}

	return w.WriteBytes(data)
}

// WriteObject add encoding.BinaryMarshaler binary data into buffer.
// Returns written bytes count and possible error.
func (w BinaryWriter) WriteObject(data interface{}) error {
	switch binaryReaderFrom := data.(type) {
	case BinaryWriterTo:
		_, err := binaryReaderFrom.BinaryWriteTo(w)
		if err != nil {
			return err
		}

		return nil

	case io.WriterTo:
		_, err := binaryReaderFrom.WriteTo(w)
		if err != nil {
			return err
		}

		return nil

	default:
		return fmt.Errorf("%w: should implement io.WriterTo or binutils.BinaryWriterTo", ErrWrite)
	}
}
