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

// ret is the "Return" handler.
//
// A return operation is unconditionally performed.
func (i *Intel8080) ret() uint16 {
	i.pc = i.stackPop()

	// Returning zero as we've manually set the program counter.
	return 0
}

// jnz is the "Jump If Not Zero" handler.
//
// If the zero bit is one, program execution continues at the memory address adr.
func (i *Intel8080) jnz() uint16 {
	// Return early if the zero bit isn't set.
	if !i.cc.z {
		return defaultInstructionLen
	}

	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.twoByteRead()

	// As we're jumping the program counter there is no need to increment the
	// program counter once this method returns.
	return 0
}

// jz is the "Jump Zero" handler.
//
// If the zero bit is not one, program execution continues at the memory address
// adr.
func (i *Intel8080) jz() uint16 {
	// Return early if the zero bit is set.
	if i.cc.z {
		return defaultInstructionLen
	}

	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.twoByteRead()

	// As we're jumping the program counter there is no need to increment the
	// program counter once this method returns.
	return 0
}

// jnc is the "Jump Not Carry" handler.
//
// If the carry bit is one, program execution continues at the memory address
// adr.
func (i *Intel8080) jnc() uint16 {
	// Return early if the carry bit is set.
	if i.cc.cy {
		return defaultInstructionLen
	}

	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.twoByteRead()

	// As we're jumping the program counter there is no need to increment the
	// program counter once this method returns.
	return 0
}

// jc is the "Jump Carry" handler.
//
// If the carry bit is not one, program execution continues at the memory
// address adr.
func (i *Intel8080) jc() uint16 {
	// Return early if the carry bit isn't set.
	if !i.cc.cy {
		return defaultInstructionLen
	}

	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.twoByteRead()

	// As we're jumping the program counter there is no need to increment the
	// program counter once this method returns.
	return 0
}

// jpo is the "Jump If Parity Odd" handler.
//
// If the Parity bit is zero (indicating a result with odd parity), program
// execution continues at the memory address adr.
func (i *Intel8080) jpo() uint16 {
	// Return early if the parity bit is set.
	if i.cc.p {
		return defaultInstructionLen
	}

	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.twoByteRead()

	// As we're jumping the program counter there is no need to increment the
	// program counter once this method returns.
	return 0
}

// jpe is the "Jump If Parity Even" handler.
//
// If the Parity bit is one (indicating a result with even parity), program
// execution continues at the memory address adr.
func (i *Intel8080) jpe() uint16 {
	// Return early if the parity bit isn't set.
	if !i.cc.p {
		return defaultInstructionLen
	}

	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.twoByteRead()

	// As we're jumping the program counter there is no need to increment the
	// program counter once this method returns.
	return 0
}

// jp is the "Jump If Positive" handler.
//
// If the Sign bit is zero (indicating a positive result), program execution
// continues at the memory address adr.
func (i *Intel8080) jp() uint16 {
	// Return early if the parity bit is set.
	if i.cc.s {
		return defaultInstructionLen
	}

	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.twoByteRead()

	// As we're jumping the program counter there is no need to increment the
	// program counter once this method returns.
	return 0
}

// jm is the "Jump If Minus" handler.
//
// If the Sign bit is one (indicating a positive result), program execution
// continues at the memory address adr.
func (i *Intel8080) jm() uint16 {
	// Return early if the parity bit is set.
	if !i.cc.s {
		return defaultInstructionLen
	}

	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.twoByteRead()

	// As we're jumping the program counter there is no need to increment the
	// program counter once this method returns.
	return 0
}

// cz is the "Call If Zero" handler.
//
// If the Zero bit is zero, a call operation is performed to subroutine sub.
func (i *Intel8080) cz() uint16 {
	// Return early if the zero bit isn't set.
	if !i.cc.z {
		return defaultInstructionLen
	}

	i.stackAdd(i.pc + 2)

	// Jump the program to the subroutine indicated by the two immediate bytes
	// in memory.
	i.pc = i.twoByteRead()

	// As we're unconditionally setting the program counter above, there is no
	// need to increment the program counter once this method returns.
	return 0
}

