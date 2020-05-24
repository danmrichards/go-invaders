package cpu

// out is the "Output" handler.
func (i *Intel8080) out() {
	if i.oh != nil {
		i.oh(i.immediateByte())
	}
}

// in is the "Input" handler.
func (i *Intel8080) in() {
	if i.ih != nil {
		i.ih(i.immediateByte())
	}
}
