package machine

// defaultInstructionLen represents the default length of an instruction.
const defaultInstructionLen = 1

// opHandler is a function that handles the instruction for a given opcode and
// returns the number of bytes to increment the program counter by.
//
// In the majority of cases the returned value will be the length of the current
// instruction.
type opHandler func() uint16

// registerOpHandlers registers the map of opcode handlers.
func (m *Machine) registerOpHandlers() {
	m.opHandlers = map[byte]opHandler{
		0x00: m.nop,
		0x01: m.lxi(&m.cpu.B, &m.cpu.C),
		//0x02: stax("B"),
		//0x03: inx("B"),
		//0x04: inr("B"),
		//0x05: dcr("B"),
		//0x06: mvi("B"),
		//0x07: rlc,
		//0x08: ignore,
		//0x09: dad("B"),
		//0x0a: ldax("B"),
		//0x0b: dcx("B"),
		//0x0c: inr("C"),
		//0x0d: dcr("C"),
		//0x0e: mvi("C"),
		//0x0f: rrc,
		//0x10: ignore,
		//0x11: lxi("D"),
		//0x12: stax("D"),
		//0x13: inx("D"),
		//0x14: inr("D"),
		//0x15: dcr("D"),
		//0x16: mvi("D"),
		//0x17: ral,
		//0x18: ignore,
		//0x19: dad("D"),
		//0x1a: ldax("D"),
		//0x1b: dcx("D"),
		//0x1c: inr("E"),
		//0x1d: dcr("E"),
		//0x1e: mvi("E"),
		//0x1f: rar,
		//0x20: rim,
		//0x21: lxi("H"),
		//0x22: shld,
		//0x23: inx("H"),
		//0x24: inr("H"),
		//0x25: dcr("H"),
		//0x26: mvi("H"),
		//0x27: daa,
		//0x28: ignore,
		//0x29: dad("H"),
		//0x2a: lhld,
		//0x2b: dcx("H"),
		//0x2c: inr("L"),
		//0x2d: dcr("L"),
		//0x2e: mvi("L"),
		//0x2f: cma,
		//0x30: sim,
		//0x31: lxi("SP"),
		//0x32: sta,
		//0x33: inx("SP"),
		//0x34: inr("M"),
		//0x35: dcr("M"),
		//0x36: mvi("M"),
		//0x37: stc,
		//0x38: ignore,
		//0x39: dad("SP"),
		//0x3a: lda,
		//0x3b: dcx("SP"),
		//0x3c: inr("A"),
		//0x3d: dcr("A"),
		//0x3e: mvi("A"),
		//0x3f: cmc,
		0x40: m.movRR(&m.cpu.B, &m.cpu.B),
		0x41: m.movRR(&m.cpu.B, &m.cpu.C),
		0x42: m.movRR(&m.cpu.B, &m.cpu.D),
		0x43: m.movRR(&m.cpu.B, &m.cpu.E),
		0x44: m.movRR(&m.cpu.B, &m.cpu.H),
		0x45: m.movRR(&m.cpu.B, &m.cpu.L),
		//0x46: m.movRR(&m.cpu.B, &m.cpu.M),
		0x47: m.movRR(&m.cpu.B, &m.cpu.A),
		0x48: m.movRR(&m.cpu.C, &m.cpu.B),
		0x49: m.movRR(&m.cpu.C, &m.cpu.C),
		0x4a: m.movRR(&m.cpu.C, &m.cpu.D),
		0x4b: m.movRR(&m.cpu.C, &m.cpu.E),
		0x4c: m.movRR(&m.cpu.C, &m.cpu.H),
		0x4d: m.movRR(&m.cpu.C, &m.cpu.L),
		//0x4e: m.movRR(&m.cpu.C, &m.cpu.M),
		0x4f: m.movRR(&m.cpu.C, &m.cpu.A),
		0x50: m.movRR(&m.cpu.D, &m.cpu.B),
		0x51: m.movRR(&m.cpu.D, &m.cpu.C),
		0x52: m.movRR(&m.cpu.D, &m.cpu.D),
		0x53: m.movRR(&m.cpu.D, &m.cpu.E),
		0x54: m.movRR(&m.cpu.D, &m.cpu.H),
		0x55: m.movRR(&m.cpu.D, &m.cpu.L),
		//0x56: m.movRR(&m.cpu.D, &m.cpu.M),
		0x57: m.movRR(&m.cpu.D, &m.cpu.A),
		0x58: m.movRR(&m.cpu.E, &m.cpu.B),
		0x59: m.movRR(&m.cpu.E, &m.cpu.C),
		0x5a: m.movRR(&m.cpu.E, &m.cpu.D),
		0x5b: m.movRR(&m.cpu.E, &m.cpu.E),
		0x5c: m.movRR(&m.cpu.E, &m.cpu.H),
		0x5d: m.movRR(&m.cpu.E, &m.cpu.L),
		//0x5e: m.movRR(&m.cpu.E, &m.cpu.M),
		0x5f: m.movRR(&m.cpu.E, &m.cpu.A),
		0x60: m.movRR(&m.cpu.H, &m.cpu.B),
		0x61: m.movRR(&m.cpu.H, &m.cpu.C),
		0x62: m.movRR(&m.cpu.H, &m.cpu.D),
		0x63: m.movRR(&m.cpu.H, &m.cpu.E),
		0x64: m.movRR(&m.cpu.H, &m.cpu.H),
		0x65: m.movRR(&m.cpu.H, &m.cpu.L),
		//0x66: m.movRR(&m.cpu.H, &m.cpu.M),
		0x67: m.movRR(&m.cpu.H, &m.cpu.A),
		0x68: m.movRR(&m.cpu.L, &m.cpu.B),
		0x69: m.movRR(&m.cpu.L, &m.cpu.C),
		0x6a: m.movRR(&m.cpu.L, &m.cpu.D),
		0x6b: m.movRR(&m.cpu.L, &m.cpu.E),
		0x6c: m.movRR(&m.cpu.L, &m.cpu.H),
		0x6d: m.movRR(&m.cpu.L, &m.cpu.L),
		//0x6e: m.movRR(&m.cpu.L, &m.cpu.M),
		0x6f: m.movRR(&m.cpu.L, &m.cpu.A),
		//0x70: m.movRR(&m.cpu.M, &m.cpu.B),
		//0x71: m.movRR(&m.cpu.M, &m.cpu.C),
		//0x72: m.movRR(&m.cpu.M, &m.cpu.D),
		//0x73: m.movRR(&m.cpu.M, &m.cpu.E),
		//0x74: m.movRR(&m.cpu.M, &m.cpu.H),
		//0x75: m.movRR(&m.cpu.M, &m.cpu.L),
		//0x76: hlt,
		//0x77: m.movRR(&m.cpu.M, &m.cpu.A),
		0x78: m.movRR(&m.cpu.A, &m.cpu.B),
		0x79: m.movRR(&m.cpu.A, &m.cpu.C),
		0x7a: m.movRR(&m.cpu.A, &m.cpu.D),
		0x7b: m.movRR(&m.cpu.A, &m.cpu.E),
		0x7c: m.movRR(&m.cpu.A, &m.cpu.H),
		0x7d: m.movRR(&m.cpu.A, &m.cpu.L),
		//0x7e: m.movRR(&m.cpu.A, &m.cpu.M),
		0x7f: m.movRR(&m.cpu.A, &m.cpu.A),
		//0x80: add("B"),
		//0x81: add("C"),
		//0x82: add("D"),
		//0x83: add("E"),
		//0x84: add("H"),
		//0x85: add("L"),
		//0x86: add("M"),
		//0x87: add("A"),
		//0x88: adc("B"),
		//0x89: adc("C"),
		//0x8a: adc("D"),
		//0x8b: adc("E"),
		//0x8c: adc("H"),
		//0x8d: adc("L"),
		//0x8e: adc("M"),
		//0x8f: adc("A"),
		//0x90: sub("B"),
		//0x91: sub("C"),
		//0x92: sub("D"),
		//0x93: sub("E"),
		//0x94: sub("H"),
		//0x95: sub("L"),
		//0x96: sub("M"),
		//0x97: sub("A"),
		//0x98: sbb("B"),
		//0x99: sbb("C"),
		//0x9a: sbb("D"),
		//0x9b: sbb("E"),
		//0x9c: sbb("H"),
		//0x9d: sbb("L"),
		//0x9e: sbb("M"),
		//0x9f: sbb("A"),
		//0xa0: ana("B"),
		//0xa1: ana("C"),
		//0xa2: ana("D"),
		//0xa3: ana("E"),
		//0xa4: ana("H"),
		//0xa5: ana("L"),
		//0xa6: ana("M"),
		//0xa7: ana("A"),
		//0xa8: xra("B"),
		//0xa9: xra("C"),
		//0xaa: xra("D"),
		//0xab: xra("E"),
		//0xac: xra("H"),
		//0xad: xra("L"),
		//0xae: xra("M"),
		//0xaf: xra("A"),
		//0xb0: ora("B"),
		//0xb1: ora("C"),
		//0xb2: ora("D"),
		//0xb3: ora("E"),
		//0xb4: ora("H"),
		//0xb5: ora("L"),
		//0xb6: ora("M"),
		//0xb7: ora("A"),
		//0xb8: cmp("B"),
		//0xb9: cmp("C"),
		//0xba: cmp("D"),
		//0xbb: cmp("E"),
		//0xbc: cmp("H"),
		//0xbd: cmp("L"),
		//0xbe: cmp("M"),
		//0xbf: cmp("A"),
		//0xc0: rnz,
		//0xc1: pop("B"),
		//0xc2: jnz,
		0xc3: m.jmp,
		//0xc4: cnz,
		//0xc5: push("B"),
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
		//0xd1: pop("D"),
		//0xd2: jnc,
		//0xd3: out,
		//0xd4: cnc,
		//0xd5: push("D"),
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
		//0xe1: pop("H"),
		//0xe2: jpo,
		//0xe3: xthl,
		//0xe4: cpo,
		//0xe5: push("H"),
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
func (m *Machine) nop() uint16 {
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
func (m *Machine) lxi(x, y *byte) opHandler {
	return func() uint16 {
		*x = m.mem[m.cpu.PC+2]
		*y = m.mem[m.cpu.PC+1]

		return 3
	}
}

// movRR is the "Move Register to Register" handler.
//
// One byte of data is moved from the register specified by src (the source
// register) to the register specified by dst (the destination register).
func (m *Machine) movRR(dst, src *byte) opHandler {
	return func() uint16 {
		*dst = *src

		return defaultInstructionLen
	}
}

// jmp is the "Jump" handler.
//
// This handler jumps the program counter to a given point in memory.
func (m *Machine) jmp() uint16 {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	m.cpu.PC = uint16(m.mem[m.cpu.PC+2])<<8 | uint16(m.mem[m.cpu.PC+1])

	// As we're jumping the program counter there is no need to return a value
	// for the main cycle to increment the counter.
	return 0
}
