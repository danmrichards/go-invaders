package cpu

// add is the "Add Register to Accumulator" handler.
//
// The given byte is added to the contents of the accumulator and relevant
// condition bits are set.
func (i *Intel8080) add(b *byte) opHandler {
	return func() {
		i.accumulatorAdd(*b)
	}
}

// adi is the "Add Immediate to Accumulator" handler.
//
// The next byte of data from memory is added to the contents of the accumulator
// and relevant condition bits are set.
func (i *Intel8080) adi() {
	i.accumulatorAdd(i.immediateByte())
}

// addM is the "Add Memory to Accumulator" handler.
//
// The byte pointed to by the HL register pair is added to the contents of the
// accumulator and relevant condition bits are set.
func (i *Intel8080) addM() {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	i.accumulatorAdd(i.mem.Read(addr))
}

// inx is the "Increment Register Pair" handler.
//
// The number held in the given register pair is incremented by one.
func (i *Intel8080) inx(x, y *byte) opHandler {
	return func() {
		// Get the full number held across the register pair.
		v := uint16(*x)<<8 | uint16(*y)
		v++

		// Split the number back up.
		*x = uint8(v >> 8)
		*y = uint8(v & 0x00ff)
	}
}

// inxSP is the "Increment Stack Pointer" handler.
func (i *Intel8080) inxSP() {
	i.sp++
}

// inr is the "Increment Register" handler.
//
// The specified register is incremented by one.
func (i *Intel8080) inr(r *byte) opHandler {
	return func() {
		// Perform the arithmetic at higher precision in order to capture the
		// carry out.
		ans := uint16(*r) + 1

		// Set the zero condition bit accordingly based on if the result of the
		// arithmetic was zero.
		i.cc.z = uint8(ans) == 0x00

		// Set the sign condition bit accordingly based on if the most
		// significant bit on the result of the arithmetic was set.
		//
		// Determine the result being zero with a bitwise AND operation against
		// 0x80 (10000000 in base 2 and 128 in base 10).
		//
		// 10000000 & 10000000 = 1
		i.cc.s = ans&0x80 == 0x80

		// Set the auxiliary carry condition bit accordingly if the result of
		// the arithmetic has a carry on the third bit.
		i.cc.ac = (*r&0x0f)+0x01 > 0x0f

		// Set the parity bit.
		i.cc.setParity(uint8(ans))

		*r = uint8(ans)
	}
}

// inrM is the "Increment Memory" handler.
//
// The byte pointed to by the HL register pair is incremented by one and
// relevant condition bits are set.
func (i *Intel8080) inrM() {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)
	n := i.mem.Read(addr)

	// Perform the arithmetic at higher precision in order to capture the
	// carry out.
	ans := uint16(n) + 1

	// Set the zero condition bit accordingly based on if the result of the
	// arithmetic was zero.
	i.cc.z = uint8(ans) == 0x00

	// Set the sign condition bit accordingly based on if the most
	// significant bit on the result of the arithmetic was set.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0x80 (10000000 in base 2 and 128 in base 10).
	//
	// 10000000 & 10000000 = 1
	i.cc.s = ans&0x80 == 0x80

	// Set the auxiliary carry condition bit accordingly if the result of
	// the arithmetic has a carry on the third bit.
	i.cc.ac = (n&0x0f)+0x01 > 0x0f

	// Set the parity bit.
	i.cc.setParity(uint8(ans))

	i.mem.Write(addr, uint8(ans))
}

// dcr is the "Decrement Register" handler.
//
// The specified register is decremented by one.
func (i *Intel8080) dcr(r *byte) opHandler {
	return func() {
		// Perform the arithmetic at higher precision in order to capture the
		// carry out.
		ans := uint16(*r) - 1

		// Set the zero condition bit accordingly based on if the result of the
		// arithmetic was zero.
		i.cc.z = uint8(ans) == 0x00

		// Set the sign condition bit accordingly based on if the most
		// significant bit on the result of the arithmetic was set.
		//
		// Determine the result being zero with a bitwise AND operation against
		// 0x80 (10000000 in base 2 and 128 in base 10).
		//
		// 10000000 & 10000000 = 1
		i.cc.s = ans&0x80 == 0x80

		// Set the auxiliary carry condition bit accordingly if the result of
		// the arithmetic has a carry on the third bit.
		i.cc.ac = (ans & 0x0f) != 0x0f

		// Set the parity bit.
		i.cc.setParity(uint8(ans))

		*r = uint8(ans)
	}
}

