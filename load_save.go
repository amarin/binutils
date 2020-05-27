package binutils

import (
	"encoding"
	"os"
	"path/filepath"
)

// SaveBinary saves binary data of encoding.BinaryMarshaler implementing object into specified file.
func SaveBinary(filename string, dict encoding.BinaryMarshaler) error {
	absFileName, err := filepath.Abs(filename)
	if err != nil {
		return WrapError(err, "cant detect absolute path")
	}
	// prepare writer
	writer, err := os.Create(absFileName)
	if err != nil {
		return WrapError(err, "cant open file for writing")
	}
	// marshal data
	data, err := dict.MarshalBinary()
	if err != nil {
		return WrapError(err, "cant marshal data")
	}
	// write data
	written, err := writer.Write(data)
	if err != nil {
		return WrapError(err, "cant write data")
	}
	// check all data written
	if written != len(data) {
		return NewError("written only %d bytes of %d", written, len(data))
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
		return NewError("loaded %d bytes, cant unmarshal", bytesLoaded)
	} else if err := dict.UnmarshalFromBuffer(buffer); err != nil {
		return err
	} else if buffer.Len() != 0 {
		return NewError("consumed %d bytes, %d bytes rest", bytesLoaded-buffer.Len(), buffer.Len())
	}

	return nil
}
