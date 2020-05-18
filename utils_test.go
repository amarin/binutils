package binutils

import (
	"encoding/hex"
	"fmt"
	"math"
	"testing"
)

func Test_uint64Bytes(t *testing.T) {
	tests := []struct {
		name  string
		value uint64
		hex   string
	}{
		{"ok_0", 0, "0000000000000000"},
		{"ok_1", 1, "0000000000000001"},
		{"ok_max", math.MaxUint64, "ffffffffffffffff"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Uint64bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2uint64(t *testing.T) {
	tests := []struct {
		name      string
		value     uint64
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "0000000000000000", false},
		{"ok_1", 1, "0000000000000001", false},
		{"ok_max", math.MaxUint64, "ffffffffffffffff", false},
		{"nok_incorrect_size", math.MaxUint64, "ffffffffffffff", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if data, err := hex.DecodeString(tt.hex); err != nil {
				t.Errorf("cannt decode string %#v to bytes: %v", tt.hex, err)
			} else if got, err := Uint64(data); (err != nil) != tt.wantError {
				t.Errorf("Uint64(%v) = %v, %v, want error %v", tt.hex, got, err, tt.wantError)
			} else if err == nil && got != tt.value {
				t.Errorf("Uint64(%v) = %v  expect %v", tt.hex, got, tt.value)
			}
		})
	}
}

func Test_int64Bytes(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		hex   string
	}{
		{"ok_0", 0, "0000000000000000"},
		{"ok_1", 1, "0000000000000001"},
		{"ok_neg_1", -1, "ffffffffffffffff"},
		{"ok_max", math.MaxInt64, "7fffffffffffffff"},
		{"ok_min", math.MinInt64, "8000000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Int64bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2int64(t *testing.T) {
	tests := []struct {
		name      string
		value     int64
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "0000000000000000", false},
		{"ok_1", 1, "0000000000000001", false},
		{"ok_neg_1", -1, "ffffffffffffffff", false},
		{"ok_max", math.MaxInt64, "7fffffffffffffff", false},
		{"ok_min", math.MinInt64, "8000000000000000", false},
		{"nok_incorrect_size", 0, "ffffffffffffff", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if data, err := hex.DecodeString(tt.hex); err != nil {
				t.Errorf("cannt decode string %#v to bytes: %v", tt.hex, err)
			} else if got, err := Int64(data); (err != nil) != tt.wantError {
				t.Errorf("Int64(%v) = %v, %v, want error %v", tt.hex, got, err, tt.wantError)
			} else if err == nil && got != tt.value {
				t.Errorf("Int64(%v) = %v  expect %v", tt.hex, got, tt.value)
			}
		})
	}
}

