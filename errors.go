package binutils

import (
	"errors"
	"fmt"
)

// Some predefined errors used during processing.
var (
	// Error indicates any binutils errors.
	Error = errors.New("binutils")

	// ErrNilPointer indicates nil pointer received when required valid pointer of specified type.
	ErrNilPointer = fmt.Errorf("%w: nill pointer", Error)

	// ErrExpected1 returned if expected exactly 1 byte.
	ErrExpected1 = fmt.Errorf("%w: expected 1 byte", Error)

	// ErrExpected2 returned if expected exactly 2 bytes.
	ErrExpected2 = fmt.Errorf("%w: expected 2 bytes", Error)

	// ErrExpected4 returned if expected exactly 4 bytes.
	ErrExpected4 = fmt.Errorf("%w: expected 4 bytes", Error)

	// ErrExpected8 returned if expected exactly 8 bytes.
	ErrExpected8 = fmt.Errorf("%w: expected 8 bytes", Error)

	// ErrMinimum1 returned if expected at least 1 byte.
	ErrMinimum1 = fmt.Errorf("%w: at least 1 byte required", Error)

	// ErrRequired0T returned if expected 0-byte termination.
	ErrRequired0T = fmt.Errorf("%w: required 0-terminated string", Error)

	// ErrDecodeTo returned if decode error.
	ErrDecodeTo = fmt.Errorf("%w: decode", Error)

	// ErrRead returned if general read error.
	ErrRead = fmt.Errorf("%w: read", Error)

	// ErrClose returned if general close error.
	ErrClose = fmt.Errorf("%w: close", Error)
)
