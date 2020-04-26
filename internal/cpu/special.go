package cpu

// ei is the "enable interrupt" handler.
func (i *Intel8080) ei() uint16 {
	i.ie = true

	return defaultInstructionLen
}

// di is the "disable interrupt" handler.
func (i *Intel8080) di() uint16 {
	i.ie = false

	return defaultInstructionLen
}

// hlt is the "Halt" handler.
func (i *Intel8080) hlt() uint16 {
	i.halted = true

	return defaultInstructionLen
}

// stc is the "Set Carry" handler.
func (i *Intel8080) stc() uint16 {
	i.cc.cy = true

	return defaultInstructionLen
}

// cmc is the "Complement Carry" handler.
//
// If the Carry bit = 0, it is set to 1. If the Carry bit = 1, it is reset to O.
func (i *Intel8080) cmc() uint16 {
	i.cc.cy = !i.cc.cy

	return defaultInstructionLen
}
