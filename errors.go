package binutils

import (
	"errors"
)

var (
	ErrExpected1   = errors.New("expected 1 byte")              // expected exactly 1 byte
	ErrExpected2   = errors.New("expected 2 bytes")             // expected exactly 2 bytes
	ErrExpected4   = errors.New("expected 4 bytes")             // expected exactly 4 bytes
	ErrExpected8   = errors.New("expected 8 bytes")             // expected exactly 8 bytes
	ErrMinimum1    = errors.New("at least 1 byte required")     // expected at least 1 byte
	ErrRequired0T  = errors.New("required 0-terminated string") // expected 0-byte termination
	ErrSizing      = errors.New("unexpected sizing")            // sizing unexpected
	ErrNegativeLen = errors.New("negative length")              // negative length
	ErrOverflowBy  = errors.New("type overflow")                // value overflows type
	ErrBuffer      = errors.New("buffer")                       // general buffer error
	ErrDecodeTo    = errors.New("decode")                       // decode error
	ErrMissedData  = errors.New("insufficient bytes")           // some bytes missed
	ErrExtraData   = errors.New("extra bytes")                  // extra bytes after decoding means error
	ErrRead        = errors.New("read")                         // general read error
)
