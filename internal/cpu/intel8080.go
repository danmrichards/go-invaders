package cpu

import (
	"fmt"

	"github.com/danmrichards/disassemble8080/pkg/dasm"
	"github.com/danmrichards/go-invaders/internal/memory"
)

// TODO: Abstract this to a separate repo and package.

// Intel8080 represents the Intel 8080 CPU.
type Intel8080 struct {
	// Working "scratchpad" registers.
	//
	// The 8080 allows these registers to be referenced in pairs, like so:
	//
	// b -> b & c
	// d -> d & e
	// h -> h & l
	b byte
	c byte
	d byte
	e byte
	h byte
	l byte

	// 8-bit accumulator.
	a byte

	// Stack pointer, stores address of last program request in the stack.
	sp uint16

	// Program counter, stores the address of the instruction being executed.
	pc uint16

	// Conditions represents the condition bits of the CPU.
	cc conditions

	// Interrupts enabled (1 = enabled, 0 = disabled).
	ie byte

	// Provides an interface to enable reads and writes to memory.
	mem memory.ReadWriteDumper

	// Each supported opcode has a handler function.
	opHandlers map[byte]opHandler

	// If set to true the emulation cycle will print debug information.
	debug bool
}

// Option is a functional option that modifies a field on the machine.
type Option func(*Intel8080)

// WithDebugEnabled enables debug mode on the machine.
func WithDebugEnabled() Option {
	return func(i *Intel8080) {
		i.debug = true
	}
}

// NewIntel8080 returns an instantiated Intel 8080.
func NewIntel8080(mem memory.ReadWriteDumper, opts ...Option) *Intel8080 {
	i := &Intel8080{
		cc:  newConditions(),
		mem: mem,
	}

	for _, o := range opts {
		o(i)
	}

	i.registerOpHandlers()

	return i
}

// Step emulates exactly one instruction on the Intel 8080.
func (i *Intel8080) Step() error {
	// Use the current value of the program counter to get the next opcode from
	// the attached memory.
	opc := i.mem.Read(i.pc)

	// Dump the assembly code if debug mode is on.
	if i.debug {
		asm, _ := dasm.Disassemble(i.mem.Dump(), int64(i.pc))
		fmt.Println(asm)
	}

	// Lookup the opcode handler.
	h, ok := i.opHandlers[opc]
	if !ok {
		return fmt.Errorf(
			"unsupported opcode 0x%02X at program counter %04x", opc, i.pc,
		)
	}

	// Handle the opcode and increment the program counter by the instruction
	// length.
	//
	// Imagine that we start at pc = 0, the first operation is 2 bytes long so
	// we increment the pc to 3 before continuing.
	i.pc += h()

	return nil
}
