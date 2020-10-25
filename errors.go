package binutils

import (
	"errors"
)

// Some predefined errors used during processing.
var (
	// ErrExpected1 returned if expected exactly 1 byte.
	ErrExpected1 = errors.New("expected 1 byte")

	// ErrExpected2 returned if expected exactly 2 bytes.
	ErrExpected2 = errors.New("expected 2 bytes")

	// ErrExpected4 returned if expected exactly 4 bytes.
	ErrExpected4 = errors.New("expected 4 bytes")

	// ErrExpected8 returned if expected exactly 8 bytes.
	ErrExpected8 = errors.New("expected 8 bytes")

	// ErrMinimum1 returned if expected at least 1 byte.
	ErrMinimum1 = errors.New("at least 1 byte required")

	// ErrRequired0T returned if expected 0-byte termination.
	ErrRequired0T = errors.New("required 0-terminated string")

	// ErrSizing returned if sizing unexpected.
	ErrSizing = errors.New("unexpected sizing")

	// ErrNegativeLen returned if negative length.
	ErrNegativeLen = errors.New("negative length")

	// ErrOverflowBy returned if value overflows type.
	ErrOverflowBy = errors.New("type overflow")

	// ErrBuffer returned if general buffer error.
	ErrBuffer = errors.New("buffer")

	// ErrDecodeTo returned if decode error.
	ErrDecodeTo = errors.New("decode")

	// ErrMissedData returned if some bytes missed.
	ErrMissedData = errors.New("insufficient bytes")

	// ErrExtraData returned if extra bytes after decoding means error.
	ErrExtraData = errors.New("extra bytes")

	// ErrRead returned if general read error.
	ErrRead = errors.New("read")

	// ErrWrite returned if general write error.
	ErrWrite = errors.New("write")

	// ErrClose returned if general close error.
	ErrClose = errors.New("close")
)
