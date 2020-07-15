package binutils_test

import (
	"math"
	"reflect"
	"testing"

	. "github.com/amarin/binutils"
)

func TestCalculateUseBitsPerIndex(t *testing.T) {
	for _, tt := range []struct { // nolint:maligned
		name     string
		sliceLen int
		reserve  bool
		want     BitsPerIndex
		wantErr  bool
	}{
		{"negative_len_error", -1, false, Use8bit, true},
		{"zero_len_use_uint8", 0, false, Use8bit, false},
		{"max_uint8_no_reserve", math.MaxUint8, false, Use8bit, false},
		{"max_uint8_reserve", math.MaxUint8, true, Use16bit, false},
		{"max_uint16_no_reserve", math.MaxUint16, false, Use16bit, false},
		{"max_uint16_reserve", math.MaxUint16, true, Use32bit, false},
		{"max_uint32_no_reserve", math.MaxUint32, false, Use32bit, false},
		{"max_uint32_reserve", math.MaxUint32, true, Use64bit, false},
		{"max_int8_no_reserve", math.MaxInt8, false, Use8bit, false},
		{"max_int8_reserve", math.MaxInt8, true, Use8bit, false},
		{"max_int16_no_reserve", math.MaxInt16, false, Use16bit, false},
		{"max_int16_reserve", math.MaxInt16, true, Use16bit, false},
		{"max_int32_no_reserve", math.MaxInt32, false, Use32bit, false},
		{"max_int32_reserve", math.MaxInt32, true, Use32bit, false},
		{"max_int64_no_reserve", math.MaxInt64, false, Use64bit, false},
		{"max_int64_reserve", math.MaxInt64, true, Use64bit, false},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			got, err := CalculateUseBitsPerIndex(tt.sliceLen, tt.reserve)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateUseBitsPerIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateUseBitsPerIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitsPerIndex_MarshalBinary(t *testing.T) {
	for _, tt := range []struct {
		name     string
		b        BitsPerIndex
		wantData []byte
		wantErr  bool
	}{
		{"uint8", Use8bit, []byte{8}, false},
		{"uint16", Use16bit, []byte{16}, false},
		{"uint32", Use32bit, []byte{32}, false},
		{"uint64", Use64bit, []byte{64}, false},
		{"uint7", BitsPerIndex(7), []byte{}, true},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if gotData, err := tt.b.MarshalBinary(); (err != nil) != tt.wantErr {
				t.Fatalf("MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil && !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("MarshalBinary() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestBitsPerIndex_UnmarshalBinary(t *testing.T) {
	for _, tt := range []struct {
		name     string
		b        BitsPerIndex
		wantData []byte
		wantErr  bool
	}{
		{"uint8", Use8bit, []byte{8}, false},
		{"uint16", Use16bit, []byte{16}, false},
		{"uint32", Use32bit, []byte{32}, false},
		{"uint64", Use64bit, []byte{64}, false},
		{"uint7_error", BitsPerIndex(0), []byte{7}, true},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			bitsPerIndex := new(BitsPerIndex)
			if err := bitsPerIndex.UnmarshalBinary(tt.wantData); (err != nil) != tt.wantErr {
				t.Fatalf("UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil && *bitsPerIndex != tt.b {
				t.Errorf("UnmarshalBinary() unmarshals = %#v, want %#v", *bitsPerIndex, tt.b)
			}
		})
	}
}

func TestWriteUint64ToBufferUsingBits(t *testing.T) {
	for _, tt := range []struct {
		name       string
		value      uint64
		usingBits  BitsPerIndex
		expectData string
		want       int
		wantErr    bool
	}{
		{"uint8",
			math.MaxUint8, Use8bit, "ff",
			Uint8size, false},
		{"uint8_overflow",
			math.MaxUint8 + 1, Use8bit, "",
			0, true},
		{"uint16",
			math.MaxUint16, Use16bit, "ffff",
			Uint16size, false},
		{"uint16_overflow",
			math.MaxUint16 + 1, Use16bit, "",
			0, true},
		{"uint32",
			math.MaxUint32, Use32bit, "ffffffff",
			Uint32size, false},
		{"uint32_overflow",
			math.MaxUint32 + 1, Use32bit, "",
			0, true},
		{"uint64",
			math.MaxUint32 + 1, Use64bit, "0000000100000000",
			8, false},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			buffer := NewEmptyBuffer()
			bytesWritten, err := WriteUint64ToBufferUsingBits(tt.value, buffer, tt.usingBits)

			switch {
			case (err != nil) != tt.wantErr:
				t.Fatalf("WriteUint64ToBufferUsingBits() error = %v, wantErr %v", err, tt.wantErr)
			case err != nil:
				return
			case bytesWritten != tt.want:
				t.Fatalf("WriteUint64ToBufferUsingBits() bytesWritten = %v, want %v", bytesWritten, tt.want)
			}

			if buffer.Hex() != tt.expectData {
				t.Fatalf("expected data:\n\t%v\ngot data:\n\t%v", tt.expectData, buffer.Hex())
			}
		})
	}
}

func TestReadUint64FromBufferUsingBits(t *testing.T) {
	for _, tt := range []struct {
		name       string
		value      uint64
		usingBits  BitsPerIndex
		expectData string
		want       int
		wantErr    bool
	}{
		{"uint8",
			math.MaxUint8, Use8bit, "ff",
			Uint8size, false},
		{"uint16",
			math.MaxUint16, Use16bit, "ffff",
			Uint16size, false},
		{"uint32",
			math.MaxUint32, Use32bit, "ffffffff",
			Uint32size, false},
		{"uint64",
			math.MaxUint32 + 1, Use64bit, "0000000100000000",
			8, false},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			buffer := NewEmptyBuffer()
			if _, err := buffer.WriteHex(tt.expectData); err != nil {
				t.Fatalf("cant prepare buffer data: %v", err)
			}
			var target uint64
			err := ReadUint64FromBufferUsingBits(&target, buffer, tt.usingBits)
			switch {
			case (err != nil) != tt.wantErr:
				t.Fatalf("ReadUint64FromBufferUsingBits() error = %v, wantErr %v", err, tt.wantErr)
			case err != nil:
				return
			case target != tt.value:
				t.Fatalf("ReadUint64FromBufferUsingBits() takes %d instead of %d", target, tt.value)
			}
		})
	}
}

func TestBitsPerIndex_UnmarshalFromBuffer(t *testing.T) {
	for _, tt := range []struct {
		name     string
		b        BitsPerIndex
		wantData []byte
		wantErr  bool
	}{
		{"uint8", Use8bit, []byte{UsingUint8Indexes}, false},
		{"uint16", Use16bit, []byte{UsingUint16Indexes}, false},
		{"uint32", Use32bit, []byte{UsingUint32Indexes}, false},
		{"uint64", Use64bit, []byte{UsingUint64Indexes}, false},
		{"uint9_error", BitsPerIndex(0), []byte{9}, true},
		{"uint17_error", BitsPerIndex(0), []byte{17}, true},
		{"uint33_error", BitsPerIndex(0), []byte{33}, true},
		{"uint65_error", BitsPerIndex(0), []byte{65}, true},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			bitsPerIndex := new(BitsPerIndex)
			buffer := NewBuffer(tt.wantData)
			if err := bitsPerIndex.UnmarshalFromBuffer(buffer); (err != nil) != tt.wantErr {
				t.Fatalf("UnmarshalFromBuffer() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil && *bitsPerIndex != tt.b {
				t.Errorf("UnmarshalFromBuffer() unmarshals = %#v, want %#v", *bitsPerIndex, tt.b)
			}
		})
	}
}
