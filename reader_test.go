package binutils_test

import (
	"bytes"
	"testing"

	. "github.com/amarin/binutils"
)

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
			switch {
			case (err != nil) != tt.wantErr:
				t.Fatalf("ReadHex(%v) err %v, want %v", tt.wantBytes, err, tt.wantErr)
			case err != nil:
				return
			case got != tt.expectString:
				t.Fatalf("ReadHex(%v) = %v, expect %v", tt.wantBytes, got, tt.expectString)
			}
		})
	}
}
