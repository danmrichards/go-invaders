package machine

import icpu "github.com/danmrichards/go-invaders/internal/cpu"

// Machine emulates the Space Invaders hardware.
type Machine struct {
	cpu *icpu.Intel8080

	// The Space Invaders Memory is mapped as follows:
	//
	// $0000-$1FFF -> 8K ROM
	// $2000-$23FF -> 1K RAM
	// $2400-$3FFF -> 7K Video RAM
	// $4000 -> RAM mirror
	//
	// For more details on the ROM structure see LoadROM.
	mem []byte

	// Each supported opcode has a handler function.
	opHandlers map[byte]opHandler

	// If set to true the emulation cycle will print debug information.
	debug bool
}

// Option is a functional option that modifies a field on the machine.
type Option func(*Machine)

// WithDebugEnabled enables debug mode on the machine.
func WithDebugEnabled() Option {
	return func(m *Machine) {
		m.debug = true
	}
}

// New returns an instantiated Space Invaders machine.
func New(opts ...Option) *Machine {
	m := &Machine{
		cpu: icpu.NewIntel8080(),
		mem: make([]byte, 16384),
	}

	for _, o := range opts {
		o(m)
	}

	m.registerOpHandlers()

	return m
}
