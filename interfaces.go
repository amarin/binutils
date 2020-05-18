package binutils

// BufferUnmarshaler defines interface to objects able to unmarshal itself from binary buffer.
// Such objects should read only own bytes from buffer leaving extra bytes intact for others.
type BufferUnmarshaler interface {
	UnmarshalFromBuffer(*Buffer) error
}
