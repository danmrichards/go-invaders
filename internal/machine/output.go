package machine

import (
	"fmt"
)

// output handles output operations for the given port.
func (m *Machine) output(port byte) {
	if m.debug {
		fmt.Printf("OUT: %02x\n", port)
	}

	switch port {
	case 0x02:
		// Set the shift register offset.
		m.so = uint16(m.c.Accumulator()) & 0x07
	case 0x03:
		m.playSound(m.c.Accumulator(), 1)
	case 0x04:
		// Set the shift data.
		m.sd = uint16(m.c.Accumulator())<<8 | m.sd>>8
	case 0x05:
		m.playSound(m.c.Accumulator(), 2)
	case 0x06:
		m.wd = m.c.Accumulator()
	}
}

// playSound plays the sound indicated by the given data and sound bank.
func (m *Machine) playSound(data byte, bank int) {
	// Returns true if the given bit is set.
	bit := func(b byte, pos int) bool {
		return (b & (0x01 << pos)) != 0
	}

	switch {
	case bank == 1 && data != m.snd1:
		if bit(data, 0) && !bit(m.snd1, 0) {
			m.p.Play("0.wav")
		} else if bit(data, 1) && !bit(m.snd1, 1) {
			m.p.Play("1.wav")
		} else if bit(data, 2) && !bit(m.snd1, 2) {
			m.p.Play("2.wav")
		} else if bit(data, 3) && !bit(m.snd1, 3) {
			m.p.Play("3.wav")
		}

		m.snd1 = data
	case bank == 2 && data != m.snd2:
		if bit(data, 0) && !bit(m.snd2, 0) {
			m.p.Play("4.wav")
		} else if bit(data, 1) && !bit(m.snd2, 1) {
			m.p.Play("5.wav")
		} else if bit(data, 2) && !bit(m.snd2, 2) {
			m.p.Play("6.wav")
		} else if bit(data, 3) && !bit(m.snd2, 3) {
			m.p.Play("7.wav")
		} else if bit(data, 4) && !bit(m.snd2, 4) {
			m.p.Play("8.wav")
		}

		m.snd2 = data
	}
}
