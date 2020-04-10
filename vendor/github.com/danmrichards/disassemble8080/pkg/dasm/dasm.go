package dasm

import (
	"fmt"
	"strings"
)

// Disassemble returns a string representing the assembly code for the current
// opcode and the length of the relevant instruction.
//
// The opcode is determined by reading the bytes indicated by the given program
// counter from the given ROM bytes.
func Disassemble(rb []byte, pc int64) (string, int64) {
	// Get the opcode.
	opc := rb[pc]

	// Build the output.
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%04x ", pc))

	var (
		asm string
		opb = int64(1)
	)
	d, ok := disassemblers[opc]
	if ok {
		asm, opb = d(rb, pc)
		sb.WriteString(asm)
	} else {
		sb.WriteString(fmt.Sprintf("OPCODE: %x UNKNOWN", opc))
	}

	return sb.String(), opb
}
