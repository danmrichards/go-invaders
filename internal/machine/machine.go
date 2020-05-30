package machine

import (
	"image/color"
	"log"
	"time"

	"github.com/danmrichards/go-invaders/internal/sound"
	cpu "github.com/danmrichards/go8080"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	// Screen dimensions. The native Space Invaders resolution is 224x256, but
	// we're also adding a scale factor to allow rendering at a higher
	// resolution on modern displays.
	screenW, screenH = 224, 256

	// The original Space Invaders machine ran at a clock speed of 2MHz. The
	// screen refreshed at a rate of 60Hz. Make some rough calculations at how
	// many CPU cycles we should process per frame.
	screenRefresh         = 60
	clockSpeed            = 2000000
	cyclesPerFrame uint32 = clockSpeed / screenRefresh

	// The video RAM starts at address 0x2400 in the machine memory.
	vramStart uint16 = 0x2400
)

type (
	// Machine emulates the Space Invaders hardware.
	Machine struct {
		c processor

		// The Space Invaders Memory is mapped as follows:
		//
		// $0000-$1FFF -> 8K ROM
		// $2000-$23FF -> 1K RAM
		// $2400-$3FFF -> 7K Video RAM
		// $4000 -> RAM mirror
		//
		// For more details on the ROM structure see LoadROM.
		mem cpu.MemReadWriter

		// The render window.
		w *pixelgl.Window

		// Sound player.
		p *sound.Player

		// The address of the next interrupt to send to the CPU.
		ni uint16

		// The video scale factor.
		sf int

		// The Intel 8080 does not include opcodes for shifting by anything
		// other than 1 bit. Hence it would take thousands of instruction calls
		// to perform a multi-bit shift.
		//
		// Consequently the Space Invaders machine had hardware for bit
		// shifting. Store the data to shift and the offset we'll shift by.
		so uint16
		sd uint16

		// Watchdog (read or write to reset).
		wd byte

		// Sound banks which store the currently playing sound.
		snd1 byte
		snd2 byte

		// Flag for debug mode.
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

// WithScaleFactor sets the video scale factor.
func WithScaleFactor(sf int) Option {
	return func(m *Machine) {
		m.sf = sf
	}
}

// New returns an instantiated Space Invaders machine.
func New(mem cpu.MemReadWriter, opts ...Option) (m *Machine, err error) {
	m = &Machine{
		mem: mem,
		ni:  0x08,
	}

	for _, o := range opts {
		o(m)
	}

	// Instantiate the CPU.
	copts := []cpu.Option{
		cpu.WithInput(m.input),
		cpu.WithOutput(m.output),
	}
	if m.debug {
		copts = append(copts, cpu.WithDebugEnabled())
	}
	m.c = cpu.NewIntel8080(mem, copts...)

	m.p, err = sound.NewPlayer()
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Run emulates the Space Invaders machine.
func (m *Machine) Run() {
	var err error
	m.w, err = pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Space Invaders",
		Bounds: pixel.R(0, 0, screenW*float64(m.sf), screenH*float64(m.sf)),
		VSync:  true,
	})
	if err != nil {
		log.Fatalf("create window: %v", err)
	}

	start := time.Now()

	for !m.w.Closed() && m.c.Running() {
		// Throttle to one step per ~16ms, to better reproduce the speed of the
		// original machine.
		dt := time.Since(start).Milliseconds()
		if float64(dt) > (1/float64(screenRefresh))*1000 {
			if err := m.step(); err != nil {
				log.Fatalf("step: %v\n", err)
			}
			m.render()
		}

		// Clear the screen, ready for the next frame.
		m.w.Clear(color.Black)
	}
}

// step performs the core CPU emulation for the machine.
//
// The space invaders machine handles it's frame rendering in two parts. It
// would draw the top half of the screen and then send the first interrupt
// (RST 8). It would then move on, draw the lower half of the screen and
// send the second interrupt (RST 10).
func (m *Machine) step() error {
	// Work out how many CPU cycles to run in half a frame. This will
	// synchronise the emulation process with the rendering process in mem.render.
	hfc := cyclesPerFrame / 2

	// Run the cycles for the first half of the frame, recording the delta in
	// cycle count at each step call.
	cyc := uint32(0)
	for cyc <= hfc {
		sc := m.c.Cycles()
		if err := m.c.Step(); err != nil {
			return err
		}
		cyc += m.c.Cycles() - sc
	}

	// Fire off the first interrupt and determine what the next one should be.
	m.c.Interrupt(m.ni)
	if m.ni == 0x08 {
		m.ni = 0x10
	} else {
		m.ni = 0x08
	}

	// Run the cycles for the second half of the frame.
	for cyc <= cyclesPerFrame {
		sc := m.c.Cycles()
		if err := m.c.Step(); err != nil {
			return err
		}
		cyc += m.c.Cycles() - sc
	}

	return nil
}
