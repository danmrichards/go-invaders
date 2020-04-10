package dasm

import "fmt"

const defaultInstructionLen = 1

// disassembler is a function that returns the relevant assembly code and length
// of the instruction.
//
// Assembly and instruction length are calculated from the given ROM bytes and
// program counter.
type disassembler func([]byte, int64) (string, int64)

var disassemblers = map[byte]disassembler{
	0x00: nop,
	0x01: lxi("B"),
	0x02: stax("B"),
	0x03: inx("B"),
	0x04: inr("B"),
	0x05: dcr("B"),
	0x06: mvi("B"),
	0x07: rlc,
	0x08: ignore,
	0x09: dad("B"),
	0x0a: ldax("B"),
	0x0b: dcx("B"),
	0x0c: inr("C"),
	0x0d: dcr("C"),
	0x0e: mvi("C"),
	0x0f: rrc,
	0x10: ignore,
	0x11: lxi("D"),
	0x12: stax("D"),
	0x13: inx("D"),
	0x14: inr("D"),
	0x15: dcr("D"),
	0x16: mvi("D"),
	0x17: ral,
	0x18: ignore,
	0x19: dad("D"),
	0x1a: ldax("D"),
	0x1b: dcx("D"),
	0x1c: inr("E"),
	0x1d: dcr("E"),
	0x1e: mvi("E"),
	0x1f: rar,
	0x20: rim,
	0x21: lxi("H"),
	0x22: shld,
	0x23: inx("H"),
	0x24: inr("H"),
	0x25: dcr("H"),
	0x26: mvi("H"),
	0x27: daa,
	0x28: ignore,
	0x29: dad("H"),
	0x2a: lhld,
	0x2b: dcx("H"),
	0x2c: inr("L"),
	0x2d: dcr("L"),
	0x2e: mvi("L"),
	0x2f: cma,
	0x30: sim,
	0x31: lxi("SP"),
	0x32: sta,
	0x33: inx("SP"),
	0x34: inr("M"),
	0x35: dcr("M"),
	0x36: mvi("M"),
	0x37: stc,
	0x38: ignore,
	0x39: dad("SP"),
	0x3a: lda,
	0x3b: dcx("SP"),
	0x3c: inr("A"),
	0x3d: dcr("A"),
	0x3e: mvi("A"),
	0x3f: cmc,
	0x40: mov("B", "B"),
	0x41: mov("B", "C"),
	0x42: mov("B", "D"),
	0x43: mov("B", "E"),
	0x44: mov("B", "H"),
	0x45: mov("B", "L"),
	0x46: mov("B", "M"),
	0x47: mov("B", "A"),
	0x48: mov("C", "B"),
	0x49: mov("C", "C"),
	0x4a: mov("C", "D"),
	0x4b: mov("C", "E"),
	0x4c: mov("C", "H"),
	0x4d: mov("C", "L"),
	0x4e: mov("C", "M"),
	0x4f: mov("C", "A"),
	0x50: mov("D", "B"),
	0x51: mov("D", "C"),
	0x52: mov("D", "D"),
	0x53: mov("D", "E"),
	0x54: mov("D", "H"),
	0x55: mov("D", "L"),
	0x56: mov("D", "M"),
	0x57: mov("D", "A"),
	0x58: mov("E", "B"),
	0x59: mov("E", "C"),
	0x5a: mov("E", "D"),
	0x5b: mov("E", "E"),
	0x5c: mov("E", "H"),
	0x5d: mov("E", "L"),
	0x5e: mov("E", "M"),
	0x5f: mov("E", "A"),
	0x60: mov("H", "B"),
	0x61: mov("H", "C"),
	0x62: mov("H", "D"),
	0x63: mov("H", "E"),
	0x64: mov("H", "H"),
	0x65: mov("H", "L"),
	0x66: mov("H", "M"),
	0x67: mov("H", "A"),
	0x68: mov("L", "B"),
	0x69: mov("L", "C"),
	0x6a: mov("L", "D"),
	0x6b: mov("L", "E"),
	0x6c: mov("L", "H"),
	0x6d: mov("L", "L"),
	0x6e: mov("L", "M"),
	0x6f: mov("L", "A"),
	0x70: mov("M", "B"),
	0x71: mov("M", "C"),
	0x72: mov("M", "D"),
	0x73: mov("M", "E"),
	0x74: mov("M", "H"),
	0x75: mov("M", "L"),
	0x76: hlt,
	0x77: mov("M", "A"),
	0x78: mov("A", "B"),
	0x79: mov("A", "C"),
	0x7a: mov("A", "D"),
	0x7b: mov("A", "E"),
	0x7c: mov("A", "H"),
	0x7d: mov("A", "L"),
	0x7e: mov("A", "M"),
	0x7f: mov("A", "A"),
	0x80: add("B"),
	0x81: add("C"),
	0x82: add("D"),
	0x83: add("E"),
	0x84: add("H"),
	0x85: add("L"),
	0x86: add("M"),
	0x87: add("A"),
	0x88: adc("B"),
	0x89: adc("C"),
	0x8a: adc("D"),
	0x8b: adc("E"),
	0x8c: adc("H"),
	0x8d: adc("L"),
	0x8e: adc("M"),
	0x8f: adc("A"),
	0x90: sub("B"),
	0x91: sub("C"),
	0x92: sub("D"),
	0x93: sub("E"),
	0x94: sub("H"),
	0x95: sub("L"),
	0x96: sub("M"),
	0x97: sub("A"),
	0x98: sbb("B"),
	0x99: sbb("C"),
	0x9a: sbb("D"),
	0x9b: sbb("E"),
	0x9c: sbb("H"),
	0x9d: sbb("L"),
	0x9e: sbb("M"),
	0x9f: sbb("A"),
	0xa0: ana("B"),
	0xa1: ana("C"),
	0xa2: ana("D"),
	0xa3: ana("E"),
	0xa4: ana("H"),
	0xa5: ana("L"),
	0xa6: ana("M"),
	0xa7: ana("A"),
	0xa8: xra("B"),
	0xa9: xra("C"),
	0xaa: xra("D"),
	0xab: xra("E"),
	0xac: xra("H"),
	0xad: xra("L"),
	0xae: xra("M"),
	0xaf: xra("A"),
	0xb0: ora("B"),
	0xb1: ora("C"),
	0xb2: ora("D"),
	0xb3: ora("E"),
	0xb4: ora("H"),
	0xb5: ora("L"),
	0xb6: ora("M"),
	0xb7: ora("A"),
	0xb8: cmp("B"),
	0xb9: cmp("C"),
	0xba: cmp("D"),
	0xbb: cmp("E"),
	0xbc: cmp("H"),
	0xbd: cmp("L"),
	0xbe: cmp("M"),
	0xbf: cmp("A"),
	0xc0: rnz,
	0xc1: pop("B"),
	0xc2: jnz,
	0xc3: jmp,
	0xc4: cnz,
	0xc5: push("B"),
	0xc6: adi,
	0xc7: rst(0),
	0xc8: rz,
	0xc9: ret,
	0xca: jz,
	0xcb: ignore,
	0xcc: cz,
	0xcd: call,
	0xce: aci,
	0xcf: rst(1),
	0xd0: rnc,
	0xd1: pop("D"),
	0xd2: jnc,
	0xd3: out,
	0xd4: cnc,
	0xd5: push("D"),
	0xd6: sui,
	0xd7: rst(2),
	0xd8: rc,
	0xd9: ignore,
	0xda: jc,
	0xdb: in,
	0xdc: cc,
	0xdd: ignore,
	0xde: sbi,
	0xdf: rst(3),
	0xe0: rpo,
	0xe1: pop("H"),
	0xe2: jpo,
	0xe3: xthl,
	0xe4: cpo,
	0xe5: push("H"),
	0xe6: ani,
	0xe7: rst(4),
	0xe8: rpe,
	0xe9: pchl,
	0xea: jpe,
	0xeb: xchg,
	0xec: cpe,
	0xed: ignore,
	0xee: xri,
	0xef: rst(5),
	0xf0: rp,
	0xf1: pop("PSW"),
	0xf2: jp,
	0xf3: di,
	0xf4: cp,
	0xf5: push("PSW"),
	0xf6: ori,
	0xf7: rst(6),
	0xf8: rm,
	0xf9: sphl,
	0xfa: jm,
	0xfb: ei,
	0xfc: cm,
	0xfd: ignore,
	0xfe: cpi,
	0xff: rst(7),
}

