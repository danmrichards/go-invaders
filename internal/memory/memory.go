package memory

// Basic is an in-memory implementation of the Space Invaders memory.
//
// The Space Invaders memory is mapped as follows:
//
// $0000-$1FFF -> 8K ROM
// $2000-$23FF -> 1K RAM
// $2400-$3FFF -> 7K Video RAM
// $4000 -> RAM mirror
//
// For more details on the ROM structure see machine.LoadROM.
type Basic []byte

// Read returns the value from memory at the given address.
func (b Basic) Read(addr uint16) byte {
	return b[addr]
}

// ReadAll returns the full memory contents.
func (b Basic) ReadAll() []byte {
	return b
}

// Write writes the value v into memory at the given address.
func (b Basic) Write(addr uint16, v byte) {
	b[addr] = v
}
