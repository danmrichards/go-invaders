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

// dad is the "Double Add" handler.
//
// The 16-bit number in the specified register pair is added to the 16-bit
// number held in the H and L registers using two's complement arithmetic. The
// result replaces the contents of the H and L registers.
func (i *Intel8080) dad(x, y byte) opHandler {
	return func() uint16 {
		// Get the full number held across the register pair.
		n := uint16(x)<<8 | uint16(y)

		// Get the full number held in the HL register pair.
		hl := uint16(i.h)<<8 | uint16(i.l)

		ans := hl + n

		// Set the carry condition bit accordingly if the result of the
		// arithmetic was greater than 0xff (11111111 in base 2 and 255 in base
		// 10).
		i.cc.cy = ans > 0xff-n

		// Store the answer back on to the HL register pair.
		i.h = uint8(ans >> 8)
		i.h = uint8(ans)

		return defaultInstructionLen
	}
}

// dad is the "Double Add Stack Pointer" handler.
//
// The 16-bit number in the stack pointer is added to the 16-bit number held in
// the H and L registers using two's complement arithmetic. The result replaces
// the contents of the H and L registers.
func (i *Intel8080) dadSP() uint16 {
	// Get the full number held in the HL register pair.
	hl := uint16(i.h)<<8 | uint16(i.l)

	ans := hl + i.sp

	// Set the carry condition bit accordingly if the result of the
	// arithmetic was greater than 0xff (11111111 in base 2 and 255 in base
	// 10).
	i.cc.cy = ans > 0xff-i.sp

	// Store the number back on to the HL register pair.
	i.h = uint8(ans >> 8)
	i.h = uint8(ans)

	return defaultInstructionLen
}

// dcx is the "Decrement Register Pair" handler.
//
// The number held in the given register pair is decremented by one.
func (i *Intel8080) dcx(x, y *byte) opHandler {
	return func() uint16 {
		// Get the full number held across the register pair.
		v := uint16(*x)<<8 | uint16(*y)
		v--

		// Split the number back up.
		*x = uint8(v >> 8)
		*y = uint8(v)

		return defaultInstructionLen
	}
}

// dcxSP is the "Decrement Stack Pointer" handler.
func (i *Intel8080) dcxSP() uint16 {
	i.sp--
	return defaultInstructionLen
}

// daa is the "Decimal Adjust Accumulator" handler.
//
// The eight-bit hexadecimal number in the accumulator is adjusted to form two
// four-bit binary coded decimal digits.
func (i *Intel8080) daa() uint16 {
	var a uint8

	lsb := i.a & 0x0f
	msb := i.a >> 4

	// If the least significant four bits of the accumulator represents a number
	// greater than 9, or if the Auxiliary Carry bit is equal to one, the
	// accumulator is incremented by six. Otherwise, no incrementing occurs.
	if lsb > 9 || i.cc.ac {
		a += 0x06
	}

	// If the most significant four bits of the accumulator now represent a
	// number greater than 9, or if the normal carry bit is equal to one, the
	// most significant four bits of the accumulator are incremented by six.
	if msb > 9 || i.cc.cy || (msb >= 9 && lsb > 9) {
		a += 0x60
	}

	i.accumulatorAdd(a)
	i.cc.cy = true

	return defaultInstructionLen
}

// adc is the "Add Register to Accumulator With Carry" handler.
//
// The specified byte plus the content of the Carry bit is added to the contents
// of the accumulator.
func (i *Intel8080) adc(b byte) opHandler {
	return func() uint16 {
		// Increment the byte if the carry bit is set.
		if i.cc.cy {
			b++
		}

		i.accumulatorAdd(b)
		return defaultInstructionLen
	}
}

// adcM is the "Add Memory to Accumulator With Carry" handler.
//
// The specified byte plus the content of the Carry bit is added to the contents
// of the accumulator.
//
// The byte pointed to by the HL register pair, plus the content of the Carry
// bit, is added to the contents of the accumulator and relevant condition bits
// are set.
func (i *Intel8080) adcM() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)
	n := i.mem.Read(addr)

	// Increment the byte if the carry bit is set.
	if i.cc.cy {
		n++
	}

	i.accumulatorAdd(n)
	return defaultInstructionLen
}

// sub is the "Subtract Register from Accumulator" handler.
//
// The given byte is subtracted from the contents of the accumulator and
// relevant condition bits are set.
func (i *Intel8080) sub(b byte) opHandler {
	return func() uint16 {
		i.accumulatorSub(b)
		return defaultInstructionLen
	}
}

// subM is the "Subtract Memory from Accumulator" handler.
//
// The byte pointed to by the HL register pair is subtracted from the contents
// of the accumulator and relevant condition bits are set.
func (i *Intel8080) subM() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	i.accumulatorSub(i.mem.Read(addr))
	return defaultInstructionLen
}

// sbb is the "Subtract Register from Accumulator With Borrow" handler.
//
// The Carry bit is internally added to the contents of the specified byte. This
// value is then subtracted from the accumulator using two's complement
// arithmetic.
func (i *Intel8080) sbb(b byte) opHandler {
	return func() uint16 {
		// Increment the byte if the carry bit is set.
		if i.cc.cy {
			b++
		}

		i.accumulatorSub(b)
		return defaultInstructionLen
	}
}

// sbbM is the "Subtract Memory from Accumulator With Borrow" handler.
//
// The Carry bit is internally added to the contents of the byte pointed to by
// the HL register pair. This value is then subtracted from the accumulator
// using two's complement arithmetic.
func (i *Intel8080) sbbM() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	n := i.mem.Read(addr)

	// Increment the byte if the carry bit is set.
	if i.cc.cy {
		n++
	}

	i.accumulatorSub(n)
	return defaultInstructionLen
}

// aci is the "Add Immediate to Accumulator With Carry" handler.
//
// The next byte of data from memory, plus the contents of the Carry bit, is
// added to the contents of the accumulator and relevant condition bits are set.
func (i *Intel8080) aci() uint16 {
	n := i.mem.Read(i.pc + 1)
	// Increment the byte if the carry bit is set.
	if i.cc.cy {
		n++
	}

	i.accumulatorAdd(n)
	return 2
}

// sui is the "Subtract Immediate from Accumulator" handler.
//
// The next byte of data from memory is subtracted from the contents of the
// accumulator and relevant condition bits are set.
func (i *Intel8080) sui() uint16 {
	i.accumulatorSub(i.mem.Read(i.pc + 1))
	return 2
}

// sbi is the "Subtract Immediate from Accumulator With Borrow" handler.
//
// The Carry bit is internally added to the byte of immediate data. This value
// is then subtracted from the accumulator using two'scomplement arithmetic.
func (i *Intel8080) sbi() uint16 {
	n := i.mem.Read(i.pc + 1)

	// Increment the byte if the carry bit is set.
	if i.cc.cy {
		n++
	}

	i.accumulatorSub(n)
	return 2
}
