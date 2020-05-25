package machine

// stepper is the interface that wraps the basic Step method.
//
// Step emulates exactly one instruction on the CPU.
type stepper interface {
	Step() error
}

// interrupter is the interface that wraps the basic Interrupt method.
//
// Interrupt instructs the CPU with a given interrupt address.
type interrupter interface {
	Interrupt(uint16)
}

// cycler is the interface that wraps the basic Cycles method.
//
// Cycles returns the current cycle count of the CPU.
type cycler interface {
	Cycles() uint32
}

// runner is the interface that wraps the basic Running method.
//
// Running returns true if the CPU is running.
type runner interface {
	Running() bool
}

// processor is the interface that implementations of a CPU are epxected to
// implement.
type processor interface {
	stepper
	interrupter
	cycler
	runner
}