// cnz is the "Call If Not Zero" handler.
//
// If the Zero bit is one, a call operation is performed to subroutine sub.
func (i *Intel8080) cnz() uint16 {
	// Return early if the zero bit is set.
	if i.cc.z {
		return defaultInstructionLen
	}

	i.stackAdd(i.pc + 2)

	// Jump the program to the subroutine indicated by the two immediate bytes
	// in memory.
	i.pc = i.twoByteRead()

	// As we're unconditionally setting the program counter above, there is no
	// need to increment the program counter once this method returns.
	return 0
}

// cc is the "Call If Carry" handler.
//
// If the Carry bit is zero, a call operation is performed to subroutine sub.
func (i *Intel8080) cic() uint16 {
	// Return early if the carry bit isn't set.
	if !i.cc.cy {
		return defaultInstructionLen
	}

	i.stackAdd(i.pc + 2)

	// Jump the program to the subroutine indicated by the two immediate bytes
	// in memory.
	i.pc = i.twoByteRead()

	// As we're unconditionally setting the program counter above, there is no
	// need to increment the program counter once this method returns.
	return 0
}

// cnc is the "Call If Not carry" handler.
//
// If the carry bit is one, a call operation is performed to subroutine sub.
func (i *Intel8080) cnc() uint16 {
	// Return early if the carry bit is set.
	if i.cc.cy {
		return defaultInstructionLen
	}

	i.stackAdd(i.pc + 2)

	// Jump the program to the subroutine indicated by the two immediate bytes
	// in memory.
	i.pc = i.twoByteRead()

	// As we're unconditionally setting the program counter above, there is no
	// need to increment the program counter once this method returns.
	return 0
}

// cpo is the "Call If Parity Odd" handler.
//
// If the Parity bit is one (indicating a result with even parity), a call
// operation is performed to subroutine sub.
func (i *Intel8080) cpo() uint16 {
	// Return early if the parity bit isn't set.
	if !i.cc.p {
		return defaultInstructionLen
	}

	i.stackAdd(i.pc + 2)

	// Jump the program to the subroutine indicated by the two immediate bytes
	// in memory.
	i.pc = i.twoByteRead()

	// As we're unconditionally setting the program counter above, there is no
	// need to increment the program counter once this method returns.
	return 0
}

// cpe is the "Call If Parity Even" handler.
//
// If the Parity bit is even (indicating a result with even parity), a call
// operation is performed to subroutine sub.
func (i *Intel8080) cpe() uint16 {
	// Return early if the parity bit is set.
	if i.cc.p {
		return defaultInstructionLen
	}

	i.stackAdd(i.pc + 2)

	// Jump the program to the subroutine indicated by the two immediate bytes
	// in memory.
	i.pc = i.twoByteRead()

	// As we're unconditionally setting the program counter above, there is no
	// need to increment the program counter once this method returns.
	return 0
}

// cp is the "Call If Positive" handler.
//
// If the Sign bit is zero (indicating a positive result), a call operation is
// performed to subroutine sub.
func (i *Intel8080) cp() uint16 {
	// Return early if the sign bit isn't set.
	if i.cc.s {
		return defaultInstructionLen
	}

	i.stackAdd(i.pc + 2)

	// Jump the program to the subroutine indicated by the two immediate bytes
	// in memory.
	i.pc = i.twoByteRead()

	// As we're unconditionally setting the program counter above, there is no
	// need to increment the program counter once this method returns.
	return 0
}

// cp is the "Call If Minus" handler.
//
// If the Sign bit is one (indicating a positive result), a call operation is
// performed to subroutine sub.
func (i *Intel8080) cm() uint16 {
	// Return early if the sign bit is set.
	if !i.cc.s {
		return defaultInstructionLen
	}

	i.stackAdd(i.pc + 2)

	// Jump the program to the subroutine indicated by the two immediate bytes
	// in memory.
	i.pc = i.twoByteRead()

	// As we're unconditionally setting the program counter above, there is no
	// need to increment the program counter once this method returns.
	return 0
}

