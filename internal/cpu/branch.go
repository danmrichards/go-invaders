package cpu

// jmp is the "Jump" handler.
//
// This handler jumps the program counter to a given point in memory.
func (i *Intel8080) jmp() uint16 {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.twoByteRead()

	// As we're jumping the program counter there is no need to return a value
	// for the main cycle to increment the counter.
	return 0
}
