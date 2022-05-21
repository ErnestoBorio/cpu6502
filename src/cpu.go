package cpu6502

// CPU Models the 6502 CPU
type CPU struct {
	PC    uint16 // Program counter
	Stack byte   // Stack pointer
	A     byte   // A register
	X     byte   // X register
	Y     byte   // Y register

	Status struct {
		Zero        bool
		Carry       bool
		Decimal     bool
		Overflow    bool
		Negative    bool
		NoInterrupt bool
	}

	// Used temporarily inside CPU.Step() <Don't use>
	tmpCycles int

	// References to functions on the host system to access memory
	ReadMemory  func(uint16) byte
	WriteMemory func(uint16, byte)
}

// Init Initializes the state of the cpu as stated in:
// http://wiki.nesdev.com/w/index.php/CPU_power_up_state
// https://www.c64-wiki.com/wiki/Reset_(Process)
func (cpu *CPU) Init(read func(uint16) byte, write func(uint16, byte)) {
	cpu.Stack = 0xFD // because of a fake push of PC and flags on reset interrupt
	cpu.A = 0
	cpu.X = 0
	cpu.Y = 0
	cpu.Status.Zero = false
	cpu.Status.Carry = false
	cpu.Status.Decimal = false
	cpu.Status.Overflow = false
	cpu.Status.Negative = false
	cpu.Status.NoInterrupt = true
	cpu.ReadMemory = read
	cpu.WriteMemory = write
}

// Jump to the address where the reset vector points to
func (cpu *CPU) Reset() {
	cpu.Stack = 0xFD
	cpu.PC = cpu.getUint16(0xFFFC)
}

// Trigger an external IRQ interrupt
func (cpu *CPU) IRQ() {
	if !cpu.Status.NoInterrupt {
		cpu.irq(false)
	}
}

// NMI Triggers an external NMI interrupt
func (cpu *CPU) NMI() {
	cpu.push(byte(cpu.PC >> 8))   // PC's high byte
	cpu.push(byte(cpu.PC & 0xFF)) // PC's low byte
	cpu.push(cpu.packStatus())
	// @todo Marat Fayzullin and others clear the decimal mode here (NES specific?)
	cpu.Status.NoInterrupt = true  // @todo: Marat Fayzullin doesn't do this (NES specific?)
	cpu.PC = cpu.getUint16(0xFFFA) // Jump to NMI vector
	cpu.tmpCycles = 7
}

// getUint16 gets a little endian 16 bits value from 2 consecutive memory addresses
func (cpu *CPU) getUint16(address uint16) uint16 {
	value := uint16(cpu.ReadMemory(address))
	value |= uint16(cpu.ReadMemory(address+1)) << 8
	return value
}

// Fetches one complete instruction from PC and executes it.
func (cpu *CPU) Step() int {
	// Fetch operation code from current PC address
	opcode := cpu.ReadMemory(cpu.PC)
	// Advance PC for either first argument byte or next instruction
	cpu.PC++
	// Count basic instruction cycles, they can be incremented afterwards in certain conditions
	cpu.tmpCycles = Opcodes[opcode].Cycles
	
	// Call the appropriate instruction and addressing mode for the fetched opcode
	address := Opcodes[opcode].Addressing(cpu)
	Opcodes[opcode].Instruction(cpu, address)
	
	// @todo: byte following opcode is read in advance, possibly causing side effects. Test it.
	return cpu.tmpCycles
}