// rnz is the "Return If Not Zero" handler.
//
// If the Zero bit is zero, a return operation is performed.
func (i *Intel8080) rnz() uint16 {
	// Return early if the zero bit is set.
	if i.cc.z {
		return defaultInstructionLen
	}

	i.pc = i.stackPop()

	// Returning zero as we've manually set the program counter.
	return 0
}

// rz is the "Return If Zero" handler.
//
// If the Zero bit is one, a return operation is performed.
func (i *Intel8080) rz() uint16 {
	// Return early if the zero bit is set.
	if !i.cc.z {
		return defaultInstructionLen
	}

	i.pc = i.stackPop()

	// Returning zero as we've manually set the program counter.
	return 0
}

// rnc is the "Return If Not Carry" handler.
//
// If the Carry bit is zero, a return operation is performed.
func (i *Intel8080) rnc() uint16 {
	// Return early if the carry bit is set.
	if i.cc.cy {
		return defaultInstructionLen
	}

	i.pc = i.stackPop()

	// Returning zero as we've manually set the program counter.
	return 0
}

// rc is the "Return If Carry" handler.
//
// If the Carry bit is one, a return operation is performed.
func (i *Intel8080) rc() uint16 {
	// Return early if the carry bit is set.
	if !i.cc.cy {
		return defaultInstructionLen
	}

	i.pc = i.stackPop()

	// Returning zero as we've manually set the program counter.
	return 0
}

// rpo is the "Return If Parity Odd" handler.
//
// If the Parity bit is zero (indicating a result with odd parity), a return
// operation is performed.
func (i *Intel8080) rpo() uint16 {
	// Return early if the parity bit is set.
	if i.cc.p {
		return defaultInstructionLen
	}

	i.pc = i.stackPop()

	// Returning zero as we've manually set the program counter.
	return 0
}

// rpe is the "Return If Parity Even" handler.
//
// If the Parity bit is one (indicating a result with event parity), a return
// operation is performed.
func (i *Intel8080) rpe() uint16 {
	// Return early if the parity bit isn't set.
	if !i.cc.p {
		return defaultInstructionLen
	}

	i.pc = i.stackPop()

	// Returning zero as we've manually set the program counter.
	return 0
}

// rp is the "Return If Positive" handler.
//
// If the Sign bit is zero (indicating a positive result), a return operation
// is performed.
func (i *Intel8080) rp() uint16 {
	// Return early if the sign bit is set.
	if i.cc.s {
		return defaultInstructionLen
	}

	i.pc = i.stackPop()

	// Returning zero as we've manually set the program counter.
	return 0
}

// rm is the "Return If Minus" handler.
//
// If the Sign bit is one (indicating a negative result), a return operation
// is performed.
func (i *Intel8080) rm() uint16 {
	// Return early if the sign bit isn't set.
	if !i.cc.s {
		return defaultInstructionLen
	}

	i.pc = i.stackPop()

	// Returning zero as we've manually set the program counter.
	return 0
}

// pchl is the "Load Program Counter" handler.
//
// The contents of the H register replaces the most significant 8 bits of the
// program counter, and the contents of the L register replace the least
// significant 8 bits of the program counter.
//
// This causes program execution to continue at the address contained in the H
// and L registers.
func (i *Intel8080) pchl() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	i.pc = addr

	// Returning zero as we've manually set the program counter.
	return 0
}

// rst is the "Restart" handler.
//
// The contents of the program counter are pushed onto the stack, providing a
// return address for later use by a RETURN instruction.
//
// The program execution continues at an address indicated by opc.
func (i *Intel8080) rst(opc byte) opHandler {
	return func() uint16 {
		i.stackAdd(i.pc)

		i.pc = uint16(opc) & 0x38

		// Returning zero as we've manually set the program counter.
		return 0
	}
}
