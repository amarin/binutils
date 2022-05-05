package binutils_test

import (
	"encoding/hex"
	"fmt"
	"math"
	"testing"

	. "github.com/amarin/binutils"
)

func Test_uint64Bytes(t *testing.T) {
	for _, tt := range []struct {
		name  string
		value uint64
		hex   string
	}{
		{"ok_0", 0, "0000000000000000"},
		{"ok_1", 1, "0000000000000001"},
		{"ok_max", math.MaxUint64, "ffffffffffffffff"},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if got := Uint64bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Uint64bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2uint64(t *testing.T) {
	for _, tt := range []struct {
		name      string
		value     uint64
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "0000000000000000", false},
		{"ok_1", 1, "0000000000000001", false},
		{"ok_max", math.MaxUint64, "ffffffffffffffff", false},
		{"nok_incorrect_size", math.MaxUint64, "ffffffffffffff", true},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
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
	for _, tt := range []struct {
		name  string
		value int64
		hex   string
	}{
		{"ok_0", 0, "0000000000000000"},
		{"ok_1", 1, "0000000000000001"},
		{"ok_neg_1", -1, "ffffffffffffffff"},
		{"ok_max", math.MaxInt64, "7fffffffffffffff"},
		{"ok_min", math.MinInt64, "8000000000000000"},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if got := Int64bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Int64bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_uint32Bytes(t *testing.T) {
	for _, tt := range []struct {
		name  string
		value uint32
		hex   string
	}{
		{"ok_0", 0, "00000000"},
		{"ok_1", 1, "00000001"},
		{"ok_max", math.MaxUint32, "ffffffff"},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if got := Uint32bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Uint32bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2uint32(t *testing.T) { // nolint:dupl
	for _, tt := range []struct { // nolint:maligned
		name      string
		value     uint32
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "00000000", false},
		{"ok_1", 1, "00000001", false},
		{"ok_max", math.MaxUint32, "ffffffff", false},
		{"nok_incorrect_size", 0, "ffffff", true},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
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
	for _, tt := range []struct {
		name  string
		value int32
		hex   string
	}{
		{"ok_0", 0, "00000000"},
		{"ok_1", 1, "00000001"},
		{"ok_neg_1", -1, "ffffffff"},
		{"ok_max", math.MaxInt32, "7fffffff"},
		{"ok_min", math.MinInt32, "80000000"},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if got := Int32bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Int32bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2int32(t *testing.T) { // nolint:dupl
	for _, tt := range []struct { // nolint:maligned
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
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
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
	for _, tt := range []struct {
		name  string
		value uint16
		hex   string
	}{
		{"ok_0", 0, "0000"},
		{"ok_1", 1, "0001"},
		{"ok_max", math.MaxUint16, hex.EncodeToString([]byte{255, 255})},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if got := Uint16bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Uint16bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2uint16(t *testing.T) {
	for _, tt := range []struct { // nolint:maligned
		name      string
		value     uint16
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "0000", false},
		{"ok_1", 1, "0001", false},
		{"ok_max", math.MaxUint16, hex.EncodeToString([]byte{255, 255}), false},
		{"nok_incorrect_size", 0, hex.EncodeToString([]byte{255}), true},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
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
	for _, tt := range []struct {
		name  string
		value int16
		hex   string
	}{
		{"ok_0", 0, "0000"},
		{"ok_1", 1, "0001"},
		{"ok_neg_1", -1, hex.EncodeToString([]byte{255, 255})},
		{"ok_max", math.MaxInt16, "7fff"},
		{"ok_min", math.MinInt16, "8000"},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if got := Int16bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Int16bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2int16(t *testing.T) { // nolint:dupl
	for _, tt := range []struct { // nolint:maligned
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
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
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
	for _, tt := range []struct {
		name  string
		value uint8
		hex   string
	}{
		{"ok_0", 0, "00"},
		{"ok_1", 1, "01"},
		{"ok_max", math.MaxUint8, "ff"},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if got := Uint8bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Uint8bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_Bytes2uint8(t *testing.T) { // nolint:dupl
	for _, tt := range []struct { // nolint:maligned
		name      string
		value     uint8
		hex       string
		wantError bool
	}{
		{"ok_0", 0, "00", false},
		{"ok_1", 1, "01", false},
		{"ok_max", math.MaxUint8, "ff", false},
		{"nok_incorrect_size_bigger", 0, "ffff", true},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
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
	for _, tt := range []struct {
		name  string
		value int8
		hex   string
	}{
		{"ok_0", 0, "00"},
		{"ok_1", 1, "01"},
		{"ok_neg_1", -1, "ff"},
		{"ok_max", math.MaxInt8, "7f"},
		{"ok_min", math.MinInt8, "80"},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if got := Int8bytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("Int8bytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func Test_bytes2int64(t *testing.T) { // nolint:dupl
	for _, tt := range []struct {
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
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
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

func Test_bytes2int8(t *testing.T) { // nolint:dupl
	for _, tt := range []struct { // nolint:maligned
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
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
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
		size := size // pin size
		t.Run(fmt.Sprintf("allocate_%v", size), func(t *testing.T) {
			size := size // pin size
			if got := AllocateBytes(size); len(got) != size {
				t.Errorf("AllocateBytes(%d) returns [%d]byte", size, len(got))
			}
		})
	}
}

func TestString(t *testing.T) {
	for _, tt := range []struct {
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
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
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
	for _, tt := range []struct {
		name  string
		value string
		hex   string
	}{
		{"ok_0", "0", "3000"},
		{"ok_1", "1", "3100"},
		{"ok_neg_1", "-1", "2d3100"},
		{"ok_ff", "ff", "666600"},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if got := StringBytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("StringBytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}

func TestRune(t *testing.T) {
	for _, tt := range []struct {
		name      string
		value     rune
		hex       string
		wantError bool
	}{
		{"latin_a", 'a', "00000061", false},
		{"latin_Z", 'Z', "0000005a", false},
		{"cyrillic_a", 'а', "00000430", false},
		{"cyrillic_YA", 'Я', "0000042f", false},
		{"chinese_one", '一', "00004e00", false},
		{"chinese_ai", '爱', "00007231", false},
		{"insufficient_bytes_caused_error", '-', "7fffff", true},
		{"extra_bytes_caused_error", '-', "7fffffffff", true},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if data, err := hex.DecodeString(tt.hex); err != nil {
				t.Errorf("cannt decode string %#v to bytes: %v", tt.hex, err)
			} else if got, err := Rune(data); (err != nil) != tt.wantError {
				t.Errorf("Rune(%v) = %v, %v, want error %v", tt.hex, got, err, tt.wantError)
			} else if err == nil && got != tt.value {
				t.Errorf("Rune(%v) = %v expect %v(%08x)", tt.hex, got, string(tt.value), tt.value)
			}
		})
	}
}

func TestRuneBytes(t *testing.T) {
	for _, tt := range []struct {
		name  string
		value rune
		hex   string
	}{
		{"latin_a", 'a', "00000061"},
		{"latin_Z", 'Z', "0000005a"},
		{"cyrillic_a", 'а', "00000430"},
		{"cyrillic_YA", 'Я', "0000042f"},
		{"chinese_one", '一', "00004e00"},
		{"chinese_ai", '爱', "00007231"},
	} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt
			if got := RuneBytes(tt.value); hex.EncodeToString(got) != tt.hex {
				t.Errorf("RuneBytes() = %v, want %v", hex.EncodeToString(got), tt.hex)
			}
		})
	}
}
