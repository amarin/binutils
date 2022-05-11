package binutils_test

import (
	"bytes"
	"encoding/hex"
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/amarin/binutils"
)

var (
	nilPtr = new(struct {
		uint8  *uint8
		uint16 *uint16
		uint32 *uint32
		uint64 *uint64
		uint   *uint
		int8   *int8
		int16  *int16
		int32  *int32
		int64  *int64
		int    *int
		string *string
		rune   *rune
		bytes  *[]byte
	})
	uint8Value   = uint8(math.MaxInt8)
	uint16Value  = uint16(math.MaxInt16)
	uint32Value  = uint32(math.MaxInt32)
	uint64Value  = uint64(math.MaxInt64)
	uintValue    = uint(uint64Value)
	int8Value    = int8(math.MinInt8)
	int16Value   = int16(math.MinInt16)
	int32Value   = int32(math.MinInt32)
	int64Value   = int64(math.MinInt64)
	intValue     = int(int64Value)
	stringValue  = "testStr"
	unicodeValue = "星"
	runeValue    = '星'
	bytesValue   = []byte{0x00, 0x01, 0x02, 0x03, 0x04}

	readWriteTests = []struct {
		name                 string
		data                 interface{}
		expectedBytesWritten int
		expectedHex          string
		wantErr              bool
	}{
		{"uint8", uint8Value, 1, "7f", false},
		{"uint8_ptr", &uint8Value, 1, "7f", false},
		{"uint8_nil_ptr_err", nilPtr.uint8, 0, "7f", true},
		{"uint16", uint16Value, 2, "7fff", false},
		{"uint16_ptr", &uint16Value, 2, "7fff", false},
		{"uint16_nil_ptr_err", nilPtr.uint16, 0, "7fff", true},
		{"uint32", uint32Value, 4, "7fffffff", false},
		{"uint32_ptr", &uint32Value, 4, "7fffffff", false},
		{"uint32_nil_ptr_err", nilPtr.uint32, 0, "7fffffff", true},
		{"uint64", uint64Value, 8, "7fffffffffffffff", false},
		{"uint64_ptr", &uint64Value, 8, "7fffffffffffffff", false},
		{"uint64_nil_ptr_err", nilPtr.uint64, 0, "7fffffffffffffff", true},
		{"uint", uintValue, 8, "7fffffffffffffff", false},
		{"uint_ptr", &uintValue, 8, "7fffffffffffffff", false},
		{"uint_nil_ptr_err", nilPtr.uint, 0, "7fffffffffffffff", true},
		{"int8", int8Value, 1, "80", false},
		{"int8_ptr", &int8Value, 1, "80", false},
		{"int8_nil_ptr_err", nilPtr.int8, 0, "80", true},
		{"int16", int16Value, 2, "8000", false},
		{"int16_ptr", &int16Value, 2, "8000", false},
		{"int16_nil_ptr_err", nilPtr.int16, 0, "8000", true},
		{"int32", int32Value, 4, "80000000", false},
		{"int32_ptr", &int32Value, 4, "80000000", false},
		{"int32_nil_ptr_err", nilPtr.int32, 0, "80000000", true},
		{"int64", int64Value, 8, "8000000000000000", false},
		{"int64_ptr", &int64Value, 8, "8000000000000000", false},
		{"int64_nil_ptr_err", nilPtr.int64, 0, "8000000000000000", true},
		{"int", intValue, 8, "8000000000000000", false},
		{"int_ptr", &intValue, 8, "8000000000000000", false},
		{"int_nil_ptr_err", nilPtr.int, 0, "8000000000000000", true},
		{"string_ascii", stringValue, 8, "7465737453747200", false},
		{"string_ascii_ptr", &stringValue, 8, "7465737453747200", false},
		{"string_unicode", unicodeValue, 4, "e6989f00", false},
		{"string_unicode_ptr", &unicodeValue, 4, "e6989f00", false},
		{"string_nil_ptr_err", nilPtr.string, 0, "e6989f00", true},
		{"rune", runeValue, 4, "0000661f", false},
		{"rune_ptr", &runeValue, 4, "0000661f", false},
		{"rune_nil_ptr_err", nilPtr.rune, 0, "", true},
		{"bytes", bytesValue, 5, "0001020304", false},
		{"bytes_ptr", &bytesValue, 5, "0001020304", false},
		{"bytes_nil_ptr_err", nilPtr.bytes, 0, "0001020304", true},
	}
)

