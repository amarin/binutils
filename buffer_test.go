package binutils_test

import (
	"strings"
	"testing"

	"github.com/amarin/binutils"
)

func TestNewEmptyBuffer(t *testing.T) {
	got := binutils.NewEmptyBuffer()
	if got.Len() != 0 {
		t.Error("NewEmptyBuffer() makes non-empty buffer")
	}
}

func TestNewBuffer(t *testing.T) {
	for _, tt := range []struct {
		name      string
		data      []byte
		expectLen int
	}{
		{"empty", []byte{}, 0},
		{"one_byte", []byte{0}, 1},
		{"ten_bytes", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 10},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt again
			if got := binutils.NewBuffer(tt.data); got.Len() != tt.expectLen {
				t.Errorf("NewBuffer() buffer len mismatch %d != %d", got.Len(), tt.expectLen)
			}
		})
	}
}

func TestBuffer_Bytes(t *testing.T) {
	for _, tt := range []struct {
		name string
		want []byte
	}{
		{"ok_empty", []byte{}},
		{"ok_five", []byte{1, 2, 3, 4, 5}},
		{"ok_fifteen", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5}},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt again
			x := binutils.NewBuffer(tt.want)
			if got := x.Bytes(); len(got) != len(tt.want) { //
				t.Errorf("slice len mismatch, got %d expected %d in %v", len(got), len(tt.want), got)
			}
		})
	}
}

func TestBuffer_Hex(t *testing.T) {
	for _, tt := range []struct {
		name   string
		buffer *binutils.Buffer
		want   string
	}{
		{"empty_buffer", binutils.NewEmptyBuffer(), ""},
		{"single_byte", binutils.NewBuffer([]byte{1}), "01"},
		{"couple_of_bytes", binutils.NewBuffer([]byte{1, 1}), "0101"},
		{"three_of_bytes", binutils.NewBuffer([]byte{1, 1, 255}), "0101ff"},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if got := tt.buffer.Hex(); got != tt.want {
				t.Errorf("Hex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_WriteHex(t *testing.T) {
	for _, tt := range []struct {
		hexString string
		want      int
		wantErr   bool
	}{
		{"01", 1, false},
		{"ff", 1, false},
		{"ffff", 2, false},
		{"ffffff", 3, false},
		{"FF", 1, false},
		{"FEEF", 2, false},
		{"FFFFFF", 3, false},
		{"not_a_hex", 0, true}, // not a hex string al all
		{"fff", 0, true},       // wrong hex string
	} {
		tt := tt // pin tt
		t.Run(tt.hexString, func(t *testing.T) {
			tt := tt // pin tt
			x := binutils.NewEmptyBuffer()
			got, err := x.WriteHex(tt.hexString)
			switch {
			case (err != nil) != tt.wantErr:
				t.Fatalf("WriteHex() error = %v, wantErr %v", err, tt.wantErr)
			case got != tt.want:
				t.Fatalf("WriteHex() got = %v, want %v", got, tt.want)
			case err == nil && x.Hex() != strings.ToLower(tt.hexString):
				t.Fatalf("WriteHex() error nil, expected data %v, have %v", tt.hexString, x.Hex())
			}
		})
	}
}

func TestBuffer_ReadHex(t *testing.T) {
	for _, tt := range []struct {
		hexString string
		want      int
		wantErr   bool
	}{
		{"01", 1, false},
		{"ff", 1, false},
		{"ffff", 2, false},
		{"ffffff", 3, false},
	} {
		tt := tt // pin tt
		t.Run(tt.hexString, func(t *testing.T) {
			tt := tt // pin tt
			buffer := binutils.NewEmptyBuffer()
			if _, err := buffer.WriteHex(tt.hexString); err != nil {
				t.Fatalf("cant prepare test data: %v", err)
			}

			var target string
			if err := buffer.ReadHex(&target, tt.want); err != nil {
				t.Fatalf("cant read data as hex: %v", err)
			} else if target != tt.hexString {
				t.Fatalf("expected %v got %v", tt.hexString, target)
			}
		})
	}
}
