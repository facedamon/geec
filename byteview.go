package geec

// ByteView holds an immutable view of bytes
type ByteView struct {
	b []byte
}

// Len returns that view`s length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of data as a byte slice
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// Implement String interface
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
