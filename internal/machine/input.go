package machine

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
)

// input returns input parsed from the given port.
func (m *Machine) input(port byte) byte {
	if m.debug {
		fmt.Printf("IN: %02x\n", port)
	}

	// TODO(dr): tweak the controls.

	var n byte
	switch port {
	case 0:
		// There technically is a port 0 in the Space Invaders hardware, but
		// it does not get used by the software.
	case 1:
		// Bit 3 is always 1.
		n |= 0x01 << 3

		// Credit.
		if m.w.Pressed(pixelgl.KeyEnter) {
			n |= 0x01
		}

		// 1P start.
		if m.w.Pressed(pixelgl.Key1) {
			n |= 0x01 << 2
		}

		// 2P start.
		if m.w.Pressed(pixelgl.Key2) {
			n |= 0x01 << 1
		}

		// 1P shot.
		if m.w.Pressed(pixelgl.KeyE) {
			n |= 0x01 << 4
		}

		// 1P left.
		if m.w.Pressed(pixelgl.KeyQ) {
			n |= 0x01 << 5
		}

		// 1P right.
		if m.w.Pressed(pixelgl.KeyW) {
			n |= 0x01 << 6
		}
	case 2:
		// 0 = 3 lives. 10 = 5 lives.
		n |= 0x00 << 0

		// 0 = extra life at 1500. 1 = extra life at 1000.
		n |= 0x00 << 3

		// Coin info on demo screen. 0 = ON.
		n |= 0x00 << 7

		// Tilt.
		if m.w.Pressed(pixelgl.KeyT) {
			n |= 0x01 << 2
		}

		// 2P shot.
		if m.w.Pressed(pixelgl.KeyI) {
			n |= 0x01 << 4
		}

		// 2P left.
		if m.w.Pressed(pixelgl.KeyO) {
			n |= 0x01 << 5
		}

		// 2P right.
		if m.w.Pressed(pixelgl.KeyP) {
			n |= 0x01 << 6
		}
	case 3:
		// Result of the shift register.
		n = uint8((m.sd >> (8 - m.so)) & 0xff)
	}

	return n
}
