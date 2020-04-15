package cpu

// lxi is the "Load Immediate Register" handler.
//
// This handler operates on a CPU register pair.
//
// Because the 8080 works on little-endian byte order, the first register in the
// pair stores the 8 most significant bits of an address while the second
// register stores the 8 least significant bits.
func (i *Intel8080) lxi(x, y *byte) opHandler {
	return func() uint16 {
		*x = i.mem.Read(i.pc + 2)
		*y = i.mem.Read(i.pc + 1)

		return 3
	}
}

// lxi is the "Load Immediate Stack Pointer" handler.
func (i *Intel8080) lxiSP() uint16 {
	i.sp = i.twoByteRead()
	return 3
}

// movRR is the "Move Register to Register" handler.
//
// One byte of data is moved from the register specified by src (the source
// register) to the register specified by dst (the destination register).
func (i *Intel8080) movRR(dst, src *byte) opHandler {
	return func() uint16 {
		*dst = *src

		return defaultInstructionLen
	}
}

// mvi is the "Move Immediate Data" handler.
//
// The byte of immediate data is stored in the specified register.
func (i *Intel8080) mvi(dst *byte) opHandler {
	return func() uint16 {
		*dst = i.mem.Read(i.pc + 1)
		return 2
	}
}

// mvi is the "Move Immediate Data Memory" handler.
//
// The byte of immediate data is stored in the register specified by the byte
// pointed by the HL register pair.
func (i *Intel8080) mviM() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	i.mem.Write(addr, i.mem.Read(i.pc+1))
	return 2
}
