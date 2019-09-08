package cpu6502

type word = uint16

// Models the 6502 CPU
type CPU struct {
	PC    word // Program counter
	Stack byte // Stack pointer
	A     byte // A register
	X     byte // X register
	Y     byte // Y register

	Status struct {
		Zero     bool
		Carry    bool
		Decimal  bool
		Overflow bool
		Negative bool
		NoInterrupt bool
	}

	cycles uint8 // Cycle count of the last executed instruction [1..7]

	// References to functions on the host system to access memory
	ReadMemory  func(word) byte
	WriteMemory func(word, byte)
}

// Initialize the state of the cpu as stated in:
// http://wiki.nesdev.com/w/index.php/CPU_power_up_state 
// https://www.c64-wiki.com/wiki/Reset_(Process)
func (cpu *CPU) Init() *CPU {
	cpu.Stack = 0xFD // because of a fake push of PC and flags
	cpu.A = 0
	cpu.X = 0
	cpu.Y = 0
	cpu.Status.Zero = false
	cpu.Status.Carry = false
	cpu.Status.Decimal = false
	cpu.Status.Overflow = false
	cpu.Status.Negative = false
	cpu.Status.NoInterrupt = true
	return cpu // to do CPU.Init().Reset()
}

// Jump to the address where the reset vector points to
func (cpu *CPU) Reset() {
	cpu.PC = cpu.getWord(0xFFFC)
}

// Trigger an external IRQ interrupt
func (cpu *CPU) IRQ() uint8{
	if ! cpu.Status.NoInterrupt {
		cpu.irq(false)
		return cpu.cycles
	}
	return 0 // 0 CPU cycles executed
}

// Trigger an external NMI interrupt
func (cpu *CPU) NMI() uint8{
	cpu.push( byte( cpu.PC >>8)) // PC's high byte
	cpu.push( byte( cpu.PC & lowByte)) // PC's low byte
	cpu.push(cpu.packStatus())
	// Marat Fayzullin and others clear the decimal mode here
	cpu.Status.NoInterrupt = true // TODO: Marat Fayzullin doesn't do this
	cpu.PC = cpu.getWord(0xFFFA) // Jump to NMI vector
	
	cpu.cycles = 7
	return cpu.cycles
}