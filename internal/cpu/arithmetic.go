package cpu

// add is the "Add Register to Accumulator" handler.
//
// The given byte is added to the contents of the accumulator and relevant
// condition bits are set.
func (i *Intel8080) add(b byte) opHandler {
	return func() uint16 {
		i.accumulatorAdd(b)
		return defaultInstructionLen
	}
}

// adi is the "Add Immediate to Accumulator" handler.
//
// The next byte of data from memory is added to the contents of the accumulator
// and relevant condition bits are set.
func (i *Intel8080) adi() uint16 {
	i.accumulatorAdd(i.mem.Read(i.pc + 1))
	return 2
}

// addM is the "Add Memory to Accumulator" handler.
//
// The byte pointed to by the HL register pair is added to the contents of the
// accumulator and relevant condition bits are set.
func (i *Intel8080) addM() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	i.accumulatorAdd(i.mem.Read(addr))
	return defaultInstructionLen
}

// inx is the "Increment Register Pair" handler.
//
// The number held in the given register pair is incremented by one.
func (i *Intel8080) inx(x, y *byte) opHandler {
	return func() uint16 {
		// Get the full number held across the register pair.
		v := uint16(*x)<<8 | uint16(*y)
		v++

		// Split the number back up.
		*x = uint8(v >> 8)
		*y = uint8(v)

		return defaultInstructionLen
	}
}

// inxSP is the "Increment Stack Pointer" handler.
func (i *Intel8080) inxSP() uint16 {
	i.sp++
	return defaultInstructionLen
}

// inr is the "Increment Register" handler.
//
// The specified register is incremented by one.
func (i *Intel8080) inr(r *byte) opHandler {
	return func() uint16 {
		// Perform the arithmetic at higher precision in order to capture the
		// carry out.
		ans := uint16(*r) + 1

		// Set the zero condition bit accordingly based on if the result of the
		// arithmetic was zero.
		//
		// Determine the result being zero with a bitwise AND operation against
		// 0xff (11111111 in base 2 and 255 in base 10).
		//
		// 00000000 & 11111111 = 0
		i.cc.z = ans&0xff == 0

		// Set the sign condition bit accordingly based on if the most
		// significant bit on the result of the arithmetic was set.
		//
		// Determine the result being zero with a bitwise AND operation against
		// 0x80 (10000000 in base 2 and 128 in base 10).
		//
		// 10000000 & 10000000 = 1
		i.cc.s = ans&0x80 == 1

		// Set the auxiliary carry condition bit accordingly if the result of
		// the arithmetic has a carry on the third bit.
		i.cc.ac = (ans & 0x0f) == 0x00

		// Set the parity bit.
		i.cc.setParity(uint8(ans))

		*r = uint8(ans)
		return defaultInstructionLen
	}
}

// inrM is the "Increment Memory" handler.
//
// The byte pointed to by the HL register pair is incremented by one and
// relevant condition bits are set.
func (i *Intel8080) inrM() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	// Perform the arithmetic at higher precision in order to capture the
	// carry out.
	ans := i.mem.Read(addr) + 1

	// Set the zero condition bit accordingly based on if the result of the
	// arithmetic was zero.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0xff (11111111 in base 2 and 255 in base 10).
	//
	// 00000000 & 11111111 = 0
	i.cc.z = ans&0xff == 0

	// Set the sign condition bit accordingly based on if the most
	// significant bit on the result of the arithmetic was set.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0x80 (10000000 in base 2 and 128 in base 10).
	//
	// 10000000 & 10000000 = 1
	i.cc.s = ans&0x80 == 1

	// Set the auxiliary carry condition bit accordingly if the result of
	// the arithmetic has a carry on the third bit.
	i.cc.ac = (ans & 0x0f) == 0x00

	// Set the parity bit.
	i.cc.setParity(ans)

	i.mem.Write(addr, ans)
	return defaultInstructionLen
}

// dcr is the "Decrement Register" handler.
//
// The specified register is decremented by one.
func (i *Intel8080) dcr(r *byte) opHandler {
	return func() uint16 {
		// Perform the arithmetic at higher precision in order to capture the
		// carry out.
		ans := uint16(*r) - 1

		// Set the zero condition bit accordingly based on if the result of the
		// arithmetic was zero.
		//
		// Determine the result being zero with a bitwise AND operation against
		// 0xff (11111111 in base 2 and 255 in base 10).
		//
		// 00000000 & 11111111 = 0
		i.cc.z = ans&0xff == 0

		// Set the sign condition bit accordingly based on if the most
		// significant bit on the result of the arithmetic was set.
		//
		// Determine the result being zero with a bitwise AND operation against
		// 0x80 (10000000 in base 2 and 128 in base 10).
		//
		// 10000000 & 10000000 = 1
		i.cc.s = ans&0x80 == 1

		// Set the auxiliary carry condition bit accordingly if the result of
		// the arithmetic has a carry on the third bit.
		i.cc.ac = !((ans & 0x0F) == 0x0F)

		// Set the parity bit.
		i.cc.setParity(uint8(ans))

		*r = uint8(ans)
		return defaultInstructionLen
	}
}

// dcrM is the "Decrement Memory" handler.
//
// The specified register is decremented by one.
func (i *Intel8080) dcrM() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	// Perform the arithmetic at higher precision in order to capture the
	// carry out.
	ans := i.mem.Read(addr) - 1

	// Set the zero condition bit accordingly based on if the result of the
	// arithmetic was zero.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0xff (11111111 in base 2 and 255 in base 10).
	//
	// 00000000 & 11111111 = 0
	i.cc.z = ans&0xff == 0

	// Set the sign condition bit accordingly based on if the most
	// significant bit on the result of the arithmetic was set.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0x80 (10000000 in base 2 and 128 in base 10).
	//
	// 10000000 & 10000000 = 1
	i.cc.s = ans&0x80 == 1

	// Set the auxiliary carry condition bit accordingly if the result of
	// the arithmetic has a carry on the third bit.
	i.cc.ac = !((ans & 0x0F) == 0x0F)

	// Set the parity bit.
	i.cc.setParity(ans)

	i.mem.Write(addr, ans)
	return defaultInstructionLen
}
