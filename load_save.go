package binutils

import (
	"encoding"
	"fmt"
	"os"
	"path/filepath"
)

// SaveBinary saves binary data of encoding.BinaryMarshaler implementing object into specified file.
func SaveBinary(filename string, dict encoding.BinaryMarshaler) error {
	absFileName, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	// prepare writer
	writer, err := os.Create(absFileName)
	if err != nil {
		return err
	}
	// marshal data
	data, err := dict.MarshalBinary()
	if err != nil {
		return err
	}
	// write data
	written, err := writer.Write(data)
	if err != nil {
		return err
	}
	// check all data written
	if written != len(data) {
		return fmt.Errorf("%w: written %d of %d", ErrMissedData, written, len(data))
	}
	// return ok
	return nil
}

// LoadBinary adds binary data from specified file into target BufferUnmarshaler implementing object
// Returns error if any file path resolution problem or file data empty or some binary data was not decoded.
func LoadBinary(filename string, dict BufferUnmarshaler) error {
	buffer := NewBuffer([]byte{})
	if bytesLoaded, err := buffer.LoadFromFilePath(filename); err != nil {
		return err
	} else if bytesLoaded == 0 {
		return fmt.Errorf("%w: loaded 0", ErrMissedData)
	} else if err := dict.UnmarshalFromBuffer(buffer); err != nil {
		return err
	} else if buffer.Len() != 0 {
		return fmt.Errorf("%w: %d consumed, %d rest", ErrExtraData, bytesLoaded-buffer.Len(), buffer.Len())
	}

	return nil
}
