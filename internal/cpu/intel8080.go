package cpu

import (
	"fmt"

	"github.com/danmrichards/disassemble8080/pkg/dasm"
)

// TODO: Abstract this to a separate repo and package.

type (
	// Intel8080 represents the Intel 8080 CPU.
	Intel8080 struct {
		// Registers including working "scratchpads" and the accumulator.
		R [8]byte

		// Stack pointer, stores address of last program request in the stack.
		sp uint16

		// Program counter, stores the address of the instruction being executed.
		pc uint16

		// Conditions represents the condition bits of the CPU.
		cc *conditions

		// Interrupts enabled.
		ie bool

		// Has the CPU been halted?
		halted bool

		// Provides an interface to enable reads and writes to memory.
		mem MemReadWriter

		// Input (i.e. keyboard) handler function.
		ih ifn

		// Output (i.e. sound) handler function.
		oh ofn

		// If set to true the emulation cycle will print debug information.
		debug bool
	}

	// Option is a functional option that modifies a field on the CPU.
	Option func(*Intel8080)

	// Input/Ouput handlers.
	ifn func(byte)
	ofn func(byte)
)

// WithDebugEnabled enables debug mode on the machine.
func WithDebugEnabled() Option {
	return func(i *Intel8080) {
		i.debug = true
	}
}

// NewIntel8080 returns an instantiated Intel 8080.
func NewIntel8080(mem MemReadWriter, in ifn, out ofn, opts ...Option) *Intel8080 {
	i := &Intel8080{
		cc:  &conditions{},
		mem: mem,
		ih:  in,
		oh:  out,
	}

	for _, o := range opts {
		o(i)
	}

	return i
}

// Step emulates exactly one instruction on the Intel 8080.
func (i *Intel8080) Step() error {
	// Use the current value of the program counter to get the next opcode from
	// the attached memory.
	opc := i.immediateByte()

	// Dump the assembly code if debug mode is on.
	if i.debug {
		asm, _ := dasm.Disassemble(i.mem.ReadAll(), int64(i.pc-1))

		fmt.Printf(
			"%s\tCY=%v\tAC=%v\tZ=%v\tP=%v\tS=%v\tSP=%04x\tA=%02x\tB=%02x\tC=%02x\tD=%02x\tE=%02x\tH=%02x\tL=%02x\n",
			asm,
			i.cc.cy,
			i.cc.ac,
			i.cc.z,
			i.cc.p,
			i.cc.s,
			i.sp,
			i.R[A],
			i.R[B],
			i.R[C],
			i.R[D],
			i.R[E],
			i.R[H],
			i.R[L],
		)
	}

	return i.handleOp(opc)
}

// immediateByte returns the next byte from memory indicated by the program
// counter.
//
// The program counter is incremented by one after the read.
func (i *Intel8080) immediateByte() byte {
	b := i.mem.Read(i.pc)
	i.pc++

	return b
}

// immediateWord returns the next two bytes from memory, merged, as a single word.
//
// The program counter is incremented by two after the read.
func (i *Intel8080) immediateWord() uint16 {
	lo := i.immediateByte()
	hi := i.immediateByte()

	return uint16(lo) | uint16(hi)<<8
}

// accumulatorAdd adds the given byte n to the accumulator and sets the relevant
// condition bits.
func (i *Intel8080) accumulatorAdd(n, carry byte) {
	// Perform the arithmetic at higher precision in order to capture the
	// carry out.
	ans := uint16(i.R[A]) + uint16(n) + uint16(carry)

	// Set the zero condition bit accordingly based on if the result of the
	// arithmetic was zero.
	i.cc.z = ans&0xff == 0x00

	// Set the sign condition bit accordingly based on if the most
	// significant bit on the result of the arithmetic was set.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0x80 (10000000 in base 2 and 128 in base 10).
	//
	// 10000000 & 10000000 = 1
	i.cc.s = (ans & 0x80) != 0

	// Set the carry condition bit accordingly if the result of the
	// arithmetic was greater than 0xff (11111111 in base 2 and 255 in base 10).
	i.cc.cy = (ans & 0x100) != 0

	// Set the auxiliary carry condition bit accordingly if the result of
	// the arithmetic has a carry on the third bit.
	i.cc.ac = (i.R[A]^uint8(ans)^n)&0x10 != 0

	// Set the parity bit.
	i.cc.setParity(uint8(ans))

	// Finally update the accumulator.
	i.R[A] = byte(ans)
}

// accumulatorSub subtracts the given byte n from the accumulator and sets the
// relevant condition bits.
func (i *Intel8080) accumulatorSub(n, carry byte) {
	// Perform the arithmetic at higher precision in order to capture the
	// carry out.
	ans := uint16(i.R[A]) - uint16(n) - uint16(carry)

	// Set the zero condition bit accordingly based on if the result of the
	// arithmetic was zero.
	i.cc.z = ans&0xff == 0x00

	// Set the sign condition bit accordingly based on if the most
	// significant bit on the result of the arithmetic was set.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0x80 (10000000 in base 2 and 128 in base 10).
	//
	// 10000000 & 10000000 = 1
	i.cc.s = (ans & 0x80) != 0

	// Set the carry condition bit accordingly if the result of the
	// arithmetic was greater than 0xff (11111111 in base 2 and 255 in base
	// 10).
	i.cc.cy = (ans & 0x100) != 0

	// Set the auxiliary carry condition bit accordingly if the result of
	// the arithmetic has a carry on the third bit.
	i.cc.ac = ^(i.R[A]^uint8(ans)^n)&0x10 != 0

	// Set the parity bit.
	i.cc.setParity(uint8(ans))

	// Finally update the accumulator.
	i.R[A] = byte(ans)
}

// stackAdd adds the given word to the stack.
func (i *Intel8080) stackAdd(n uint16) {
	i.sp -= 2
	i.mem.Write(i.sp, uint8(n&0xff))
	i.mem.Write(i.sp+1, uint8(n>>8))
}

// stackPop returns the immediate word from the stack as indicated by the stack
// pointer.
func (i *Intel8080) stackPop() uint16 {
	n := uint16(i.mem.Read(i.sp)) | uint16(i.mem.Read(i.sp+1))<<8
	i.sp += 2

	return n
}