func TestBinaryWriter_Write(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		args    args
		wantN   int
		wantErr bool
	}{
		{"empty_slice_writes_0", args{[]byte{}}, 0, false},
		{"write_one_byte", args{[]byte{0x01}}, 1, false},
		{"write_some_byte", args{[]byte{0x01, 0xff, 0x13, 0xee}}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := bytes.NewBuffer(make([]byte, 0))
			writer := binutils.NewBinaryWriter(collector)
			gotN, err := writer.Write(tt.args.p)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.wantN, gotN)
			require.Equal(t, tt.wantN, writer.BytesWritten())
			if err != nil {
				return
			}
			writtenData := collector.Bytes()
			require.Equal(t, len(tt.args.p), len(writtenData))
			require.Equal(t, tt.args.p, writtenData)
			writer.ResetBytesWritten()
			require.Equal(t, 0, writer.BytesWritten())
		})
	}
}

func TestBinaryWriter_WriteUint8(t *testing.T) {
	type args struct {
		data uint8
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"write_000", args{0}, false},
		{"write_100", args{100}, false},
		{"write_255", args{255}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := bytes.NewBuffer(nil)
			writer := binutils.NewBinaryWriter(collector)
			err := writer.WriteUint8(tt.args.data)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, binutils.Uint8size, writer.BytesWritten())
			if err != nil {
				return
			}
			writtenData := collector.Bytes()
			require.Equal(t, binutils.Uint8size, len(writtenData))
			require.Equal(t, tt.args.data, writtenData[0])
			writer.ResetBytesWritten()
			require.Equal(t, 0, writer.BytesWritten())
		})
	}
}

func TestBinaryWriter_WriteUint16(t *testing.T) {
	type args struct {
		data uint16
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"write_000", args{0}, false},
		{"write_100", args{100}, false},
		{"write_max_uint16", args{math.MaxUint16}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := bytes.NewBuffer(nil)
			writer := binutils.NewBinaryWriter(collector)
			err := writer.WriteUint16(tt.args.data)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, binutils.Uint16size, writer.BytesWritten())
			if err != nil {
				return
			}
			writtenData := collector.Bytes()
			require.Equal(t, binutils.Uint16size, len(writtenData))
			var (
				value uint16
			)
			value, err = binutils.Uint16(writtenData[0:binutils.Uint16size])
			require.NoError(t, err)
			require.Equal(t, tt.args.data, value)
			writer.ResetBytesWritten()
			require.Equal(t, 0, writer.BytesWritten())
		})
	}
}

func TestBinaryWriter_WriteUint32(t *testing.T) {
	type args struct {
		data uint32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"write_000", args{0}, false},
		{"write_100", args{100}, false},
		{"write_max_uint32", args{math.MaxUint32}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := bytes.NewBuffer(nil)
			writer := binutils.NewBinaryWriter(collector)
			err := writer.WriteUint32(tt.args.data)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, binutils.Uint32size, writer.BytesWritten())
			if err != nil {
				return
			}
			writtenData := collector.Bytes()
			require.Equal(t, binutils.Uint32size, len(writtenData))
			var (
				value uint32
			)
			value, err = binutils.Uint32(writtenData[0:binutils.Uint32size])
			require.NoError(t, err)
			require.Equal(t, tt.args.data, value)
			writer.ResetBytesWritten()
			require.Equal(t, 0, writer.BytesWritten())
		})
	}
}

func TestBinaryWriter_WriteUint64(t *testing.T) {
	type args struct {
		data uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"write_000", args{0}, false},
		{"write_100", args{100}, false},
		{"write_max_uint64", args{math.MaxUint64}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := bytes.NewBuffer(nil)
			writer := binutils.NewBinaryWriter(collector)
			err := writer.WriteUint64(tt.args.data)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, binutils.Uint64size, writer.BytesWritten())
			if err != nil {
				return
			}
			writtenData := collector.Bytes()
			require.Equal(t, binutils.Uint64size, len(writtenData))
			var (
				value uint64
			)
			value, err = binutils.Uint64(writtenData[0:binutils.Uint64size])
			require.NoError(t, err)
			require.Equal(t, tt.args.data, value)
			writer.ResetBytesWritten()
			require.Equal(t, 0, writer.BytesWritten())
		})
	}
}

