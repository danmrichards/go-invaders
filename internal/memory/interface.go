package memory

// Reader is the interface that wraps the basic Read and ReadAll methods.
//
// Read returns the value from memory at the given address.
//
// ReadAll returns the full memory contents.
type Reader interface {
	Read(addr uint16) byte
	ReadAll() []byte
}

// Writer is the interface that wraps the basic Write method.
//
// Write writes the value v into memory at the given address.
type Writer interface {
	Write(addr uint16, v byte)
}

// ReadWriteDumper is the interface that groups the basic Read, Write and Dump
// methods.
type ReadWriteDumper interface {
	Reader
	Writer
}
