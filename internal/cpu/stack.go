package cpu

// pop is the "Pop Data Off Stack" handler.
//
// The contents of the specified register pair are restored from two bytes of
// memory indicated by the stack pointer SP.
func (i *Intel8080) pop(x, y *byte) opHandler {
	return func() uint16 {
		*x = i.mem.Read(i.sp)
		*y = i.mem.Read(i.sp + 1)
		i.sp += 2

		return defaultInstructionLen
	}
}

// push is the "Push Data Onto Stack" handler.
//
// The contents of the specified register pair are saved in two bytes of memory
// indicated by the stack pointer SP.
func (i *Intel8080) push(x, y byte) opHandler {
	return func() uint16 {
		i.mem.Write(i.sp-1, x)
		i.mem.Write(i.sp-2, y)
		i.sp -= 2

		return defaultInstructionLen
	}
}

// popPSW is the "Pop Data Off Stack PSW" handler.
//
// The contents of the PSW register pair are restored from two bytes of
// memory indicated by the stack pointer SP.
func (i *Intel8080) popPSW() uint16 {
	i.a = i.mem.Read(i.sp + 1)
	psw := i.mem.Read(i.sp)

	i.cc.z = (psw & 0x01) == 0x01
	i.cc.s = (psw & 0x02) == 0x02
	i.cc.p = (psw & 0x04) == 0x04
	i.cc.cy = (psw & 0x08) == 0x08
	i.cc.ac = (psw & 0x10) == 0x08

	i.sp += 2

	return defaultInstructionLen
}

// pushPSW is the "Push Data Onto Stack PSW" handler.
//
// The contents of the PSW register pair are saved in two bytes of memory
// indicated by the stack pointer SP.
func (i *Intel8080) pushPSW() uint16 {
	i.mem.Write(i.sp-1, i.a)

	var z, s, p, cy, ac byte
	if i.cc.z {
		z = 1
	}
	if i.cc.s {
		s = 1
	}
	if i.cc.p {
		p = 1
	}
	if i.cc.cy {
		cy = 1
	}
	if i.cc.ac {
		ac = 1
	}

	psw := z | s<<1 | p<<2 | cy<<3 | ac<<4

	i.mem.Write(i.sp-2, psw)

	i.sp -= 2
	return defaultInstructionLen
}

// xthl is the "Exchange Stack" handler.
//
// The contents of the L register are exchanged with the contents of the memory
// byte whose address is held in the stack pointer SP. The contents of the H
// register are exchanged with the contents of the memory byte whose address is
// one greater than that held in the stack pointer.
func (i *Intel8080) xthl() uint16 {
	b := uint16(i.mem.Read(i.sp+1))<<8 | uint16(i.mem.Read(i.sp))
	i.sp += 2

	hl := uint16(i.h)<<8 | uint16(i.l)

	i.h = uint8(b >> 8)
	i.l = uint8(b)

	i.push(byte(hl&0xff), byte(hl>>8))

	return defaultInstructionLen
}

// sphl is the "Load SP from H and L" handler.
//
// The 16 bits of data held in the H and L registers replace the contents of the
// stack pointer SP. The contents of the H and L registers are unchanged.
func (i *Intel8080) sphl() uint16 {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	i.sp = addr

	return defaultInstructionLen
}
