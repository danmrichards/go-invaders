package cpu

// conditions represents the condition bits of the Intel 8080.
//
// Condition bits are used to reflect the results of data operations, they can
// be effectively considered as flags.
type conditions struct {
	// Carry bit is set and reset by certain data operations, and its status can
	// be directly tested by a program. The operations which affect the Carry
	// bit are addition, subtraction, rotate, and logical operations.
	cy bool

	// Auxiliary Carry bit indicates carry out of bit 3. The state of the
	// Auxiliary Carry bit cannot be directly tested by a program instruction
	// and is present only to enable one instruction (DAA).
	ac bool

	// Sign bit is set at the conclusion of certain instructions, it will be set
	// to the condition of the most significant bit of the answer (bit 7).
	s bool

	// Zero bit is set if the result generated by the execution of certain
	// instructions is zero. The Zero bit is reset if the result is not zero.
	z bool

	// Parity bit is set to 1 for even parity, and is reset to 0 for odd parity.
	// Byte "parity" is checked after certain operations. The number of 1 bits
	// in a byte are counted, and if the total is odd, "odd" parity is flagged;
	// if the total is even, "even" parity is flagged.
	p bool
}

// setParity sets the parity bit based upon the number of set bits in byte b.
func (c *conditions) setParity(b byte) {
	var n int

	// Iterate through the bits in the given byte and count how many are set.
	for i := 0; i < 8; i++ {
		if (b>>i)&0x1 == 1 {
			n++
		}
	}

	// Set parity based on the number of set bits being even.
	c.p = n%2 == 0
}

// status returns a special byte which represents the current status of the
// conditions.
//
// Intended for use with the accumulator to form the "Program Status Word".
func (c *conditions) status() (s byte) {
	if c.s {
		s |= 1 << 7
	}
	if c.z {
		s |= 1 << 6
	}
	if c.ac {
		s |= 1 << 4
	}
	if c.p {
		s |= 1 << 2
	}
	s |= 1 << 1
	if c.cy {
		s |= 1
	}

	return s
}

// setStatus sets the value of the conditions based on the given special byte.
//
// Intended for use with the accumulator to form the "Program Status Word".
func (c *conditions) setStatus(b byte) {
	c.s = (b >> 7 & 0x01) == 0x01
	c.z = (b >> 6 & 0x01) == 0x01
	c.ac = (b >> 4 & 0x01) == 0x01
	c.p = (b >> 2 & 0x01) == 0x01
	c.cy = (b & 0x1) == 0x1
}

// carryByte returns a byte representation of the carry flag.
func (c *conditions) carryByte() byte {
	if c.cy {
		return 1
	}
	return 0
}
