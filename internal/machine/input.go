package machine

import "fmt"

func (m *Machine) input(b byte) {
	if m.debug {
		fmt.Printf("IN: %02x\n", b)
	}

	// TODO: Input handler
}
