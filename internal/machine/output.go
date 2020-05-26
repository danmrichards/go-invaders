package machine

import "fmt"

func (m *Machine) output(b byte) {
	if m.debug {
		fmt.Printf("OUT: %02x\n", b)
	}

	// TODO: Output handler
}
