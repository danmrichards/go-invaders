package cpu

// cma is the "Compliment Accumulator" handler.
//
// Each bit of the contents of the accumulator is complemented (producing the
// one's complement).
//
// E.g. 01010001 -> 10101110
func (i *Intel8080) cma() {
	i.a = ^i.a
}

// ani is the "And Immediate With Accumulator" handler.
//
// The byte of immediate data is logically ANDed with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) ani() {
	n := i.a & i.immediateByte()

	// Set the condition bits accordingly.
	i.cc.z = n == 0
	i.cc.s = n&0x80 == 0x80
	i.cc.setParity(n)
	i.cc.cy = false

	i.a = n
}

// rrc is the "Rotate Accumulator Right" handler.
//
// The carry bit is set equal to the low-order bit of the accumulator. The
// contents of the accumulator are rotated one bit position to the right, with
// the low-order bit being transferred to the high-order bit position of the
// accumulator.
func (i *Intel8080) rrc() {
	n := i.a

	i.a = (n & 1 << 7) | (n >> 1)
	i.cc.cy = n&1 == 1
}

// rar is the "Rotate Accumulator Right Through Carry" handler.
//
// The contents of the accumulator are rotated one bit position to the right.
//
// The low-order bit of the accumulator replaces the carry bit, while the carry
// bit replaces the high-order bit of the accumulator.
func (i *Intel8080) rar() {
	n := i.a

	var c byte
	if i.cc.cy {
		c = 1
	}
	i.a = (c << 7) | (n >> 1)
	i.cc.cy = n&1 == 1
}

// cpi is the "Compare Immediate With Accumulator" handler.
func (i *Intel8080) cpi() {
	b := i.immediateByte()
	n := i.a - b

	// Set the zero condition bit accordingly based on if the result of the
	// arithmetic was zero.
	i.cc.z = uint8(n) == 0x00

	// Set the sign condition bit accordingly based on if the most
	// significant bit on the result of the arithmetic was set.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0x80 (10000000 in base 2 and 128 in base 10).
	//
	// 10000000 & 10000000 = 1
	i.cc.s = n&0x80 == 0x80

	// Set the carry condition bit accordingly if the result of the
	// arithmetic was greater than 0xff (11111111 in base 2 and 255 in base
	// 10).
	i.cc.cy = i.a < n

	// Set the auxiliary carry condition bit accordingly if the result of
	// the arithmetic has a carry on the third bit.
	i.cc.ac = (i.a&0x0f)-(n&0x0f) >= 0x00

	// Set the parity bit.
	i.cc.setParity(uint8(n))
}

// xra is the "Logical Exclusive-Or Register With Accumulator" handler.
//
// The specified byte is EXCLUSIVE-ORed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) xra(b *byte) opHandler {
	return func() {
		n := *b ^ i.a

		i.cc.z = n == 0
		i.cc.s = n&0x80 == 0x80
		i.cc.setParity(n)
		i.cc.ac = false
		i.cc.cy = false

		i.a = n
	}
}

// xraM is the "Logical Exclusive-Or Memory With Accumulator" handler.
//
// The specified byte is EXCLUSIVE-ORed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) xraM() {
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
}

// ana is the "Logical AND Register With Accumulator" handler.
//
// The specified byte is logically ANDed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) ana(b *byte) opHandler {
	return func() {
		n := *b & i.a

		i.cc.z = n == 0
		i.cc.s = n&0x80 == 0x80
		i.cc.setParity(n)
		i.cc.ac = false
		i.cc.cy = false

		i.a = n
	}
}

// anaM is the "Logical AND Memory With Accumulator" handler.
//
// The specified byte is logically ANDed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) anaM() {
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
}

// ora is the "Logical OR Register With Accumulator" handler.
//
// The specified byte is logically ORed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) ora(b *byte) opHandler {
	return func() {
		n := *b | i.a

		i.cc.z = n == 0
		i.cc.s = n&0x80 == 0x80
		i.cc.setParity(n)
		i.cc.ac = false
		i.cc.cy = false

		i.a = n
	}
}

// oraM is the "Logical OR Memory With Accumulator" handler.
//
// The specified byte is logically ORed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) oraM() {
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
func (i *Intel8080) cmp(b *byte) opHandler {
	return func() {
		n := i.a - *b

		// Set the condition bits.
		i.cc.z = n == 0
		i.cc.s = n&0x80 == 0x80
		i.cc.setParity(n)
		i.cc.ac = false
		i.cc.cy = i.a < *b
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
func (i *Intel8080) cmpM() {
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
}

// rlc is the "Rotate Accumulator Left" handler.
//
// The Carry bit is set equal to the high-order bit of the accumulator. The
// contents of the accumulator are rotated one bit position to the left, with
// the high-order bit being transferred to the low-order bit position of the
// accumulator.
func (i *Intel8080) rlc() {
	i.cc.cy = (i.a & (1 << 7)) != 0

	var c byte
	if i.cc.cy {
		c = 1
	}
	i.a = i.a<<1 | c
}

// ral is the "Rotate Accumulator Left Through Carry" handler.
//
// The contents of the accumulator are rotated one bit position to the left.
//
// The high-order bit of the accumulator replaces the carry bit, while the carry
// bit replaces the high-order bit of the accumulator.
func (i *Intel8080) ral() {
	a := i.a
	cy := (a & 0x80) != 0

	i.a = a << 1

	if i.cc.cy {
		i.a |= 0x01
	}

	if cy {
		i.cc.cy = true
	} else {
		i.cc.cy = false
	}
}

// xri is the "Exclusive-Or Immediate With Accumulator" handler.
//
// The byte of immediate data is EXCLUSIVE-ORed with the contents of the
// accumulator. The carry bit is set to zero.
func (i *Intel8080) xri() {
	n := i.immediateByte() ^ i.a

	i.cc.z = n == 0
	i.cc.s = n&0x80 == 0x80
	i.cc.setParity(n)
	i.cc.ac = false
	i.cc.cy = false

	i.a = n
}

// ori is the "Logical OR Immediate Register With Accumulator" handler.
//
// The byte of immediate data is logically ORed bit by bit with the contents of
// the accumulator. The Carry bit is reset to zero.
func (i *Intel8080) ori() {
	n := i.immediateByte() | i.a

	i.cc.z = n == 0
	i.cc.s = n&0x80 == 0x80
	i.cc.setParity(n)
	i.cc.ac = false
	i.cc.cy = false

	i.a = n
}