func TestBinaryWriter_WriteInt8(t *testing.T) {
	type args struct {
		data int8
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"write_000", args{0}, false},
		{"write_100", args{100}, false},
		{"write_min_int8", args{math.MinInt8}, false},
		{"write_max_int8", args{math.MaxInt8}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := bytes.NewBuffer(nil)
			writer := binutils.NewBinaryWriter(collector)
			err := writer.WriteInt8(tt.args.data)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, binutils.Int8size, writer.BytesWritten())
			if err != nil {
				return
			}
			writtenData := collector.Bytes()
			require.Equal(t, binutils.Int8size, len(writtenData))
			var (
				value int8
			)
			value, err = binutils.Int8(writtenData[0:binutils.Int8size])
			require.Equal(t, tt.args.data, value)
			writer.ResetBytesWritten()
			require.Equal(t, 0, writer.BytesWritten())
		})
	}
}

func TestBinaryWriter_WriteInt16(t *testing.T) {
	type args struct {
		data int16
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"write_000", args{0}, false},
		{"write_100", args{100}, false},
		{"write_min_int16", args{math.MinInt16}, false},
		{"write_max_int16", args{math.MaxInt16}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := bytes.NewBuffer(nil)
			writer := binutils.NewBinaryWriter(collector)
			err := writer.WriteInt16(tt.args.data)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, binutils.Int16size, writer.BytesWritten())
			if err != nil {
				return
			}
			writtenData := collector.Bytes()
			require.Equal(t, binutils.Int16size, len(writtenData))
			var (
				value int16
			)
			value, err = binutils.Int16(writtenData[0:binutils.Int16size])
			require.NoError(t, err)
			require.Equal(t, tt.args.data, value)
			writer.ResetBytesWritten()
			require.Equal(t, 0, writer.BytesWritten())
		})
	}
}

func TestBinaryWriter_WriteInt32(t *testing.T) {
	type args struct {
		data int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"write_000", args{0}, false},
		{"write_100", args{100}, false},
		{"write_min_int32", args{math.MinInt32}, false},
		{"write_max_int32", args{math.MaxInt32}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := bytes.NewBuffer(nil)
			writer := binutils.NewBinaryWriter(collector)
			err := writer.WriteInt32(tt.args.data)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, binutils.Int32size, writer.BytesWritten())
			if err != nil {
				return
			}
			writtenData := collector.Bytes()
			require.Equal(t, binutils.Int32size, len(writtenData))
			var (
				value int32
			)
			value, err = binutils.Int32(writtenData[0:binutils.Int32size])
			require.NoError(t, err)
			require.Equal(t, tt.args.data, value)
			writer.ResetBytesWritten()
			require.Equal(t, 0, writer.BytesWritten())
		})
	}
}

func TestBinaryWriter_WriteInt64(t *testing.T) {
	type args struct {
		data int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"write_000", args{0}, false},
		{"write_100", args{100}, false},
		{"write_min_int64", args{math.MinInt64}, false},
		{"write_max_int64", args{math.MaxInt64}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := bytes.NewBuffer(nil)
			writer := binutils.NewBinaryWriter(collector)
			err := writer.WriteInt64(tt.args.data)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, binutils.Int64size, writer.BytesWritten())
			if err != nil {
				return
			}
			writtenData := collector.Bytes()
			require.Equal(t, binutils.Int64size, len(writtenData))
			var (
				value int64
			)
			value, err = binutils.Int64(writtenData[0:binutils.Int64size])
			require.NoError(t, err)
			require.Equal(t, tt.args.data, value)
			writer.ResetBytesWritten()
			require.Equal(t, 0, writer.BytesWritten())
		})
	}
}

func TestBinaryWriter_WriteObject(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	writer := binutils.NewBinaryWriter(buffer)
	reader := binutils.NewBinaryReader(buffer)

	for _, tt := range readWriteTests {
		t.Run(tt.name, func(t *testing.T) {
			var actualHex string
			buffer.Reset()
			writer.ResetBytesWritten()
			reader.ResetBytesTaken()
			err := writer.WriteObject(tt.data)
			require.Equalf(t, tt.wantErr, err != nil, "expected error %v, got %v", tt.wantErr, err)
			if err != nil {
				return
			}
			require.Equalf(t, tt.expectedBytesWritten, buffer.Len(), "hex: %v", hex.EncodeToString(buffer.Bytes()))
			require.Equal(t, tt.expectedBytesWritten, writer.BytesWritten())
			actualHex, err = reader.ReadHex(buffer.Len())
			require.NoError(t, err)
			require.Equal(t, tt.expectedHex, actualHex)
		})
	}
}
