package binutils

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// BinaryReader implements binary writing for various data types into file writer.
type BinaryReader struct {
	mu         *sync.Mutex // read mutex protects underlying fields
	source     io.Reader
	bytesTaken int
}

// OpenFile opens specified file path and returns BinaryReader wrapping it.
// Target filePath must be present and readable before opening.
func OpenFile(filePath string) (*BinaryReader, error) {
	if absFileName, err := filepath.Abs(filePath); err != nil {
		return nil, err
	} else if source, err := os.Open(absFileName); err != nil {
		return nil, err
	} else {
		return NewBinaryReader(source), nil
	}
}

// NewBinaryReader wraps existing io.Reader into BinaryReader.
func NewBinaryReader(source io.Reader) *BinaryReader {
	return &BinaryReader{source: source, mu: new(sync.Mutex), bytesTaken: 0}
}

// ResetBytesTaken zeroes internal bytes taken counter.
func (r *BinaryReader) ResetBytesTaken() {
	r.mu.Lock()
	r.bytesTaken = 0
	r.mu.Unlock()
}

// BytesTaken returns bytesTaken counter since last ResetBytesTaken invoke or created if never ResetBytesTaken called.
func (r *BinaryReader) BytesTaken() (bytesTaken int) {
	r.mu.Lock()
	bytesTaken = r.bytesTaken
	r.mu.Unlock()

	return bytesTaken
}

// increaseBytesTaken adds specified addBytesTaken value to internal bytes taken counter.
func (r *BinaryReader) increaseBytesTaken(addBytesTaken int) {
	r.mu.Lock()
	r.bytesTaken += addBytesTaken
	r.mu.Unlock()
}

// Close closes underlying reader. Implements io.Closer.
// Returns error if underlying reader not  implements io.Closer.
func (r *BinaryReader) Close() error {
	closer, ok := r.source.(io.Closer)
	if !ok {
		return fmt.Errorf("%w: %T is not io.Closer", ErrClose, r.source)
	}

	return closer.Close()
}

// Read reads up to len(p) bytes into p. It returns the number of bytes taken (0 <= n <= len(p))
// and any error encountered, just calling underlying io.Reader Read method.
// Note it extends internal taken bytes counter to taken bytes value.
// Implements io.Reader itself.
func (r *BinaryReader) Read(p []byte) (n int, err error) {
	r.mu.Lock()
	n, err = r.source.Read(p)
	r.bytesTaken += n
	r.mu.Unlock()

	return n, err
}

// read wraps Read returning only error.
func (r *BinaryReader) read(p []byte) (err error) {
	_, err = r.Read(p)

	return err
}

// ReadBytesCount reads exactly specified amount of bytes.
// Returns read bytes or error if insufficient bytes count ready to read or any underlying reader error encountered.
func (r *BinaryReader) ReadBytesCount(amount int) (buffer []byte, err error) {
	buffer = make([]byte, amount)
	if err = r.read(buffer); err != nil { // read required bytes amount counting taken bytes internally
		return buffer, err
	}

	return buffer, nil
}

// ReadUint8 reads uint8 value from underlying reader.
// Returns uint8 value and any error encountered.
func (r *BinaryReader) ReadUint8() (res uint8, err error) {
	byteBuffer := AllocateBytes(Uint8size)

	if err = r.read(byteBuffer); err != nil { // read required bytes amount counting taken bytes internally
		return 0, err
	}

	return Uint8(byteBuffer)
}

// ReadUint16 reads uint16 value from underlying reader.
// Returns uint16 value and any error encountered.
func (r *BinaryReader) ReadUint16() (res uint16, err error) {
	byteBuffer := AllocateBytes(Uint16size)
	if err = r.read(byteBuffer); err != nil { // read required bytes amount counting taken bytes internally
		return 0, err
	}

	return Uint16(byteBuffer)
}

// ReadUint32 reads uint32 value from underlying reader.
// Returns uint32 value and any error encountered.
func (r *BinaryReader) ReadUint32() (res uint32, err error) {
	byteBuffer := AllocateBytes(Uint32size)

	if err = r.read(byteBuffer); err != nil { // read required bytes amount counting taken bytes internally
		return 0, err
	}

	return Uint32(byteBuffer)
}

// ReadUint64 reads uint64 value from underlying reader.
// Returns uint64 value and any error encountered.
func (r *BinaryReader) ReadUint64() (res uint64, err error) {
	byteBuffer := AllocateBytes(Uint64size)

	if err = r.read(byteBuffer); err != nil { // read required bytes amount counting taken bytes internally
		return 0, err
	}
	return Uint64(byteBuffer)

}

// ReadUint reads uint value from underlying reader.
// Returns uint value and any error encountered.
func (r *BinaryReader) ReadUint() (uint, error) {
	uint64result, err := r.ReadUint64()
	return uint(uint64result), err
}

// ReadInt8 reads int8 value from underlying reader.
// Returns int8 value and any error encountered.
func (r *BinaryReader) ReadInt8() (res int8, err error) {
	byteBuffer := AllocateBytes(Int8size)
	if err = r.read(byteBuffer); err != nil { // read required bytes amount counting taken bytes internally
		return 0, err
	}

	return Int8(byteBuffer)
}

// ReadInt16 reads int16 value from underlying reader.
// Returns int16 value and any error encountered.
func (r *BinaryReader) ReadInt16() (res int16, err error) {
	byteBuffer := AllocateBytes(Int16size)
	if err = r.read(byteBuffer); err != nil { // read required bytes amount counting taken bytes internally
		return 0, err
	}

	return Int16(byteBuffer)
}

