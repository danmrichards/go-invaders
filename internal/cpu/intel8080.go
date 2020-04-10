package cpu

type Intel8080 struct {
	// Working "scratchpad" registers.
	B byte
	C byte
	D byte
	E byte
	H byte
	L byte

	// 8-bit accumulator.
	A byte

	// Stack pointer, stores address of last program request in the stack.
	SP uint16

	// Program counter, stores the address of the instruction being executed.
	PC uint16

	// Conditions represents the condition bits of the CPU.
	CC Conditions

	// Interrupts enabled (1 = enabled, 0 = disabled).
	IE byte
}

func NewIntel8080() *Intel8080 {
	return &Intel8080{
		CC: NewConditions(),
	}
}
