package cpu

// jmp is the "Jump" handler.
//
// This handler jumps the program counter to a given point in memory.
func (i *Intel8080) jmp() uint16 {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.twoByteRead()

	// As we're jumping the program counter there is no need to increment the
	// program counter once this method returns.
	return 0
}

// call is the "Call subroutine" handler.
//
// A call operation is unconditionally performed to subroutine sub.
func (i *Intel8080) call() uint16 {
	i.stackAdd(i.pc + 2)

	// Jump the program to the subroutine indicated by the two immediate bytes
	// in memory.
	i.pc = i.twoByteRead()

	// As we're unconditionally setting the program counter above, there is no
	// need to increment the program counter once this method returns.
	return 0
}
