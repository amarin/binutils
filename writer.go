package binutils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var (
	// ErrWriter identifies any writer errors.
	ErrWriter = errors.New("writer")

	// ErrWriterWrite identifies writer failed during write.
	ErrWriterWrite = fmt.Errorf("%w: write", ErrWriter)
)

// BinaryWriter implements binary writing for various data types into file writer.
type BinaryWriter struct {
	mu           *sync.Mutex // write mutex protects underlying fields
	writer       io.Writer   // underlying io.Writer
	bytesWritten int         // written bytes counter
}

// NewBinaryWriter wraps existing io.Writer instance into BinaryWriter.
func NewBinaryWriter(writer io.Writer) *BinaryWriter {
	return &BinaryWriter{writer: writer, bytesWritten: 0, mu: new(sync.Mutex)}
}

// BytesWritten returns written bytes counter value.
// Note counter can be reset to 0 using ResetBytesWritten.
func (w *BinaryWriter) BytesWritten() (res int) {
	w.mu.Lock()
	res = w.bytesWritten
	w.mu.Unlock()

	return res
}

// ResetBytesWritten sets written bytes counter to zero.
func (w *BinaryWriter) ResetBytesWritten() {
	w.mu.Lock()
	w.bytesWritten = 0
	w.mu.Unlock()
}

// CreateFile creates file and wrap file writer into BinaryWriter.
// Target file will be created.
func CreateFile(filePath string) (*BinaryWriter, error) {
	w := NewBinaryWriter(nil)

	if absFileName, err := filepath.Abs(filePath); err != nil {
		return nil, fmt.Errorf("%v: resolve path: %w", ErrWriter, err)
	} else if w.writer, err = os.Create(absFileName); err != nil {
		return nil, fmt.Errorf("%v: create file: %w", ErrWriter, err)
	} else {
		return w, nil
	}
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
// Note it extends internal written bytes counter to written bytes value.
// Implements io.Writer.
func (w *BinaryWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	n, err = w.writer.Write(p)
	w.bytesWritten += n
	w.mu.Unlock()

	return n, err
}

// WriteUint8 writes uint8 value into writer as bytes.
func (w *BinaryWriter) WriteUint8(data uint8) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	bytesWritten, err := w.writer.Write(Uint8bytes(data))
	w.bytesWritten += bytesWritten

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
func (w *BinaryWriter) WriteUint16(data uint16) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	bytesWritten, err := w.writer.Write(Uint16bytes(data))
	w.bytesWritten += bytesWritten

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
func (w *BinaryWriter) WriteUint32(data uint32) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	bytesWritten, err := w.writer.Write(Uint32bytes(data))
	w.bytesWritten += bytesWritten

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

// WriteRune writes rune value into writer as uint32 bytes.
func (w *BinaryWriter) WriteRune(char rune) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	bytesWritten, err := w.writer.Write(RuneBytes(char))
	w.bytesWritten += bytesWritten

	switch {
	case err != nil:
		return fmt.Errorf("%v: rune: %w", ErrWriterWrite, err)
	case bytesWritten != RuneSize:
		return fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, RuneSize, bytesWritten)
	}

	return nil
}

// WriteUint64 writes uint64 value into writer as bytes.
func (w *BinaryWriter) WriteUint64(data uint64) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	bytesWritten, err := w.writer.Write(Uint64bytes(data))
	w.bytesWritten += bytesWritten

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
func (w *BinaryWriter) WriteUint(data uint) (err error) {
	return w.WriteUint64(uint64(data))
}

// WriteInt8 writes int8 value into writer as byte.
func (w *BinaryWriter) WriteInt8(data int8) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	bytesWritten, err := w.writer.Write(Int8bytes(data))
	w.bytesWritten += bytesWritten

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
func (w *BinaryWriter) WriteInt16(data int16) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	bytesWritten, err := w.writer.Write(Int16bytes(data))
	w.bytesWritten += bytesWritten

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
func (w *BinaryWriter) WriteInt32(data int32) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	bytesWritten, err := w.writer.Write(Int32bytes(data))
	w.bytesWritten += bytesWritten

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
func (w *BinaryWriter) WriteInt64(data int64) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	bytesWritten, err := w.writer.Write(Int64bytes(data))
	w.bytesWritten += bytesWritten

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
func (w *BinaryWriter) WriteInt(data int) error {
	return w.WriteInt64(int64(data))
}

// WriteStringZ writes string bytes into underlying writer as Zero-terminated string.
func (w *BinaryWriter) WriteStringZ(data string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	stringBytes := StringBytes(data)
	bytesWritten, err := w.writer.Write(stringBytes)
	w.bytesWritten += bytesWritten

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
func (w *BinaryWriter) WriteBytes(data []byte) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	bytesWritten, err := w.writer.Write(data)
	w.bytesWritten += bytesWritten

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
func (w *BinaryWriter) WriteHex(hexString string) error {
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return fmt.Errorf("%v: %v: hex: %w", ErrWriter, ErrDecodeTo, err)
	}

	return w.WriteBytes(data)
}

// WriteObject writes object data into underlying writer.
// Specified data could be one of io.WriterTo, BinaryWriterTo, BinaryUint8, BinaryUint16, BinaryUint32, BinaryUint64,
// BinaryInt8, BinaryInt16, BinaryInt32, BinaryInt64 or BinaryRune interface implementation.
// If multiple interfaces implemented first of described order will be used.
// Use required method directly to fully determined behaviour.
// Returns error if caused internally. To get written bytes counter use BytesWritten result.
func (w *BinaryWriter) WriteObject(data interface{}) (err error) {
	n := int64(0)

	switch binaryObject := data.(type) {
	case io.WriterTo:
		n, err = binaryObject.WriteTo(w)
	case BinaryWriterTo:
		err = binaryObject.BinaryWriteTo(w)
	case BinaryUint8:
		err = w.WriteUint8(binaryObject.Uint8())
	case BinaryUint16:
		err = w.WriteUint16(binaryObject.Uint16())
	case BinaryUint32:
		err = w.WriteUint32(binaryObject.Uint32())
	case BinaryUint64:
		err = w.WriteUint64(binaryObject.Uint64())
	case BinaryInt8:
		err = w.WriteInt8(binaryObject.Int8())
	case BinaryInt16:
		err = w.WriteInt16(binaryObject.Int16())
	case BinaryInt32:
		err = w.WriteInt32(binaryObject.Int32())
	case BinaryInt64:
		err = w.WriteInt64(binaryObject.Int64())
	case BinaryRune:
		err = w.WriteRune(binaryObject.Rune())
	case BinaryString:
		err = w.WriteStringZ(binaryObject.BinaryString())
	default:
		return fmt.Errorf(
			"%w: should implement io.WriterTo, binutils.BinaryWriterTo or any Binary<Type> interface",
			ErrWriterWrite,
		)
	}

	if n != 0 {
		w.mu.Lock()
		w.bytesWritten += int(n)
		w.mu.Unlock()
	}

	if err != nil {
		return fmt.Errorf("%v: %w", ErrWriterWrite, err)
	}

	return nil
}
