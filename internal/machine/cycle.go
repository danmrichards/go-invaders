package machine

import "fmt"

// Cycle emulates exactly one operation (i.e. CPU cycle) on the Space Invaders
// machine.
func (m *Machine) Cycle() error {
	// Use the current value of the program counter to get the next opcode from
	// the machine memory.
	opc := m.mem[m.cpu.PC]

	fmt.Printf("OPCODE: %X PC: %04x\n", opc, m.cpu.PC)

	// TODO: Lookup opcode handler and execute. Should return number of bytes
	// to increment the program counter by.

	// TODO: This should be incremented by the number of bytes the executed
	// instruction needed.
	m.cpu.PC++

	return nil
}
