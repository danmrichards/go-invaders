package machine

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

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
	x1 := float64(x * m.sf)
	y1 := float64(y * m.sf)
	imd.Push(
		pixel.V(x1, y1),
		pixel.V(x1+float64(m.sf), y1+float64(m.sf)),
	)
	imd.Rectangle(0)
}
