package machine

import (
	icpu "github.com/danmrichards/go-invaders/internal/cpu"
)

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

	// TODO: Opcode handlers map.
}

// New returns an instantiated Space Invaders machine.
func New() *Machine {
	return &Machine{
		cpu: icpu.NewIntel8080(),
		mem: make([]byte, 16384),
	}
}