func nop(rb []byte, pc int64) (string, int64) {
	return "NOP", defaultInstructionLen
}

func lxi(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return fmt.Sprintf("LXI %s,#$%02X%02X", r, rb[pc+2], rb[pc+1]), 3
	}
}

func stax(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "STAX " + r, defaultInstructionLen
	}
}

func inx(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "INX " + r, defaultInstructionLen
	}
}

func inr(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "INR " + r, defaultInstructionLen
	}
}

func dcr(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "DCR " + r, defaultInstructionLen
	}
}

func mvi(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return fmt.Sprintf("MVI %s,#$%02X", r, rb[pc+1]), 2
	}
}

func rlc(rb []byte, pc int64) (string, int64) {
	return "RLC", defaultInstructionLen
}

func jmp(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("JMP $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func push(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "PUSH " + r, defaultInstructionLen
	}
}

func ignore(rb []byte, pc int64) (string, int64) {
	return "-", defaultInstructionLen
}

func dad(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "DAD " + r, defaultInstructionLen
	}
}

func ldax(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "LDAX " + r, defaultInstructionLen
	}
}

func dcx(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "DCX " + r, defaultInstructionLen
	}
}

func rrc(rb []byte, pc int64) (string, int64) {
	return "RRC", defaultInstructionLen
}

