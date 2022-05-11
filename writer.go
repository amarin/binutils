package binutils

import (
	"encoding"
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

func (w *BinaryWriter) increaseBytesWritten(addBytesCount int) {
	w.mu.Lock()
	w.bytesWritten += addBytesCount
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
func (w *BinaryWriter) Write(p []byte) (bytesWritten int, err error) {
	w.mu.Lock()
	bytesWritten, err = w.writer.Write(p)
	w.bytesWritten += bytesWritten
	w.mu.Unlock()

	switch {
	case err != nil:
		return bytesWritten, fmt.Errorf("%v: uint8: %w", ErrWriterWrite, err)
	case bytesWritten != len(p):
		return bytesWritten, fmt.Errorf(
			"%v: %w: expected %v written %v",
			ErrWriter, io.ErrShortWrite, len(p), bytesWritten)
	}

	return bytesWritten, err
}

// write wraps Write returning only error. Counts written bytes internally.
func (w *BinaryWriter) write(p []byte) (err error) {
	_, err = w.Write(p)

	return err
}

// WriteUint8 writes uint8 value into writer as bytes.
func (w *BinaryWriter) WriteUint8(data uint8) error {
	return w.write(Uint8bytes(data))
}

// WriteUint16 writes uint16 value into writer as bytes.
func (w *BinaryWriter) WriteUint16(data uint16) error {
	return w.write(Uint16bytes(data))
}

// WriteUint32 writes uint16 value into writer as bytes.
func (w *BinaryWriter) WriteUint32(data uint32) error {
	return w.write(Uint32bytes(data))
}

// WriteRune writes rune value into writer as uint32 bytes.
func (w *BinaryWriter) WriteRune(char rune) error {
	return w.write(RuneBytes(char))
}

// WriteUint64 writes uint64 value into writer as bytes.
func (w *BinaryWriter) WriteUint64(data uint64) error {
	return w.write(Uint64bytes(data))
}

// WriteUint uint value into writer as bytes.
func (w *BinaryWriter) WriteUint(data uint) (err error) {
	return w.write(Uint64bytes(uint64(data)))
}

// WriteInt8 writes int8 value into writer as byte.
func (w *BinaryWriter) WriteInt8(data int8) error {
	return w.write(Int8bytes(data))
}

// WriteInt16 writes int16 value into writer as bytes.
func (w *BinaryWriter) WriteInt16(data int16) error {
	return w.write(Int16bytes(data))
}

// WriteInt32 writes int32 value into writer as bytes.
func (w *BinaryWriter) WriteInt32(data int32) error {
	return w.write(Int32bytes(data))
}

// WriteInt64 writes int64 value into writer as bytes.
func (w *BinaryWriter) WriteInt64(data int64) error {
	return w.write(Int64bytes(data))
}

// WriteInt int value into writer as bytes.
func (w *BinaryWriter) WriteInt(data int) error {
	return w.write(Int64bytes(int64(data)))
}

// WriteStringZ writes string bytes into underlying writer as Zero-terminated string.
func (w *BinaryWriter) WriteStringZ(data string) error {
	return w.write(StringBytes(data))
}

// WriteBytes writes byte string into underlying writer.
// Returns error if written bytes count mismatch specified byte string length or any underlying error if occurs.
func (w *BinaryWriter) WriteBytes(data []byte) error {
	return w.write(data)
}

// WriteHex adds byte string defined by hex string into writer.
func (w *BinaryWriter) WriteHex(hexString string) error {
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return fmt.Errorf("%v: %v: hex: %w", ErrWriter, ErrDecodeTo, err)
	}

	return w.write(data)
}

// WriteObject writes object data into underlying writer.
// User specified data types data must be one of io.WriterTo, BinaryWriterTo, BinaryUint8, BinaryUint16, BinaryUint32, BinaryUint64,
// BinaryInt8, BinaryInt16, BinaryInt32, BinaryInt64 or BinaryRune interface implementation.
// Basic Int[8-64], Uint[8-64] or pointers to it are simply generates bigEndian bytes.
//
// If multiple interfaces implemented first of described order will be used.
// Use required method directly to fully determined behaviour.
// Returns error if caused internally. To get written bytes counter use BytesWritten result.
func (w *BinaryWriter) WriteObject(data interface{}) (err error) {
	n := int64(0)

	switch typedValue := data.(type) {
	case io.WriterTo:
		n, err = typedValue.WriteTo(w)
		w.increaseBytesWritten(int(n))
	case encoding.BinaryMarshaler:
		var binaryData []byte
		if binaryData, err = typedValue.MarshalBinary(); err != nil {
			return fmt.Errorf("%w: marshal: %v", Error, err)
		}

		return w.write(binaryData)
	case BinaryWriterTo:
		return typedValue.BinaryWriteTo(w)
	case uint8:
		return w.WriteUint8(typedValue)
	case *uint8:
		if typedValue == nil {
			return fmt.Errorf("%w: uint8", ErrNilPointer)
		}
		return w.WriteUint8(*typedValue)
	case uint16:
		return w.WriteUint16(typedValue)
	case *uint16:
		if typedValue == nil {
			return fmt.Errorf("%w: uint16", ErrNilPointer)
		}
		return w.WriteUint16(*typedValue)
	case uint32:
		return w.WriteUint32(typedValue)
	case *uint32:
		if typedValue == nil {
			return fmt.Errorf("%w: uint32", ErrNilPointer)
		}
		return w.WriteUint32(*typedValue)
	case uint64:
		return w.WriteUint64(typedValue)
	case *uint64:
		if typedValue == nil {
			return fmt.Errorf("%w: uint64", ErrNilPointer)
		}
		return w.WriteUint64(*typedValue)
	case uint:
		return w.WriteUint(typedValue)
	case *uint:
		if typedValue == nil {
			return fmt.Errorf("%w: uint", ErrNilPointer)
		}
		return w.WriteUint(*typedValue)
	case int8:
		return w.WriteInt8(typedValue)
	case *int8:
		if typedValue == nil {
			return fmt.Errorf("%w: int8", ErrNilPointer)
		}
		return w.WriteInt8(*typedValue)
	case int16:
		return w.WriteInt16(typedValue)
	case *int16:
		if typedValue == nil {
			return fmt.Errorf("%w: int16", ErrNilPointer)
		}
		return w.WriteInt16(*typedValue)
	case int32:
		return w.WriteInt32(typedValue)
	case *int32:
		if typedValue == nil {
			return fmt.Errorf("%w: int32", ErrNilPointer)
		}
		return w.WriteInt32(*typedValue)
	case int64:
		return w.WriteInt64(typedValue)
	case *int64:
		if typedValue == nil {
			return fmt.Errorf("%w: int64", ErrNilPointer)
		}
		return w.WriteInt64(*typedValue)
	case int:
		return w.WriteInt(typedValue)
	case *int:
		if typedValue == nil {
			return fmt.Errorf("%w: int", ErrNilPointer)
		}
		return w.WriteInt(*typedValue)
	case string:
		return w.WriteStringZ(typedValue)
	case *string:
		if typedValue == nil {
			return fmt.Errorf("%w: string", ErrNilPointer)
		}
		return w.WriteStringZ(*typedValue)
	case []uint8: // cover []byte
		return w.WriteBytes(typedValue)
	case *[]uint8:
		if typedValue == nil {
			return fmt.Errorf("%w: []byte", ErrNilPointer)
		}
		return w.WriteBytes(*typedValue)
	default:
		return fmt.Errorf(
			"%w: %T should implement io.WriterTo, binutils.BinaryWriterTo or any Binary<Type> interface",
			ErrWriterWrite, typedValue,
		)
	}

	if err != nil {
		return fmt.Errorf("%v: %w", ErrWriterWrite, err)
	}

	return nil
}
