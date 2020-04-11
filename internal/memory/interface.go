package memory

// Reader is the interface that wraps the basic Read method.
//
// Read returns the value from memory at the given address.
type Reader interface {
	Read(addr uint16) byte
}

// Writer is the interface that wraps the basic Write method.
//
// Write writes the value v into memory at the given address.
type Writer interface {
	Write(addr uint16, v byte)
}

// Dumper is the interface that wraps the basic Dump method.
//
// Dump returns a full dump of the memory contents.
type Dumper interface {
	Dump() []byte
}

// ReadWriteDumper is the interface that groups the basic Read, Write and Dump
// methods.
type ReadWriteDumper interface {
	Reader
	Writer
	Dumper
}
