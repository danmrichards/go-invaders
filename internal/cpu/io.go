package cpu

// out is the "Output" handler.
//
// The contents of the accumulator are sent to output device number exp.
func (i *Intel8080) out() uint16 {
	// TODO: How do we output to a device?
	return defaultInstructionLen
}

// in is the "Input" handler.
//
// An eight-bit data byte is read from input device number expand replaces the
// contents of the accumulator.
func (i *Intel8080) in() uint16 {
	// TODO: How do we get input from a device?
	return defaultInstructionLen
}
