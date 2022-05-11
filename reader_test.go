package binutils_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/amarin/binutils"
)

func TestBinaryReader_Read(t *testing.T) {
	buffer := bytes.NewBuffer(bytesValue)
	reader := NewBinaryReader(buffer)
	target := make([]byte, len(bytesValue))
	bytesTaken, err := reader.Read(target)
	require.NoError(t, err)
	require.Equal(t, bytesTaken, len(bytesValue))
	require.Equal(t, bytesTaken, reader.BytesTaken())
	reader.ResetBytesTaken()
	bytesTaken, err = reader.Read(target)
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadBytesCount(t *testing.T) {
	buffer := bytes.NewBuffer(bytesValue)
	reader := NewBinaryReader(buffer)
	for byteCount := 0; byteCount < len(bytesValue); byteCount++ {
		target, err := reader.ReadBytesCount(1)
		require.NoError(t, err)
		require.Equal(t, byteCount+1, reader.BytesTaken())
		require.Equal(t, bytesValue[byteCount], target[0])
	}

	reader.ResetBytesTaken()
	_, err := reader.ReadBytesCount(1)
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadUint8(t *testing.T) {
	for _, expected := range []uint8{0x11, 0x23, 0x34} {
		expectedHex := fmt.Sprintf("%x", expected)
		bufferBytes, err := hex.DecodeString(expectedHex)
		require.NoErrorf(t, err, "%v: %v", err, expectedHex)
		require.Equal(t, Uint8size, len(bufferBytes))
		buffer := bytes.NewBuffer(bufferBytes)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadUint8()
		require.NoError(t, err)
		require.Equal(t, Uint8size, reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadUint8()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadUint16(t *testing.T) {
	for _, expected := range []uint16{0x1112, 0x2334, 0x3445} {
		expectedHex := fmt.Sprintf("%x", expected)
		bufferBytes, err := hex.DecodeString(expectedHex)
		require.NoErrorf(t, err, "%v: %v", err, expectedHex)
		require.Equal(t, Uint16size, len(bufferBytes))
		buffer := bytes.NewBuffer(bufferBytes)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadUint16()
		require.NoError(t, err)
		require.Equal(t, Uint16size, reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadUint16()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadUint32(t *testing.T) {
	for _, expected := range []uint32{0x11121314, 0x23344556, 0x34455667} {
		expectedHex := fmt.Sprintf("%x", expected)
		bufferBytes, err := hex.DecodeString(expectedHex)
		require.NoErrorf(t, err, "%v: %v", err, expectedHex)
		require.Equal(t, Uint32size, len(bufferBytes))
		buffer := bytes.NewBuffer(bufferBytes)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadUint32()
		require.NoError(t, err)
		require.Equal(t, Uint32size, reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadUint32()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadUint64(t *testing.T) {
	for _, expected := range []uint64{0x1112131415161718, 0x2334455678900102, 0x3445566701020304} {
		expectedHex := fmt.Sprintf("%x", expected)
		bufferBytes, err := hex.DecodeString(expectedHex)
		require.NoErrorf(t, err, "%v: %v", err, expectedHex)
		require.Equal(t, Uint64size, len(bufferBytes))
		buffer := bytes.NewBuffer(bufferBytes)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadUint64()
		require.NoError(t, err)
		require.Equal(t, Uint64size, reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadUint64()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadUint(t *testing.T) {
	for _, expected := range []uint{0x1112131415161718, 0x2334455678900102, 0x3445566701020304} {
		expectedHex := fmt.Sprintf("%x", expected)
		bufferBytes, err := hex.DecodeString(expectedHex)
		require.NoErrorf(t, err, "%v: %v", err, expectedHex)
		require.Equal(t, Uint64size, len(bufferBytes))
		buffer := bytes.NewBuffer(bufferBytes)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadUint()
		require.NoError(t, err)
		require.Equal(t, Uint64size, reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadUint()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadInt8(t *testing.T) {
	for _, expected := range []int8{0x11, 0x23, 0x34} {
		expectedHex := fmt.Sprintf("%x", expected)
		bufferBytes, err := hex.DecodeString(expectedHex)
		require.NoErrorf(t, err, "%v: %v", err, expectedHex)
		require.Equal(t, Int8size, len(bufferBytes))
		buffer := bytes.NewBuffer(bufferBytes)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadInt8()
		require.NoError(t, err)
		require.Equal(t, Int8size, reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadInt8()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadInt16(t *testing.T) {
	for _, expected := range []int16{0x1112, 0x2334, 0x3445} {
		expectedHex := fmt.Sprintf("%x", expected)
		bufferBytes, err := hex.DecodeString(expectedHex)
		require.NoErrorf(t, err, "%v: %v", err, expectedHex)
		require.Equal(t, Int16size, len(bufferBytes))
		buffer := bytes.NewBuffer(bufferBytes)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadInt16()
		require.NoError(t, err)
		require.Equal(t, Int16size, reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadInt16()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadInt32(t *testing.T) {
	for _, expected := range []int32{0x11121314, 0x23344556, 0x34455667} {
		expectedHex := fmt.Sprintf("%x", expected)
		bufferBytes, err := hex.DecodeString(expectedHex)
		require.NoErrorf(t, err, "%v: %v", err, expectedHex)
		require.Equal(t, Int32size, len(bufferBytes))
		buffer := bytes.NewBuffer(bufferBytes)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadInt32()
		require.NoError(t, err)
		require.Equal(t, Int32size, reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadInt32()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadInt64(t *testing.T) {
	for _, expected := range []int64{0x1112131415161718, 0x2334455678900102, 0x3445566701020304} {
		expectedHex := fmt.Sprintf("%x", expected)
		bufferBytes, err := hex.DecodeString(expectedHex)
		require.NoErrorf(t, err, "%v: %v", err, expectedHex)
		require.Equal(t, Int64size, len(bufferBytes))
		buffer := bytes.NewBuffer(bufferBytes)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadInt64()
		require.NoError(t, err)
		require.Equal(t, Int64size, reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadInt64()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadInt(t *testing.T) {
	for _, expected := range []int{0x1112131415161718, 0x2334455678900102, 0x3445566701020304} {
		expectedHex := fmt.Sprintf("%x", expected)
		bufferBytes, err := hex.DecodeString(expectedHex)
		require.NoErrorf(t, err, "%v: %v", err, expectedHex)
		require.Equal(t, Int64size, len(bufferBytes))
		buffer := bytes.NewBuffer(bufferBytes)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadInt()
		require.NoError(t, err)
		require.Equal(t, Int64size, reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadInt()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadRune(t *testing.T) {
	for _, expected := range []rune{'Я', '±', 'ა', 'タ', 'W'} {
		expectedHex := fmt.Sprintf("%#08x", uint32(expected))[2:]
		bufferBytes, err := hex.DecodeString(expectedHex)
		require.NoErrorf(t, err, "%v: %v", err, expectedHex)
		require.Equal(t, RuneSize, len(bufferBytes))
		buffer := bytes.NewBuffer(bufferBytes)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadRune()
		require.NoError(t, err)
		require.Equal(t, RuneSize, reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadRune()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadBytes(t *testing.T) {
	for _, expected := range [][]byte{{0x10, 0x11, 0x12}, {0x20, 0x21, 0x31, 0x41}} {
		var stopByte uint8
		for testStop := 0; testStop < 255; testStop++ {
			intersected := false
			for idx := 0; idx < len(expected); idx++ {
				if uint8(testStop) == expected[idx] {
					intersected = true
					continue
				}

			}
			if !intersected {
				stopByte = uint8(testStop)
				break
			}
		}
		expected = append(expected, stopByte)
		buffer := bytes.NewBuffer(expected)
		reader := NewBinaryReader(buffer)
		target, err := reader.ReadBytes(stopByte)
		require.NoError(t, err)
		require.Equal(t, len(expected), reader.BytesTaken())
		require.Equal(t, expected, target)
	}

	reader := NewBinaryReader(bytes.NewBuffer(nil))
	reader.ResetBytesTaken()
	_, err := reader.ReadRune()
	require.Error(t, err) // error as buffer empty
	require.Equal(t, 0, reader.BytesTaken())
}

func TestBinaryReader_ReadStringZ(t *testing.T) {
	tests := []struct {
		name     string
		hexBytes string
		wantLine string
		wantErr  bool
	}{
		{"testStr", "7465737453747200", "testStr", false},
		{"testUnicode", "e6989f00", "星", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytesData, err := hex.DecodeString(tt.hexBytes)
			require.NoError(t, err)
			reader := NewBinaryReader(bytes.NewBuffer(bytesData))
			gotLine, err := reader.ReadStringZ()
			require.Equalf(t, tt.wantErr, err != nil, "want err %v, got %v", tt.wantErr, err)
			if err != nil {
				return
			}
			require.Equal(t, tt.wantLine, gotLine)
		})
	}
}

func TestBinaryReader_ReadHex(t *testing.T) {
	for _, tt := range []struct {
		name         string
		buffer       *bytes.Buffer
		wantBytes    int
		expectString string
		wantErr      bool
	}{
		{"empty_buffer_take_0_bytes",
			bytes.NewBuffer([]byte{}), 0,
			"", false},
		{"overflow_or_eof",
			bytes.NewBuffer([]byte{}), 8,
			"", true},
		{"single_byte",
			bytes.NewBuffer([]byte{1}), 1,
			"01", false},
		{"couple_of_bytes",
			bytes.NewBuffer([]byte{1, 1}), 2,
			"0101", false},
		{"three_of_bytes",
			bytes.NewBuffer([]byte{1, 1, 255}), 3,
			"0101ff", false},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			reader := NewBinaryReader(tt.buffer)
			got, err := reader.ReadHex(tt.wantBytes)
			require.Equalf(t, tt.wantErr, err != nil, "expected error %v, got %v", tt.wantErr, err)
			if err != nil {
				return
			}
			require.Equal(t, tt.expectString, got)
			require.Equal(t, tt.wantBytes, reader.BytesTaken())
			reader.ResetBytesTaken()
			require.Equal(t, 0, reader.BytesTaken())
		})
	}
}

func TestBinaryReader_ReadObject(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	writer := NewBinaryWriter(buffer)
	reader := NewBinaryReader(buffer)

	for _, tt := range readWriteTests {
		notAPointerAndNotABytes := !strings.Contains(tt.name, "ptr") && !strings.Contains(tt.name, "bytes")
		if notAPointerAndNotABytes || strings.Contains(tt.name, "nil") {
			continue
		}
		buffer.Reset()
		writer.ResetBytesWritten()
		reader.ResetBytesTaken()

		dataBytes, err := hex.DecodeString(tt.expectedHex)

		require.NoError(t, writer.WriteBytes(dataBytes))
		var expectedInstance interface{}
		switch tt.data.(type) {
		case *uint8:
			expectedInstance = new(uint8)
		case *uint16:
			expectedInstance = new(uint16)
		case *uint32:
			expectedInstance = new(uint32)
		case *uint64:
			expectedInstance = new(uint64)
		case *uint:
			expectedInstance = new(uint)
		case *int8:
			expectedInstance = new(int8)
		case *int16:
			expectedInstance = new(int16)
		case *int32:
			if strings.Contains(tt.name, "rune") {
				expectedInstance = new(rune)
			} else {
				expectedInstance = new(int32)
			}
		case *int64:
			expectedInstance = new(int64)
		case *int:
			expectedInstance = new(int)
		case *string:
			expectedInstance = new(string)
		case *[]byte:
			target := make([]byte, writer.BytesWritten())
			expectedInstance = &target
		case []byte:
			expectedInstance = make([]byte, writer.BytesWritten())
		}

		t.Run(tt.name, func(t *testing.T) {
			err = reader.ReadObject(expectedInstance)
			require.Equalf(t, tt.wantErr, err != nil, "want error %v, got %v", tt.wantErr, err)
			if err != nil {
				return
			}
			require.Equal(t, tt.expectedBytesWritten, reader.BytesTaken())

			require.EqualValues(t, tt.data, expectedInstance)
			require.NoError(t, writer.WriteBytes(dataBytes))
		})
	}
}
