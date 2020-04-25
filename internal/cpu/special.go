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
