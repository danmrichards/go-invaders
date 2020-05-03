package cpu

// jmp is the "Jump" handler.
//
// This handler jumps the program counter to a given point in memory.
func (i *Intel8080) jmp() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.immediateWord()
}

// call is the "Call subroutine" handler.
//
// A call operation is unconditionally performed to subroutine sub.
func (i *Intel8080) call() {
	// We will dump the program to the subroutine indicated by the two immediate
	// bytes in memory.
	addr := i.immediateWord()

	// Update the stack pointer. Note we're incrementing by 3 to account for
	// this operation and the two bytes we've read.
	i.stackAdd(i.pc)

	i.pc = addr
}

// ret is the "Return" handler.
//
// A return operation is unconditionally performed.
func (i *Intel8080) ret() {
	i.pc = i.stackPop()
}

// jnz is the "Jump If Not Zero" handler.
//
// If the zero bit is one, program execution continues at the memory address adr.
func (i *Intel8080) jnz() {
	if !i.cc.z {
		i.jmp()
	}
}

// jz is the "Jump Zero" handler.
//
// If the zero bit is not one, program execution continues at the memory address
// adr.
func (i *Intel8080) jz() {
	if i.cc.z {
		i.jmp()
	}
}

// jnc is the "Jump Not Carry" handler.
//
// If the carry bit is one, program execution continues at the memory address
// adr.
func (i *Intel8080) jnc() {
	if !i.cc.cy {
		i.jmp()
	}
}

// jc is the "Jump Carry" handler.
//
// If the carry bit is not one, program execution continues at the memory
// address adr.
func (i *Intel8080) jc() {
	// Return early if the carry bit isn't set.
	if i.cc.cy {
		i.jmp()
	}
}

// jpo is the "Jump If Parity Odd" handler.
//
// If the Parity bit is zero (indicating a result with odd parity), program
// execution continues at the memory address adr.
func (i *Intel8080) jpo() {
	if !i.cc.p {
		i.jmp()
	}
}

// jpe is the "Jump If Parity Even" handler.
//
// If the Parity bit is one (indicating a result with even parity), program
// execution continues at the memory address adr.
func (i *Intel8080) jpe() {
	// Return early if the parity bit isn't set.
	if i.cc.p {
		i.jmp()
	}
}

// jp is the "Jump If Positive" handler.
//
// If the Sign bit is zero (indicating a positive result), program execution
// continues at the memory address adr.
func (i *Intel8080) jp() {
	if !i.cc.s {
		i.jmp()
	}
}

// jm is the "Jump If Minus" handler.
//
// If the Sign bit is one (indicating a positive result), program execution
// continues at the memory address adr.
func (i *Intel8080) jm() {
	if i.cc.s {
		i.jmp()
	}
}

// cz is the "Call If Zero" handler.
//
// If the Zero bit is zero, a call operation is performed to subroutine sub.
func (i *Intel8080) cz() {
	if i.cc.z {
		i.call()
	}
}

// cnz is the "Call If Not Zero" handler.
//
// If the Zero bit is one, a call operation is performed to subroutine sub.
func (i *Intel8080) cnz() {
	if !i.cc.z {
		i.call()
	}
}

// cc is the "Call If Carry" handler.
//
// If the Carry bit is zero, a call operation is performed to subroutine sub.
func (i *Intel8080) cic() {
	if i.cc.cy {
		i.call()
	}
}

// cnc is the "Call If Not carry" handler.
//
// If the carry bit is one, a call operation is performed to subroutine sub.
func (i *Intel8080) cnc() {
	if !i.cc.cy {
		i.call()
	}
}

// cpo is the "Call If Parity Odd" handler.
//
// If the Parity bit is one (indicating a result with even parity), a call
// operation is performed to subroutine sub.
func (i *Intel8080) cpo() {
	if i.cc.p {
		i.call()
	}
}

// cpe is the "Call If Parity Even" handler.
//
// If the Parity bit is even (indicating a result with even parity), a call
// operation is performed to subroutine sub.
func (i *Intel8080) cpe() {
	if !i.cc.p {
		i.call()
	}
}

// cp is the "Call If Positive" handler.
//
// If the Sign bit is zero (indicating a positive result), a call operation is
// performed to subroutine sub.
func (i *Intel8080) cp() {
	if !i.cc.s {
		i.call()
	}
}

// cp is the "Call If Minus" handler.
//
// If the Sign bit is one (indicating a positive result), a call operation is
// performed to subroutine sub.
func (i *Intel8080) cm() {
	if i.cc.s {
		i.call()
	}
}

// rnz is the "Return If Not Zero" handler.
//
// If the Zero bit is zero, a return operation is performed.
func (i *Intel8080) rnz() {
	if !i.cc.z {
		i.ret()
	}
}

// rz is the "Return If Zero" handler.
//
// If the Zero bit is one, a return operation is performed.
func (i *Intel8080) rz() {
	if i.cc.z {
		i.ret()
	}
}

// rnc is the "Return If Not Carry" handler.
//
// If the Carry bit is zero, a return operation is performed.
func (i *Intel8080) rnc() {
	if !i.cc.cy {
		i.ret()
	}
}

// rc is the "Return If Carry" handler.
//
// If the Carry bit is one, a return operation is performed.
func (i *Intel8080) rc() {
	if i.cc.cy {
		i.ret()
	}
}

// rpo is the "Return If Parity Odd" handler.
//
// If the Parity bit is zero (indicating a result with odd parity), a return
// operation is performed.
func (i *Intel8080) rpo() {
	if !i.cc.p {
		i.ret()
	}
}

// rpe is the "Return If Parity Even" handler.
//
// If the Parity bit is one (indicating a result with event parity), a return
// operation is performed.
func (i *Intel8080) rpe() {
	if i.cc.p {
		i.ret()
	}
}

// rp is the "Return If Positive" handler.
//
// If the Sign bit is zero (indicating a positive result), a return operation
// is performed.
func (i *Intel8080) rp() {
	if !i.cc.s {
		i.ret()
	}
}

// rm is the "Return If Minus" handler.
//
// If the Sign bit is one (indicating a negative result), a return operation
// is performed.
func (i *Intel8080) rm() {
	if i.cc.s {
		i.ret()
	}
}

// pchl is the "Load Program Counter" handler.
//
// The contents of the H register replaces the most significant 8 bits of the
// program counter, and the contents of the L register replace the least
// significant 8 bits of the program counter.
//
// This causes program execution to continue at the address contained in the H
// and L registers.
func (i *Intel8080) pchl() {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	i.pc = addr
}

// rst is the "Restart" handler.
//
// The contents of the program counter are pushed onto the stack, providing a
// return address for later use by a RETURN instruction.
//
// The program execution continues at an address indicated by opc.
func (i *Intel8080) rst(opc byte) opHandler {
	return func() {
		i.stackAdd(i.pc)

		i.pc = uint16(opc) & 0x38
	}
}
