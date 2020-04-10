package machine

import (
	"fmt"

	"github.com/danmrichards/disassemble8080/pkg/dasm"
)

// Cycle emulates exactly one operation (i.e. CPU cycle) on the Space Invaders
// machine.
func (m *Machine) Cycle() error {
	// Use the current value of the program counter to get the next opcode from
	// the machine memory.
	opc := m.mem[m.cpu.PC]

	// Dump the assembly code if debug mode is on.
	if m.debug {
		asm, _ := dasm.Disassemble(m.mem, int64(m.cpu.PC))
		fmt.Println(asm)
	}

	// Lookup the opcode handler.
	h, ok := m.opHandlers[opc]
	if !ok {
		return fmt.Errorf(
			"unsupported opcode 0x%02X at program counter %04x", opc, m.cpu.PC,
		)
	}

	// Handle the opcode and increment the program counter by the instruction
	// length.
	//
	// Imagine that we start at PC = 0, the first operation is 2 bytes long so
	// we increment the PC to 3 before continuing.
	m.cpu.PC += h()

	return nil
}
