package machine

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/danmrichards/go-invaders/internal/cpu"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	// Screen dimensions. The native Space Invaders resolution is 224x256, but
	// we're also adding a scaler factor to allow rendering at a higher
	// resolution on modern displays.
	//
	// TODO: Make scale factor configurable.
	screenW, screenH = 224, 256
	scaleFactor      = 3

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
		m cpu.MemReadWriter

		// The render window.
		w *pixelgl.Window

		// The address of the next interrupt to send to the CPU.
		ni uint16

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

// New returns an instantiated Space Invaders machine.
func New(mem cpu.MemReadWriter, opts ...Option) *Machine {
	m := &Machine{
		m:  mem,
		ni: 0x08,
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

	return m
}

// Run emulates the Space Invaders machine.
func (m *Machine) Run() {
	var err error
	m.w, err = pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Space Invaders",
		Bounds: pixel.R(0, 0, screenW*scaleFactor, screenH*scaleFactor),
		VSync:  true,
	})
	if err != nil {
		log.Fatalf("create window: %v", err)
	}

	start := time.Now()

	// TODO: Implement keyboard input.

	// TODO: Also check for halted here.
	for !m.w.Closed() {
		dt := time.Since(start)
		time.Now()
		if float64(dt.Milliseconds()) > (1/float64(60))*1000 {
			if err := m.step(); err != nil {
				log.Fatalf("step: %w\n", err)
			}
			m.render()
		}
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
	// synchronise the emulation process with the rendering process in m.render.
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

// render renders the current screen to the window.
func (m *Machine) render() {
	// Prepare the drawing object.
	imd := imdraw.New(nil)
	imd.Color = color.White

	// Draw the screen.
	m.draw(imd)

	// Update the window.
	imd.Draw(m.w)
	m.w.Update()

	// Clear the screen, ready for the next frame.
	m.w.Clear(color.Black)
}

// draw draws the current screen onto the given draw object.
//
// The screen is drawn by iterating over the range of memory between the VRAM
// start address and the start address plus 256*224 bytes. Each byte in this
// range represents 8 pixels.
func (m *Machine) draw(imd *imdraw.IMDraw) {
	var (
		bit  uint = 0
		vb   uint8
		addr = vramStart
	)
	for x := 0; x < screenW; x++ {
		for y := 0; y < screenH; y++ {
			// Read the next VRAM byte.
			if bit == 0 {
				vb = m.m.Read(addr)
				addr++
			}

			// Check if the pixel is lit.
			if (vb>>bit)&0x01 != 0x00 {
				m.pixel(imd, x, y)
			}

			// Move on to the next bit.
			bit++

			// Last bit, move on to the next byte.
			if bit == 8 {
				bit = 0
			}
		}
	}
}

// pixel draws a pixel to the draw object at the give co-ordinates.
//
// The pixel is scaled to a size determined by the scale factor.
func (m *Machine) pixel(imd *imdraw.IMDraw, x, y int) {
	x1 := float64(x * scaleFactor)
	y1 := float64(y * scaleFactor)
	imd.Push(
		pixel.V(x1, y1),
		pixel.V(x1+scaleFactor, y1+scaleFactor),
	)
	imd.Rectangle(0)
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
