package cpu

// cma is the "Compliment Accumulator" handler.
//
// Each bit of the contents of the accumulator is complemented (producing the
// one's complement).
//
// E.g. 01010001 -> 10101110
func (i *Intel8080) cma() uint16 {
	i.a = ^i.a

	return defaultInstructionLen
}

// ani is the "And Immediate With Accumulator" handler.
//
// The byte of immediate data is logically ANDed with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) ani() uint16 {
	n := i.a & i.mem.Read(i.pc+1)

	// Set the condition bits accordingly.
	i.cc.z = n == 0
	i.cc.s = n&0x80 == 0x80
	i.cc.setParity(n)
	i.cc.cy = false

	i.a = n
	return 2
}

// rrc is the "Rotate Accumulator Right" handler.
//
// The carry bit is set equal to the low-order bit of the accumulator. The
// contents of the accumulator are rotated one bit position to the right, with
// the low-order bit being transferred to the high-order bit position of the
// accumulator.
func (i *Intel8080) rrc() uint16 {
	n := i.a

	i.a = (n & 1 << 7) | (n >> 1)
	i.cc.cy = n&1 == 1

	return defaultInstructionLen
}

// rar is the "Rotate Accumulator Right Through Carry" handler.
//
// The contents of the accumulator are rotated one bit position to the right.
//
// The low-order bit of the accumulator replaces the carry bit, while the carry
// bit replaces the high-order bit of the accumulator.
func (i *Intel8080) rar() uint16 {
	n := i.a

	var c byte
	if i.cc.cy {
		c = 1
	}
	i.a = (c << 7) | (n >> 1)
	i.cc.cy = n&1 == 1

	return defaultInstructionLen
}

// cpi is the "Compare Immediate With Accumulator" handler.
func (i *Intel8080) cpi() uint16 {
	b := i.mem.Read(i.pc + 1)
	n := i.a - b

	// Set the condition bits.
	i.cc.z = n == 0
	i.cc.s = n&0x80 == 0x80
	i.cc.setParity(n)
	i.cc.ac = false
	i.cc.cy = i.a < b

	return 2
}

// xra is the "Logical Exclusive-Or Register With Accumulator" handler.
//
// The specified byte is EXCLUSIVE-ORed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) xra(b byte) opHandler {
	return func() uint16 {
		n := b ^ i.a

		i.cc.z = n == 0
		i.cc.s = n&0x80 == 0x80
		i.cc.setParity(n)
		i.cc.ac = false
		i.cc.cy = false

		i.a = n
		return defaultInstructionLen
	}
}

// xraM is the "Logical Exclusive-Or Memory With Accumulator" handler.
//
// The specified byte is EXCLUSIVE-ORed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) xraM() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	n := i.mem.Read(addr) ^ i.a

	i.cc.z = n == 0
	i.cc.s = n&0x80 == 0x80
	i.cc.setParity(n)
	i.cc.ac = false
	i.cc.cy = false

	i.a = n
	return defaultInstructionLen
}

// ana is the "Logical AND Register With Accumulator" handler.
//
// The specified byte is logically ANDed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) ana(b byte) opHandler {
	return func() uint16 {
		n := b & i.a

		i.cc.z = n == 0
		i.cc.s = n&0x80 == 0x80
		i.cc.setParity(n)
		i.cc.ac = false
		i.cc.cy = false

		i.a = n
		return defaultInstructionLen
	}
}

// anaM is the "Logical AND Memory With Accumulator" handler.
//
// The specified byte is logically ANDed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) anaM() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	n := i.mem.Read(addr) & i.a

	i.cc.z = n == 0
	i.cc.s = n&0x80 == 0x80
	i.cc.setParity(n)
	i.cc.ac = false
	i.cc.cy = false

	i.a = n
	return defaultInstructionLen
}

// ora is the "Logical OR Register With Accumulator" handler.
//
// The specified byte is logically ORed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) ora(b byte) opHandler {
	return func() uint16 {
		n := b | i.a

		i.cc.z = n == 0
		i.cc.s = n&0x80 == 0x80
		i.cc.setParity(n)
		i.cc.ac = false
		i.cc.cy = false

		i.a = n
		return defaultInstructionLen
	}
}

// oraM is the "Logical OR Memory With Accumulator" handler.
//
// The specified byte is logically ORed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) oraM() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	n := i.mem.Read(addr) | i.a

	i.cc.z = n == 0
	i.cc.s = n&0x80 == 0x80
	i.cc.setParity(n)
	i.cc.ac = false
	i.cc.cy = false

	i.a = n
	return defaultInstructionLen
}

// cmp is the "Compare Register With Accumulator" handler.
//
// The specified byte is compared to the contents of the accumulator. The
// comparison is performed by internally subtracting the contents of REG from
// the accumulator (leaving both unchanged) and setting the condition bits
// according to the result.
//
// In particular, the Zero bit is set if the quantities are equal, and reset if
// they are unequal. Since a subtract operation is performed, the Carry bit will
// be set if there is no carry out of bit 7, indicating that the contents of REG
// are greater than the contents of the accumulator, and reset otherwise.
func (i *Intel8080) cmp(b byte) opHandler {
	return func() uint16 {
		n := i.a - b

		// Set the condition bits.
		i.cc.z = n == 0
		i.cc.s = n&0x80 == 0x80
		i.cc.setParity(n)
		i.cc.ac = false
		i.cc.cy = i.a < b

		return defaultInstructionLen
	}
}

// cmpM is the "Compare Memory With Accumulator" handler.
//
// The specified byte is compared to the contents of the accumulator. The
// comparison is performed by internally subtracting the contents of REG from
// the accumulator (leaving both unchanged) and setting the condition bits
// according to the result.
//
// In particular, the Zero bit is set if the quantities are equal, and reset if
// they are unequal. Since a subtract operation is performed, the Carry bit will
// be set if there is no carry out of bit 7, indicating that the contents of REG
// are greater than the contents of the accumulator, and reset otherwise.
func (i *Intel8080) cmpM() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	n := i.a - i.mem.Read(addr)

	// Set the condition bits.
	i.cc.z = n == 0
	i.cc.s = n&0x80 == 0x80
	i.cc.setParity(n)
	i.cc.ac = false
	i.cc.cy = i.a < n

	return defaultInstructionLen
}