// dcrM is the "Decrement Memory" handler.
//
// The specified register is decremented by one.
func (i *Intel8080) dcrM() {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)
	n := i.mem.Read(addr)

	// Perform the arithmetic at higher precision in order to capture the
	// carry out.
	ans := uint16(n) - 1

	// Set the zero condition bit accordingly based on if the result of the
	// arithmetic was zero.
	i.cc.z = uint8(ans) == 0x00

	// Set the sign condition bit accordingly based on if the most
	// significant bit on the result of the arithmetic was set.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0x80 (10000000 in base 2 and 128 in base 10).
	//
	// 10000000 & 10000000 = 1
	i.cc.s = ans&0x80 == 0x80

	// Set the auxiliary carry condition bit accordingly if the result of
	// the arithmetic has a carry on the third bit.
	i.cc.ac = (ans & 0x0f) != 0x0f

	// Set the parity bit.
	i.cc.setParity(uint8(ans))

	i.mem.Write(addr, uint8(ans))
}

// dad is the "Double Add" handler.
//
// The 16-bit number in the specified register pair is added to the 16-bit
// number held in the H and L registers using two's complement arithmetic. The
// result replaces the contents of the H and L registers.
func (i *Intel8080) dad(x, y *byte) opHandler {
	return func() {
		// Get the full number held across the register pair.
		n := uint16(*x)<<8 | uint16(*y)

		// Get the full number held in the HL register pair.
		hl := uint16(i.h)<<8 | uint16(i.l)

		ans := hl + n

		// Set the carry condition bit accordingly.
		i.cc.cy = hl > 0xffff-n

		// Store the answer back on to the HL register pair.
		i.h = uint8(ans >> 8)
		i.l = uint8(ans)
	}
}

// dad is the "Double Add Stack Pointer" handler.
//
// The 16-bit number in the stack pointer is added to the 16-bit number held in
// the H and L registers using two's complement arithmetic. The result replaces
// the contents of the H and L registers.
func (i *Intel8080) dadSP() {
	// Get the full number held in the HL register pair.
	hl := uint16(i.h)<<8 | uint16(i.l)

	ans := hl + i.sp

	// Set the carry condition bit accordingly.
	i.cc.cy = hl > 0xffff-i.sp

	// Store the number back on to the HL register pair.
	i.h = uint8(ans >> 8)
	i.l = uint8(ans)
}

// dcx is the "Decrement Register Pair" handler.
//
// The number held in the given register pair is decremented by one.
func (i *Intel8080) dcx(x, y *byte) opHandler {
	return func() {
		// Get the full number held across the register pair.
		v := uint16(*x)<<8 | uint16(*y)
		v--

		// Split the number back up.
		*x = uint8(v >> 8)
		*y = uint8(v & 0x00ff)
	}
}

// dcxSP is the "Decrement Stack Pointer" handler.
func (i *Intel8080) dcxSP() {
	i.sp--
}

// daa is the "Decimal Adjust Accumulator" handler.
//
// The eight-bit hexadecimal number in the accumulator is adjusted to form two
// four-bit binary coded decimal digits.
func (i *Intel8080) daa() {
	var (
		a uint8
		c = i.cc.cy
	)

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
		c = true
	}

	i.accumulatorAdd(a)
	i.cc.cy = c
}

