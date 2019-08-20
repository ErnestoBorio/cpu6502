package cpu6502

type word = uint16

type Cpu struct {
	PC    word // Program counter
	Stack byte // Stack pointer
	A     byte // A register
	X     byte // X register
	Y     byte // Y register

	Status struct {
		Zero     bool
		Carry    bool
		Decimal  bool
		IntDis   bool
		Overflow bool
		Negative bool
	}

	cycles byte // Cycle count of the last executed instruction [1..7]

	readMemory  [0x10000]func(word) byte
	writeMemory [0x10000]func(word,byte)
}

func (cpu *Cpu) init() {
	cpu.Stack = 0xFD
	cpu.A = 0
	cpu.X = 0
	cpu.Y = 0
	cpu.Status.Zero = false
	cpu.Status.Carry = false
	cpu.Status.Decimal = false
	cpu.Status.IntDis = true
	cpu.Status.Overflow = false
	cpu.Status.Negative = false
}

func (cpu *Cpu) reset() {
	cpu.PC = word(cpu.readMemory[0xFFFC](0xFFFC))
	cpu.PC += (word(cpu.readMemory[0xFFFC](0xFFFC)) << 8)
}