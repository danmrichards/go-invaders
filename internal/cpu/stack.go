package cpu

// pop is the "Pop Data Off Stack" handler.
//
// The contents of the specified register pair are restored from two bytes of
// memory indicated by the stack pointer SP.
func (i *Intel8080) pop(x, y *byte) opHandler {
	return func() {
		n := i.stackPop()

		*x = uint8(n >> 8)
		*y = uint8(n & 0x00ff)
	}
}

// push is the "Push Data Onto Stack" handler.
//
// The contents of the specified register pair are saved in two bytes of memory
// indicated by the stack pointer SP.
func (i *Intel8080) push(x, y *byte) opHandler {
	return func() {
		w := uint16(*x)<<8 | uint16(*y)

		i.stackAdd(w)
	}
}

// popPSW is the "Pop Data Off Stack PSW" handler.
//
// The contents of the PSW register pair are restored from two bytes of
// memory indicated by the stack pointer SP.
func (i *Intel8080) popPSW() {
	n := i.stackPop()

	i.a = uint8(n >> 8)
	i.cc.setStatus(uint8(n & 0x00ff))
}

// pushPSW is the "Push Data Onto Stack PSW" handler.
//
// The contents of the PSW register pair are saved in two bytes of memory
// indicated by the stack pointer SP.
func (i *Intel8080) pushPSW() {
	i.stackAdd(uint16(i.a)<<8 | uint16(i.cc.status()))
}

// xthl is the "Exchange Stack" handler.
//
// The contents of the L register are exchanged with the contents of the memory
// byte whose address is held in the stack pointer SP. The contents of the H
// register are exchanged with the contents of the memory byte whose address is
// one greater than that held in the stack pointer.
func (i *Intel8080) xthl() {
	b := uint16(i.mem.Read(i.sp)) | uint16(i.mem.Read(i.sp+1))<<8
	hl := uint16(i.h)<<8 | uint16(i.l)

	i.h = uint8(b >> 8)
	i.l = uint8(b & 0x00ff)

	i.mem.Write(i.sp, uint8(hl&0xff))
	i.mem.Write(i.sp+1, uint8(hl>>8))
}

// sphl is the "Load SP from H and L" handler.
//
// The 16 bits of data held in the H and L registers replace the contents of the
// stack pointer SP. The contents of the H and L registers are unchanged.
func (i *Intel8080) sphl() {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := uint16(i.h)<<8 | uint16(i.l)

	i.sp = addr
}
