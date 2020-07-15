package binutils

import (
	"errors"
	"fmt"
)

var (
	ErrExpected1byte                = errors.New("expected 1 byte")
	ErrExpected2bytes               = errors.New("expected 2 bytes")
	ErrExpected4bytes               = errors.New("expected 4 bytes")
	ErrExpected8bytes               = errors.New("expected 8 bytes")
	ErrExpectedAtLeast1Byte         = errors.New("at least 1 byte required")
	ErrRequiredZeroTerminatedString = errors.New("required 0-terminated sring")
)

// Error type implements error interface and used in package routines.
type Error struct {
	wrapped     error
	description string
}

// Error returns formatted error string. Implements error interface.
func (e Error) Error() string {
	if e.wrapped != nil {
		return e.description + ": " + e.wrapped.Error()
	}

	return e.description
}

// GoString returns formatted error string prefixes with package name and type. Implements GoStringer.
func (e Error) GoString() string {
	return "binutils.Error:" + e.Error()
}

// NewError is a default error constructor.
// Takes format and arguments to make error description using fmt.Sprintf() formatting rules.
func NewError(format string, args ...interface{}) error {
	return &Error{description: fmt.Sprintf(format, args...)}
}

// WrapError creates new error, wrapping underlying error with formatted string.
func WrapError(wrap error, format string, args ...interface{}) *Error {
	return &Error{description: fmt.Sprintf(format, args...), wrapped: wrap}
}
