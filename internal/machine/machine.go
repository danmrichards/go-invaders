package machine

import (
	"github.com/danmrichards/go-invaders/internal/memory"
)

// Machine emulates the Space Invaders hardware.
type Machine struct {
	cpu cpuStepper

	// The Space Invaders Memory is mapped as follows:
	//
	// $0000-$1FFF -> 8K ROM
	// $2000-$23FF -> 1K RAM
	// $2400-$3FFF -> 7K Video RAM
	// $4000 -> RAM mirror
	//
	// For more details on the ROM structure see LoadROM.
	mem memory.ReadWriteDumper

	done <-chan struct{}
}

// New returns an instantiated Space Invaders machine.
func New(cpu cpuStepper, mem memory.ReadWriteDumper, done <-chan struct{}) *Machine {
	m := &Machine{
		cpu:  cpu,
		mem:  mem,
		done: done,
	}

	return m
}

// Run emulates the Space Invaders machine.
func (m *Machine) Run() error {
	for {
		select {
		case <-m.done:
			return nil
		default:
		}

		// Emulate an instruction.
		if err := m.cpu.Step(); err != nil {
			return err
		}
	}
}
