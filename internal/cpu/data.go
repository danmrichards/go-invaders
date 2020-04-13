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