func Test_uint32Bytes(t *testing.T) {
	tests := []struct {
		name  string
		value uint32
		hex   string
	}{
		{"ok_0", 0, "00000000"},
		{"ok_1", 1, "00000001"},
		{"ok_max", math.MaxUint32, "ffffffff"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Uint32bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2uint32(t *testing.T) {
	tests := []struct {
		name      string
		value     uint32
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "00000000", false},
		{"ok_1", 1, "00000001", false},
		{"ok_max", math.MaxUint32, "ffffffff", false},
		{"nok_incorrect_size", 0, "ffffff", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if data, err := hex.DecodeString(tt.hex); err != nil {
				t.Errorf("cannt decode string %#v to bytes: %v", tt.hex, err)
			} else if got, err := Uint32(data); (err != nil) != tt.wantError {
				t.Errorf("Uint32(%v) = %v, %v, want error %v", tt.hex, got, err, tt.wantError)
			} else if err == nil && got != tt.value {
				t.Errorf("Uint32(%v) = %v  expect %v", tt.hex, got, tt.value)
			}
		})
	}
}

func Test_int32Bytes(t *testing.T) {
	tests := []struct {
		name  string
		value int32
		hex   string
	}{
		{"ok_0", 0, "00000000"},
		{"ok_1", 1, "00000001"},
		{"ok_neg_1", -1, "ffffffff"},
		{"ok_max", math.MaxInt32, "7fffffff"},
		{"ok_min", math.MinInt32, "80000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Int32bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2int32(t *testing.T) {
	tests := []struct {
		name      string
		value     int32
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "00000000", false},
		{"ok_1", 1, "00000001", false},
		{"ok_neg_1", -1, "ffffffff", false},
		{"ok_max", math.MaxInt32, "7fffffff", false},
		{"ok_min", math.MinInt32, "80000000", false},
		{"nok_incorrect_size_bigger", 0, "ffffffffff", true},
		{"nok_incorrect_size_less", 0, "ffff", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if data, err := hex.DecodeString(tt.hex); err != nil {
				t.Errorf("cannt decode string %#v to bytes: %v", tt.hex, err)
			} else if got, err := Int32(data); (err != nil) != tt.wantError {
				t.Errorf("Int32(%v) = %v, %v, want error %v", tt.hex, got, err, tt.wantError)
			} else if err == nil && got != tt.value {
				t.Errorf("Int32(%v) = %v  expect %v", tt.hex, got, tt.value)
			}
		})
	}
}

func Test_uint16Bytes(t *testing.T) {
	tests := []struct {
		name  string
		value uint16
		hex   string
	}{
		{"ok_0", 0, "0000"},
		{"ok_1", 1, "0001"},
		{"ok_max", math.MaxUint16, "ffff"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Uint16bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2uint16(t *testing.T) {
	tests := []struct {
		name      string
		value     uint16
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "0000", false},
		{"ok_1", 1, "0001", false},
		{"ok_max", math.MaxUint16, "ffff", false},
		{"nok_incorrect_size", 0, "ff", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if data, err := hex.DecodeString(tt.hex); err != nil {
				t.Errorf("cannt decode string %#v to bytes: %v", tt.hex, err)
			} else if got, err := Uint16(data); (err != nil) != tt.wantError {
				t.Errorf("Uint16(%v) = %v, %v, want error %v", tt.hex, got, err, tt.wantError)
			} else if err == nil && got != tt.value {
				t.Errorf("Uint16(%v) = %v  expect %v", tt.hex, got, tt.value)
			}
		})
	}
}

func Test_int16Bytes(t *testing.T) {
	tests := []struct {
		name  string
		value int16
		hex   string
	}{
		{"ok_0", 0, "0000"},
		{"ok_1", 1, "0001"},
		{"ok_neg_1", -1, "ffff"},
		{"ok_max", math.MaxInt16, "7fff"},
		{"ok_min", math.MinInt16, "8000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Int16bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2int16(t *testing.T) {
	tests := []struct {
		name      string
		value     int16
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "0000", false},
		{"ok_1", 1, "0001", false},
		{"ok_neg_1", -1, "ffff", false},
		{"ok_max", math.MaxInt16, "7fff", false},
		{"ok_min", math.MinInt16, "8000", false},
		{"nok_incorrect_size_bigger", 0, "ffffffff", true},
		{"nok_incorrect_size_less", 0, "ff", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if data, err := hex.DecodeString(tt.hex); err != nil {
				t.Errorf("cannt decode string %#v to bytes: %v", tt.hex, err)
			} else if got, err := Int16(data); (err != nil) != tt.wantError {
				t.Errorf("Int16(%v) = %v, %v, want error %v", tt.hex, got, err, tt.wantError)
			} else if err == nil && got != tt.value {
				t.Errorf("Int16(%v) = %v  expect %v", tt.hex, got, tt.value)
			}
		})
	}
}

func Test_Uint8bytes(t *testing.T) {
	tests := []struct {
		name  string
		value uint8
		hex   string
	}{
		{"ok_0", 0, "00"},
		{"ok_1", 1, "01"},
		{"ok_max", math.MaxUint8, "ff"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Uint8bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_Bytes2uint8(t *testing.T) {
	tests := []struct {
		name      string
		value     uint8
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "00", false},
		{"ok_1", 1, "01", false},
		{"ok_max", math.MaxUint8, "ff", false},
		{"nok_incorrect_size_bigger", 0, "ffff", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if data, err := hex.DecodeString(tt.hex); err != nil {
				t.Errorf("cannt decode string %#v to bytes: %v", tt.hex, err)
			} else if got, err := Uint8(data); (err != nil) != tt.wantError {
				t.Errorf("Int8(%v) = %v, %v, want error %v", tt.hex, got, err, tt.wantError)
			} else if err == nil && got != tt.value {
				t.Errorf("Int8(%v) = %v  expect %v", tt.hex, got, tt.value)
			}
		})
	}
}

func Test_int8Bytes(t *testing.T) {
	tests := []struct {
		name  string
		value int8
		hex   string
	}{
		{"ok_0", 0, "00"},
		{"ok_1", 1, "01"},
		{"ok_neg_1", -1, "ff"},
		{"ok_max", math.MaxInt8, "7f"},
		{"ok_min", math.MinInt8, "80"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Int8bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2int8(t *testing.T) {
	tests := []struct {
		name      string
		value     int8
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "00", false},
		{"ok_1", 1, "01", false},
		{"ok_neg_1", -1, "ff", false},
		{"ok_max", math.MaxInt8, "7f", false},
		{"ok_min", math.MinInt8, "80", false},
		{"nok_incorrect_size_bigger", 0, "ffff", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if data, err := hex.DecodeString(tt.hex); err != nil {
				t.Errorf("cannt decode string %#v to bytes: %v", tt.hex, err)
			} else if got, err := Int8(data); (err != nil) != tt.wantError {
				t.Errorf("Int8(%v) = %v, %v, want error %v", tt.hex, got, err, tt.wantError)
			} else if err == nil && got != tt.value {
				t.Errorf("Int8(%v) = %v  expect %v", tt.hex, got, tt.value)
			}
		})
	}
}

func TestAllocateBytes(t *testing.T) {
	tests := []int{0, 1, 2, 3, 4, 5, 7, 8, 11, 13, 16, 17, 23, 29, 32, 37}
	for _, size := range tests {
		t.Run(fmt.Sprintf("allocate_%v", size), func(t *testing.T) {
			if got := AllocateBytes(size); len(got) != size {
				t.Errorf("AllocateBytes(%d) returns [%d]byte", size, len(got))
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		hex       string
		wantError bool
	}{
		{"ok_0", "0", "3000", false},
		{"ok_1", "1", "3100", false},
		{"ok_neg_1", "-1", "2d3100", false},
		{"ok_neg_1", "ff", "666600", false},
		{"nok_no_termination_byte", "", "ff", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if data, err := hex.DecodeString(tt.hex); err != nil {
				t.Errorf("cannt decode string %#v to bytes: %v", tt.hex, err)
			} else if got, err := String(data); (err != nil) != tt.wantError {
				t.Errorf("Int8(%v) = %v, %v, want error %v", tt.hex, got, err, tt.wantError)
			} else if err == nil && got != tt.value {
				t.Errorf("Int8(%v) = %v  expect %v", tt.hex, got, tt.value)
			}
		})
	}
}

func TestStringBytes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		hex   string
	}{
		{"ok_0", "0", "3000"},
		{"ok_1", "1", "3100"},
		{"ok_neg_1", "-1", "2d3100"},
		{"ok_ff", "ff", "666600"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringBytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("StringBytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}