// ReadInt32 reads int32 value from underlying reader.
// Returns int32 value and any error encountered.
func (r *BinaryReader) ReadInt32() (res int32, err error) {
	byteBuffer := AllocateBytes(Int32size)
	if err = r.read(byteBuffer); err != nil { // read required bytes amount counting taken bytes internally
		return 0, err
	}

	return Int32(byteBuffer)
}

// ReadInt64 reads int64 value from underlying reader.
// Returns int64 value and any error encountered.
func (r *BinaryReader) ReadInt64() (res int64, err error) {
	byteBuffer := AllocateBytes(Int64size)
	if err = r.read(byteBuffer); err != nil { // read required bytes amount counting taken bytes internally
		return 0, err
	}

	return Int64(byteBuffer)
}

// ReadInt reads int value from underlying reader.
// Returns int value and any error encountered.
func (r *BinaryReader) ReadInt() (int, error) {
	int64result, err := r.ReadInt64()
	return int(int64result), err
}

// ReadRune reads rune value from underlying io.Reader.
// Returns rune value and any error encountered.
func (r *BinaryReader) ReadRune() (res rune, err error) {
	byteBuffer := AllocateBytes(RuneSize)
	if err = r.read(byteBuffer); err != nil { // read required bytes amount counting taken bytes internally
		return 0, err
	}

	return Rune(byteBuffer)
}

// ReadBytes reads bytes sequence until the first occurrence of stop byte in the input.
// Returns a bytes slice containing the data up to and including the delimiter.
// If ReadBytes encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
func (r *BinaryReader) ReadBytes(stop byte) (dataTaken []byte, err error) {
	alreadyImplemented, ok := r.source.(untilStopByteReader)
	if ok {
		dataTaken, err = alreadyImplemented.ReadBytes(stop)
		r.increaseBytesTaken(len(dataTaken)) // increase counter to taken bytes len

		return dataTaken, err
	}
	// underlying reader does not implement read bytes until stop,
	// so read byte-by-byte and compare next ones until stop byte found or any read error happened.
	var (
		currentByte uint8
		buffer      = make([]byte, 0)
	)

	for {
		if currentByte, err = r.ReadUint8(); err != nil { // counter increased internally in ReadUint8
			return buffer, err
		}

		buffer = append(buffer, currentByte)

		if currentByte == stop {
			return buffer, nil
		}
	}
}

// ReadStringZ reads zero-terminated string from underlying reader.
func (r *BinaryReader) ReadStringZ() (line string, err error) {
	var dataTaken []byte

	if dataTaken, err = r.ReadBytes(0); err != nil {
		return "", fmt.Errorf("%w: read: %v", ErrRequired0T, err)
	}

	return string(dataTaken[:len(dataTaken)-1]), nil
}

// ReadHex reads exactly specified amount of bytes and return hex representation string for received bytes.
// Returns underlying reader errors encountered.
func (r *BinaryReader) ReadHex(amount int) (hexString string, err error) {
	var dataTaken []byte

	if dataTaken, err = r.ReadBytesCount(amount); err != nil {
		return "", err
	}

	return hex.EncodeToString(dataTaken), nil
}

// ReadObject reads object data from underlying io.Reader.
// Returns written bytes count and possible error.
func (r *BinaryReader) ReadObject(target interface{}) error {
	switch tgtType := target.(type) {
	case *uint8:
		receivedValue, err := r.ReadUint8()
		*tgtType = receivedValue
		return err
	case *uint16:
		receivedValue, err := r.ReadUint16()
		*tgtType = receivedValue
		return err
	case *uint32:
		receivedValue, err := r.ReadUint32()
		*tgtType = receivedValue
		return err
	case *uint64:
		receivedValue, err := r.ReadUint64()
		*tgtType = receivedValue
		return err
	case *uint:
		receivedValue, err := r.ReadUint()
		*tgtType = receivedValue
		return err
	case *int8:
		receivedValue, err := r.ReadInt8()
		*tgtType = receivedValue
		return err
	case *int16:
		receivedValue, err := r.ReadInt16()
		*tgtType = receivedValue
		return err
	case *int32: // covers both int32 and rune read
		receivedValue, err := r.ReadInt32()
		*tgtType = receivedValue
		return err
	case *int64:
		receivedValue, err := r.ReadInt64()
		*tgtType = receivedValue
		return err
	case *int:
		receivedValue, err := r.ReadInt()
		*tgtType = receivedValue
		return err
	case []uint8: // cover []byte
		receivedValue, err := r.ReadBytesCount(len(tgtType))
		copy(tgtType, receivedValue)
		return err
	case *[]uint8:
		if tgtType == nil {
			return fmt.Errorf("%w: []byte", ErrNilPointer)
		}
		receivedValue, err := r.ReadBytesCount(len(*tgtType))
		copy(*tgtType, receivedValue)
		return err
	case *string:
		receivedValue, err := r.ReadStringZ()
		*tgtType = receivedValue
		return err
	case BinaryReaderFrom:
		bytesTaken, err := tgtType.BinaryReadFrom(r)
		r.increaseBytesTaken(int(bytesTaken))
		if err != nil {
			return err
		}

		return nil
	case io.ReaderFrom:
		bytesTaken, err := tgtType.ReadFrom(r)
		r.increaseBytesTaken(int(bytesTaken))

		if err != nil {
			return err
		}

		return nil

	default:
		return fmt.Errorf("%w: %T should implement io.ReaderFrom or binutils.BinaryReaderFrom", ErrRead, tgtType)
	}
}
