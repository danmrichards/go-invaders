package machine

// stepper is the interface that wraps the basic Step method.
//
// Step emulates exactly one instruction on the CPU.
type stepper interface {
	Step() error
}

type interrupter interface {
	Interrupt(uint16)
}

type cycler interface {
	Cycles() uint32
}

type processor interface {
	stepper
	interrupter
	cycler
}
