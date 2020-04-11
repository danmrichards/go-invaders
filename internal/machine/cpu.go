package machine

// cpuStepper is the interface that wraps the basic Step method.
//
// Step emulates exactly one instruction on the CPU.
type cpuStepper interface {
	Step() error
}
