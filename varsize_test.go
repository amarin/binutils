package binutils_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/amarin/binutils"
)

func TestCalculateUseBitsPerIndex(t *testing.T) {
	for _, tt := range []struct { // nolint:maligned
		name     string
		sliceLen int
		reserve  bool
		want     binutils.BitsPerIndex
		wantErr  bool
	}{
		{"negative_len_error", -1, false, binutils.Use8bit, true},
		{"zero_len_use_uint8", 0, false, binutils.Use8bit, false},
		{"max_uint8_no_reserve", math.MaxUint8, false, binutils.Use8bit, false},
		{"max_uint8_reserve", math.MaxUint8, true, binutils.Use16bit, false},
		{"max_uint16_no_reserve", math.MaxUint16, false, binutils.Use16bit, false},
		{"max_uint16_reserve", math.MaxUint16, true, binutils.Use32bit, false},
		{"max_uint32_no_reserve", math.MaxUint32, false, binutils.Use32bit, false},
		{"max_uint32_reserve", math.MaxUint32, true, binutils.Use64bit, false},
		{"max_int8_no_reserve", math.MaxInt8, false, binutils.Use8bit, false},
		{"max_int8_reserve", math.MaxInt8, true, binutils.Use8bit, false},
		{"max_int16_no_reserve", math.MaxInt16, false, binutils.Use16bit, false},
		{"max_int16_reserve", math.MaxInt16, true, binutils.Use16bit, false},
		{"max_int32_no_reserve", math.MaxInt32, false, binutils.Use32bit, false},
		{"max_int32_reserve", math.MaxInt32, true, binutils.Use32bit, false},
		{"max_int64_no_reserve", math.MaxInt64, false, binutils.Use64bit, false},
		{"max_int64_reserve", math.MaxInt64, true, binutils.Use64bit, false},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			got, err := binutils.CalculateUseBitsPerIndex(tt.sliceLen, tt.reserve)
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
		b        binutils.BitsPerIndex
		wantData []byte
		wantErr  bool
	}{
		{"uint8", binutils.Use8bit, []byte{8}, false},
		{"uint16", binutils.Use16bit, []byte{16}, false},
		{"uint32", binutils.Use32bit, []byte{32}, false},
		{"uint64", binutils.Use64bit, []byte{64}, false},
		{"uint7", binutils.BitsPerIndex(7), []byte{}, true},
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
		b        binutils.BitsPerIndex
		wantData []byte
		wantErr  bool
	}{
		{"uint8", binutils.Use8bit, []byte{8}, false},
		{"uint16", binutils.Use16bit, []byte{16}, false},
		{"uint32", binutils.Use32bit, []byte{32}, false},
		{"uint64", binutils.Use64bit, []byte{64}, false},
		{"uint7_error", binutils.BitsPerIndex(0), []byte{7}, true},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			bitsPerIndex := new(binutils.BitsPerIndex)
			if err := bitsPerIndex.UnmarshalBinary(tt.wantData); (err != nil) != tt.wantErr {
				t.Fatalf("UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil && *bitsPerIndex != tt.b {
				t.Errorf("UnmarshalBinary() unmarshals = %#v, want %#v", *bitsPerIndex, tt.b)
			}
		})
	}
}
