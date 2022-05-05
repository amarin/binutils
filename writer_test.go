package binutils_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/amarin/binutils"
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
			collector := binutils.NewBuffer(nil)
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
			collector := binutils.NewBuffer(nil)
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
			collector := binutils.NewBuffer(nil)
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
			collector := binutils.NewBuffer(nil)
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
			collector := binutils.NewBuffer(nil)
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
			collector := binutils.NewBuffer(nil)
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
			collector := binutils.NewBuffer(nil)
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
			collector := binutils.NewBuffer(nil)
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
			collector := binutils.NewBuffer(nil)
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