// adc is the "Add Register to Accumulator With Carry" handler.
//
// The specified byte plus the content of the Carry bit is added to the contents
// of the accumulator.
func (i *Intel8080) adc(b *byte) opHandler {
	return func() {
		n := *b

		// Increment the byte if the carry bit is set.
		if i.cc.cy {
			n++
		}

		i.accumulatorAdd(n)
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
func (i *Intel8080) adcM() {
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
}

// sub is the "Subtract Register from Accumulator" handler.
//
// The given byte is subtracted from the contents of the accumulator and
// relevant condition bits are set.
func (i *Intel8080) sub(b *byte) opHandler {
	return func() {
		i.accumulatorSub(*b)
	}
}

// subM is the "Subtract Memory from Accumulator" handler.
//
// The byte pointed to by the HL register pair is subtracted from the contents
// of the accumulator and relevant condition bits are set.
func (i *Intel8080) subM() {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	i.accumulatorSub(i.mem.Read(addr))
}

// sbb is the "Subtract Register from Accumulator With Borrow" handler.
//
// The Carry bit is internally added to the contents of the specified byte. This
// value is then subtracted from the accumulator using two's complement
// arithmetic.
func (i *Intel8080) sbb(b *byte) opHandler {
	return func() {
		n := *b

		// Increment the byte if the carry bit is set.
		if i.cc.cy {
			n++
		}

		i.accumulatorSub(n)
	}
}

// sbbM is the "Subtract Memory from Accumulator With Borrow" handler.
//
// The Carry bit is internally added to the contents of the byte pointed to by
// the HL register pair. This value is then subtracted from the accumulator
// using two's complement arithmetic.
func (i *Intel8080) sbbM() {
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
}

// aci is the "Add Immediate to Accumulator With Carry" handler.
//
// The next byte of data from memory, plus the contents of the Carry bit, is
// added to the contents of the accumulator and relevant condition bits are set.
func (i *Intel8080) aci() {
	n := i.immediateByte()

	c := uint8(0)
	if i.cc.cy {
		c++
	}

	// Perform the arithmetic at higher precision in order to capture the
	// carry out.
	ans := uint16(i.a) + uint16(n) + uint16(c)
	r := byte(ans & 0xff)

	// Set the zero condition bit accordingly based on if the result of the
	// arithmetic was zero.
	i.cc.z = r == 0x00

	// Set the sign condition bit accordingly based on if the most
	// significant bit on the result of the arithmetic was set.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0x80 (10000000 in base 2 and 128 in base 10).
	//
	// 10000000 & 10000000 = 1
	i.cc.s = r&0x80 == 0x80

	// Set the carry condition bit accordingly if the result of the
	// arithmetic was greater than 0xff (11111111 in base 2 and 255 in base 10).
	i.cc.cy = ans > 0xff

	// Set the auxiliary carry condition bit accordingly if the result of
	// the arithmetic has a carry on the third bit.
	i.cc.ac = (i.a&0x0f)+(n&0x0f)+c > 0x0f

	// Set the parity bit.
	i.cc.setParity(r)

	// Finally update the accumulator.
	i.a = r
}

// sui is the "Subtract Immediate from Accumulator" handler.
//
// The next byte of data from memory is subtracted from the contents of the
// accumulator and relevant condition bits are set.
func (i *Intel8080) sui() {
	i.accumulatorSub(i.immediateByte())
}

// sbi is the "Subtract Immediate from Accumulator With Borrow" handler.
//
// The Carry bit is internally added to the byte of immediate data. This value
// is then subtracted from the accumulator using two'scomplement arithmetic.
func (i *Intel8080) sbi() {
	n := i.immediateByte()

	c := uint8(0)
	if i.cc.cy {
		c++
	}

	// Perform the arithmetic at higher precision in order to capture the
	// carry out.
	ans := uint16(i.a) - uint16(n) - uint16(c)
	r := byte(ans & 0xff)

	// Set the zero condition bit accordingly based on if the result of the
	// arithmetic was zero.
	i.cc.z = r == 0x00

	// Set the sign condition bit accordingly based on if the most
	// significant bit on the result of the arithmetic was set.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0x80 (10000000 in base 2 and 128 in base 10).
	//
	// 10000000 & 10000000 = 1
	i.cc.s = r&0x80 == 0x80

	// Set the carry condition bit accordingly if the result of the
	// arithmetic was greater than 0xff (11111111 in base 2 and 255 in base
	// 10).
	i.cc.cy = uint16(i.a) < uint16(n)+uint16(c)

	// Set the auxiliary carry condition bit accordingly if the result of
	// the arithmetic has a carry on the third bit.
	i.cc.ac = (i.a&0x0f)-(n&0x0f)-c >= 0x00

	// Set the parity bit.
	i.cc.setParity(r)

	// Finally update the accumulator.
	i.a = r
}