func ral(rb []byte, pc int64) (string, int64) {
	return "RAL", defaultInstructionLen
}

func rar(rb []byte, pc int64) (string, int64) {
	return "RAR", defaultInstructionLen
}

func rim(rb []byte, pc int64) (string, int64) {
	return "RIM", defaultInstructionLen
}

func shld(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("SHLD $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func daa(rb []byte, pc int64) (string, int64) {
	return "DAA", defaultInstructionLen
}

func lhld(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("LHLD $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func cma(rb []byte, pc int64) (string, int64) {
	return "CMA", defaultInstructionLen
}

func sim(rb []byte, pc int64) (string, int64) {
	return "SIM", defaultInstructionLen
}

func sta(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("STA $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func stc(rb []byte, pc int64) (string, int64) {
	return "STC", defaultInstructionLen
}

func lda(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("LDA $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func cmc(rb []byte, pc int64) (string, int64) {
	return "CMC", defaultInstructionLen
}

func mov(dst, src string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return fmt.Sprintf("MOV %s,%s", dst, src), defaultInstructionLen
	}
}

func hlt(rb []byte, pc int64) (string, int64) {
	return "HLT", defaultInstructionLen
}

func add(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "ADD " + r, defaultInstructionLen
	}
}

func adc(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "ADC " + r, defaultInstructionLen
	}
}

func sub(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "SUB " + r, defaultInstructionLen
	}
}

func sbb(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "SBB " + r, defaultInstructionLen
	}
}

func ana(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "ANA " + r, defaultInstructionLen
	}
}

func xra(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "XRA " + r, defaultInstructionLen
	}
}

func ora(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "ORA " + r, defaultInstructionLen
	}
}

func cmp(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "CMP " + r, defaultInstructionLen
	}
}

func rnz(rb []byte, pc int64) (string, int64) {
	return "RNZ", defaultInstructionLen
}

func pop(r string) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return "POP " + r, defaultInstructionLen
	}
}

func jnz(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("JNZ $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func cnz(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("CNZ $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func adi(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("ADI #$%02X", rb[pc+1]), 2
}

func rst(i int) disassembler {
	return func(rb []byte, pc int64) (string, int64) {
		return fmt.Sprintf("RST %d", i), defaultInstructionLen
	}
}

func rz(rb []byte, pc int64) (string, int64) {
	return "RZ", defaultInstructionLen
}

func ret(rb []byte, pc int64) (string, int64) {
	return "RET", defaultInstructionLen
}

func jz(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("JZ $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func cz(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("CZ $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func call(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("CALL $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func aci(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("ACI #$%02X", rb[pc+1]), 2
}

func rnc(rb []byte, pc int64) (string, int64) {
	return "RNC", defaultInstructionLen
}

func jnc(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("JNC $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func out(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("OUT #$%02X", rb[pc+1]), 2
}

func cnc(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("CNC $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func sui(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("SUI #$%02X", rb[pc+1]), 2
}

func rc(rb []byte, pc int64) (string, int64) {
	return "RC", defaultInstructionLen
}

func jc(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("JC $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func in(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("IN #$%02X", rb[pc+1]), 2
}

func cc(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("CC $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func sbi(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("SBI #$%02X", rb[pc+1]), 2
}

func rpo(rb []byte, pc int64) (string, int64) {
	return "RPO", defaultInstructionLen
}

func jpo(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("JPO $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func xthl(rb []byte, pc int64) (string, int64) {
	return "XTHL", defaultInstructionLen
}

func cpo(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("CPO $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func ani(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("ANI #$%02X", rb[pc+1]), 2
}

func rpe(rb []byte, pc int64) (string, int64) {
	return "RPE", defaultInstructionLen
}

func pchl(rb []byte, pc int64) (string, int64) {
	return "PCHL", defaultInstructionLen
}

func jpe(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("JPE $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func xchg(rb []byte, pc int64) (string, int64) {
	return "XCHG", defaultInstructionLen
}

func cpe(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("CPE $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func xri(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("XRI #$%02X", rb[pc+1]), 2
}

func rp(rb []byte, pc int64) (string, int64) {
	return "RP", defaultInstructionLen
}

func jp(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("JP $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func di(rb []byte, pc int64) (string, int64) {
	return "DI", defaultInstructionLen
}

func cp(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("CP $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func ori(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("ORI #$%02X", rb[pc+1]), 2
}

func rm(rb []byte, pc int64) (string, int64) {
	return "RM", defaultInstructionLen
}

func sphl(rb []byte, pc int64) (string, int64) {
	return "SPHL", defaultInstructionLen
}

func jm(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("JM $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func ei(rb []byte, pc int64) (string, int64) {
	return "EI", defaultInstructionLen
}

func cm(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("CM $%02X%02X", rb[pc+2], rb[pc+1]), 3
}

func cpi(rb []byte, pc int64) (string, int64) {
	return fmt.Sprintf("CPI #$%02X", rb[pc+1]), 2
}
