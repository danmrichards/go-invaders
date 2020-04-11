package cpu

// defaultInstructionLen represents the default length of an instruction.
const defaultInstructionLen = 1

// opHandler is a function that handles the instruction for a given opcode and
// returns the number of bytes to increment the program counter by.
//
// In the majority of cases the returned value will be the length of the current
// instruction.
type opHandler func() uint16

// registerOpHandlers registers the map of opcode handlers.
func (i *Intel8080) registerOpHandlers() {
	i.opHandlers = map[byte]opHandler{
		0x00: i.nop,
		0x01: i.lxi(&i.b, &i.c),
		//0x02: stax("b"),
		//0x03: inx("b"),
		//0x04: inr("b"),
		//0x05: dcr("b"),
		//0x06: mvi("b"),
		//0x07: rlc,
		//0x08: ignore,
		//0x09: dad("b"),
		//0x0a: ldax("b"),
		//0x0b: dcx("b"),
		//0x0c: inr("c"),
		//0x0d: dcr("c"),
		//0x0e: mvi("c"),
		//0x0f: rrc,
		//0x10: ignore,
		//0x11: lxi("d"),
		//0x12: stax("d"),
		//0x13: inx("d"),
		//0x14: inr("d"),
		//0x15: dcr("d"),
		//0x16: mvi("d"),
		//0x17: ral,
		//0x18: ignore,
		//0x19: dad("d"),
		//0x1a: ldax("d"),
		//0x1b: dcx("d"),
		//0x1c: inr("e"),
		//0x1d: dcr("e"),
		//0x1e: mvi("e"),
		//0x1f: rar,
		//0x20: rim,
		//0x21: lxi("h"),
		//0x22: shld,
		//0x23: inx("h"),
		//0x24: inr("h"),
		//0x25: dcr("h"),
		//0x26: mvi("h"),
		//0x27: daa,
		//0x28: ignore,
		//0x29: dad("h"),
		//0x2a: lhld,
		//0x2b: dcx("h"),
		//0x2c: inr("l"),
		//0x2d: dcr("l"),
		//0x2e: mvi("l"),
		//0x2f: cma,
		//0x30: sim,
		//0x31: lxi("sp"),
		//0x32: sta,
		//0x33: inx("sp"),
		//0x34: inr("M"),
		//0x35: dcr("M"),
		//0x36: mvi("M"),
		//0x37: stc,
		//0x38: ignore,
		//0x39: dad("sp"),
		//0x3a: lda,
		//0x3b: dcx("sp"),
		//0x3c: inr("a"),
		//0x3d: dcr("a"),
		//0x3e: mvi("a"),
		//0x3f: cmc,
		0x40: i.movRR(&i.b, &i.b),
		0x41: i.movRR(&i.b, &i.c),
		0x42: i.movRR(&i.b, &i.d),
		0x43: i.movRR(&i.b, &i.e),
		0x44: i.movRR(&i.b, &i.h),
		0x45: i.movRR(&i.b, &i.l),
		//0x46: i.movRR(&i.b, &i.M),
		0x47: i.movRR(&i.b, &i.a),
		0x48: i.movRR(&i.c, &i.b),
		0x49: i.movRR(&i.c, &i.c),
		0x4a: i.movRR(&i.c, &i.d),
		0x4b: i.movRR(&i.c, &i.e),
		0x4c: i.movRR(&i.c, &i.h),
		0x4d: i.movRR(&i.c, &i.l),
		//0x4e: i.movRR(&i.c, &i.M),
		0x4f: i.movRR(&i.c, &i.a),
		0x50: i.movRR(&i.d, &i.b),
		0x51: i.movRR(&i.d, &i.c),
		0x52: i.movRR(&i.d, &i.d),
		0x53: i.movRR(&i.d, &i.e),
		0x54: i.movRR(&i.d, &i.h),
		0x55: i.movRR(&i.d, &i.l),
		//0x56: i.movRR(&i.d, &i.M),
		0x57: i.movRR(&i.d, &i.a),
		0x58: i.movRR(&i.e, &i.b),
		0x59: i.movRR(&i.e, &i.c),
		0x5a: i.movRR(&i.e, &i.d),
		0x5b: i.movRR(&i.e, &i.e),
		0x5c: i.movRR(&i.e, &i.h),
		0x5d: i.movRR(&i.e, &i.l),
		//0x5e: i.movRR(&i.e, &i.M),
		0x5f: i.movRR(&i.e, &i.a),
		0x60: i.movRR(&i.h, &i.b),
		0x61: i.movRR(&i.h, &i.c),
		0x62: i.movRR(&i.h, &i.d),
		0x63: i.movRR(&i.h, &i.e),
		0x64: i.movRR(&i.h, &i.h),
		0x65: i.movRR(&i.h, &i.l),
		//0x66: i.movRR(&i.h, &i.M),
		0x67: i.movRR(&i.h, &i.a),
		0x68: i.movRR(&i.l, &i.b),
		0x69: i.movRR(&i.l, &i.c),
		0x6a: i.movRR(&i.l, &i.d),
		0x6b: i.movRR(&i.l, &i.e),
		0x6c: i.movRR(&i.l, &i.h),
		0x6d: i.movRR(&i.l, &i.l),
		//0x6e: i.movRR(&i.l, &i.M),
		0x6f: i.movRR(&i.l, &i.a),
		//0x70: i.movRR(&i.M, &i.b),
		//0x71: i.movRR(&i.M, &i.c),
		//0x72: i.movRR(&i.M, &i.d),
		//0x73: i.movRR(&i.M, &i.e),
		//0x74: i.movRR(&i.M, &i.h),
		//0x75: i.movRR(&i.M, &i.l),
		//0x76: hlt,
		//0x77: i.movRR(&i.M, &i.a),
		0x78: i.movRR(&i.a, &i.b),
		0x79: i.movRR(&i.a, &i.c),
		0x7a: i.movRR(&i.a, &i.d),
		0x7b: i.movRR(&i.a, &i.e),
		0x7c: i.movRR(&i.a, &i.h),
		0x7d: i.movRR(&i.a, &i.l),
		//0x7e: i.movRR(&i.a, &i.M),
		0x7f: i.movRR(&i.a, &i.a),
		//0x80: add("b"),
		//0x81: add("c"),
		//0x82: add("d"),
		//0x83: add("e"),
		//0x84: add("h"),
		//0x85: add("l"),
		//0x86: add("M"),
		//0x87: add("a"),
		//0x88: adc("b"),
		//0x89: adc("c"),
		//0x8a: adc("d"),
		//0x8b: adc("e"),
		//0x8c: adc("h"),
		//0x8d: adc("l"),
		//0x8e: adc("M"),
		//0x8f: adc("a"),
		//0x90: sub("b"),
		//0x91: sub("c"),
		//0x92: sub("d"),
		//0x93: sub("e"),
		//0x94: sub("h"),
		//0x95: sub("l"),
		//0x96: sub("M"),
		//0x97: sub("a"),
		//0x98: sbb("b"),
		//0x99: sbb("c"),
		//0x9a: sbb("d"),
		//0x9b: sbb("e"),
		//0x9c: sbb("h"),
		//0x9d: sbb("l"),
		//0x9e: sbb("M"),
		//0x9f: sbb("a"),
		//0xa0: ana("b"),
		//0xa1: ana("c"),
		//0xa2: ana("d"),
		//0xa3: ana("e"),
		//0xa4: ana("h"),
		//0xa5: ana("l"),
		//0xa6: ana("M"),
		//0xa7: ana("a"),
		//0xa8: xra("b"),
		//0xa9: xra("c"),
		//0xaa: xra("d"),
		//0xab: xra("e"),
		//0xac: xra("h"),
		//0xad: xra("l"),
		//0xae: xra("M"),
		//0xaf: xra("a"),
		//0xb0: ora("b"),
		//0xb1: ora("c"),
		//0xb2: ora("d"),
		//0xb3: ora("e"),
		//0xb4: ora("h"),
		//0xb5: ora("l"),
		//0xb6: ora("M"),
		//0xb7: ora("a"),
		//0xb8: cmp("b"),
		//0xb9: cmp("c"),
		//0xba: cmp("d"),
		//0xbb: cmp("e"),
		//0xbc: cmp("h"),
		//0xbd: cmp("l"),
		//0xbe: cmp("M"),
		//0xbf: cmp("a"),
		//0xc0: rnz,
		//0xc1: pop("b"),
		//0xc2: jnz,
		0xc3: i.jmp,
		//0xc4: cnz,
		//0xc5: push("b"),
		//0xc6: adi,
		//0xc7: rst(0),
		//0xc8: rz,
		//0xc9: ret,
		//0xca: jz,
		//0xcb: ignore,
		//0xcc: cz,
		//0xcd: call,
		//0xce: aci,
		//0xcf: rst(1),
		//0xd0: rnc,
		//0xd1: pop("d"),
		//0xd2: jnc,
		//0xd3: out,
		//0xd4: cnc,
		//0xd5: push("d"),
		//0xd6: sui,
		//0xd7: rst(2),
		//0xd8: rc,
		//0xd9: ignore,
		//0xda: jc,
		//0xdb: in,
		//0xdc: cc,
		//0xdd: ignore,
		//0xde: sbi,
		//0xdf: rst(3),
		//0xe0: rpo,
		//0xe1: pop("h"),
		//0xe2: jpo,
		//0xe3: xthl,
		//0xe4: cpo,
		//0xe5: push("h"),
		//0xe6: ani,
		//0xe7: rst(4),
		//0xe8: rpe,
		//0xe9: pchl,
		//0xea: jpe,
		//0xeb: xchg,
		//0xec: cpe,
		//0xed: ignore,
		//0xee: xri,
		//0xef: rst(5),
		//0xf0: rp,
		//0xf1: pop("PSW"),
		//0xf2: jp,
		//0xf3: di,
		//0xf4: cp,
		//0xf5: push("PSW"),
		//0xf6: ori,
		//0xf7: rst(6),
		//0xf8: rm,
		//0xf9: sphl,
		//0xfa: jm,
		//0xfb: ei,
		//0xfc: cm,
		//0xfd: ignore,
		//0xfe: cpi,
		//0xff: rst(7),
	}
}

// nop is a no-op and just returns the default instruction length.
func (i *Intel8080) nop() uint16 {
	return defaultInstructionLen
}

// lxi is the "Load Immediate Register" handler.
//
// This handler operates on a CPU register pair, the two components of the pair
// accessed by x and y.
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

// jmp is the "Jump" handler.
//
// This handler jumps the program counter to a given point in memory.
func (i *Intel8080) jmp() uint16 {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge thei.
	i.pc = uint16(i.mem.Read(i.pc+2))<<8 | uint16(i.mem.Read(i.pc+1))

	// As we're jumping the program counter there is no need to return a value
	// for the main cycle to increment the counter.
	return 0
}