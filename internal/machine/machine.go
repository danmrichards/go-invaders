package machine

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"

	"github.com/danmrichards/go-invaders/internal/cpu"
	"github.com/faiface/pixel/pixelgl"
)

type (
	// Machine emulates the Space Invaders hardware.
	Machine struct {
		c cpuStepper

		// The Space Invaders Memory is mapped as follows:
		//
		// $0000-$1FFF -> 8K ROM
		// $2000-$23FF -> 1K RAM
		// $2400-$3FFF -> 7K Video RAM
		// $4000 -> RAM mirror
		//
		// For more details on the ROM structure see LoadROM.
		m cpu.MemReadWriter

		w *pixelgl.Window

		debug bool
	}

	// Option is a functional option that modifies a field on the machine.
	Option func(*Machine)
)

// WithDebugEnabled enables debug mode on the machine.
func WithDebugEnabled() Option {
	return func(m *Machine) {
		m.debug = true
	}
}

// New returns an instantiated Space Invaders machine.
func New(mem cpu.MemReadWriter, opts ...Option) *Machine {
	m := &Machine{
		m: mem,
	}

	for _, o := range opts {
		o(m)
	}

	// Instantiate the CPU.
	var copts []cpu.Option
	if m.debug {
		copts = append(copts, cpu.WithDebugEnabled())
	}
	m.c = cpu.NewIntel8080(mem, m.input, m.output, copts...)

	return m
}

// Run emulates the Space Invaders machine.
func (m *Machine) Run() {
	var err error
	m.w, err = pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Space Invaders",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	})
	if err != nil {
		log.Fatalf("create window: %v", err)
	}

	// Emulation loop.
	for !m.w.Closed() {
		m.w.UpdateInput()

		// Emulate an instruction.
		if err := m.c.Step(); err != nil {
			log.Fatal(err)
		}
	}
}

func (m *Machine) input(b byte) {
	if m.debug {
		fmt.Printf("IN: %02x\n", b)
	}

	// TODO: Input handler
}

func (m *Machine) output(b byte) {
	if m.debug {
		fmt.Printf("OUT: %02x\n", b)
	}

	// TODO: Output handler
}
