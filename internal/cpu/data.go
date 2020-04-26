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

// ldax is the "Load Accumulator" handler.
//
// The contents of the memory location addressed by registers B and C, or by
// registers D and E, replace the contents of the accumulator.
func (i *Intel8080) ldax(x, y byte) opHandler {
	return func() uint16 {
		// Determine the address of the byte pointed by the given register pair.
		// The address is two bytes long, so merge the two bytes stored in each
		// side of the register pair.
		addr := uint16(x)<<8 | uint16(y)

		i.a = i.mem.Read(addr)
		return defaultInstructionLen
	}
}

// movRR is the "Move Register to Memory" handler.
//
// One byte of data is moved from the register specified by r (the source
// register) to the memory address pointed by the HL register pair.
func (i *Intel8080) movRM(r byte) opHandler {
	return func() uint16 {
		// Determine the address of the byte pointed by the HL register pair.
		// The address is two bytes long, so merge the two bytes stored in each
		// side of the register pair.
		addr := uint16(i.h)<<8 | uint16(i.l)

		i.mem.Write(addr, r)
		return defaultInstructionLen
	}
}

// sta is the "Store Accumulator Direct" handler.
//
// The contents of the accumulator replace the byte at the memory address formed
// by concatenating HI ADD with LOW ADD.
func (i *Intel8080) sta() uint16 {
	i.mem.Write(i.twoByteRead(), i.a)
	return 3
}

// movRR is the "Move Memory to Register" handler.
//
// One byte of data is moved from the memory address pointed by the HL register
// pair, to the given register r.
func (i *Intel8080) movMR(r *byte) opHandler {
	return func() uint16 {
		// Determine the address of the byte pointed by the HL register pair.
		// The address is two bytes long, so merge the two bytes stored in each
		// side of the register pair.
		addr := uint16(i.h)<<8 | uint16(i.l)

		*r = i.mem.Read(addr)
		return defaultInstructionLen
	}
}

// lda is the "Load Accumulator Direct" handler.
//
// The byte at the memory address formed by concatenating HI ADD with LOW ADD
// replaces the contents of the accumulator.
func (i *Intel8080) lda() uint16 {
	i.a = i.mem.Read(i.twoByteRead())
	return 2
}

// stax is the "Store Accumulator" handler.
//
// The contents of the accumulator are stored in the memory location addressed
// by registers B an dC, or by registers 0 and E.
func (i *Intel8080) stax(x, y byte) opHandler {
	return func() uint16 {
		// Determine the address of the byte pointed by the given register pair.
		// The address is two bytes long, so merge the two bytes stored in each
		// side of the register pair.
		addr := uint16(x)<<8 | uint16(y)

		i.mem.Write(addr, i.a)
		return defaultInstructionLen
	}
}

// shld is the "Store H and L Direct" handler.
//
// The contents of the L register are stored at the memory address formed by
// concatenating HI ADD with LOW ADD. The contents of the H register are stored
// at the next higher memory address.
func (i *Intel8080) shld() uint16 {
	addr := i.twoByteRead()

	hl := uint16(i.h)<<8 | uint16(i.l)

	i.mem.Write(addr, byte(hl&0xff))
	i.mem.Write(addr+1, byte(hl>>8))

	return 3
}

// lhld is the "Load H and L Direct" handler.
//
// The byte at the memory address formed by concatenating HI ADD with LOW ADD
// replaces the contents of the L register. The byte at the next higher memory
// address replaces the contents of the H register.
func (i *Intel8080) lhld() uint16 {
	addr := i.twoByteRead()
	b := uint16(i.mem.Read(addr+1))<<8 | uint16(i.mem.Read(addr))

	i.h = uint8(b >> 8)
	i.l = uint8(b)

	return 3
}

// xchg is the "Exchange Registers" handler.
//
// The 16 bits of data held in the H and L registers are exchanged with the 16
// bits of data held in the D and E registers.
func (i *Intel8080) xchg() uint16 {
	// Swap H <-> D.
	d := i.d
	h := i.h
	i.h = d
	i.d = h

	// Swap L <-> E.
	l := i.l
	e := i.e
	i.l = e
	i.e = l

	return defaultInstructionLen
}
