package binutils

import (
	"testing"
)

func TestNewEmptyBuffer(t *testing.T) {
	got := NewEmptyBuffer()
	if got.buffer.Len() != 0 {
		t.Error("NewEmptyBuffer() makes non-empty buffer")
	}
}

func TestNewBuffer(t *testing.T) {
	tests := []struct {
		name      string
		data      []byte
		expectLen int
	}{
		{"empty", []byte{}, 0},
		{"one_byte", []byte{0}, 1},
		{"ten_bytes", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuffer(tt.data); got.buffer.Len() != tt.expectLen {
				t.Errorf("NewBuffer() buffer len mismatch %d != %d", got.buffer.Len(), tt.expectLen)
			}
		})
	}
}

func TestBuffer_Bytes(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{"ok_empty", []byte{}},
		{"ok_five", []byte{1, 2, 3, 4, 5}},
		{"ok_fifteen", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewBuffer(tt.want)
			if got := x.Bytes(); len(got) != len(tt.want) {
				t.Errorf("slice len mismatch, got %d expected %d in %v", len(got), len(tt.want), got)
			}
		})
	}
}
