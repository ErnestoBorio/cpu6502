package cpu6502

type word = uint16

// Models the 6502 CPU
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

	readMemory  [0x10000]MemoryReader
	writeMemory [0x10000]MemoryWriter
}

type MemoryReader = func(word) byte
type MemoryWriter = func(word, byte)

// Initialize the state of the cpu as stated in:
// http://wiki.nesdev.com/w/index.php/CPU_power_up_state
func (cpu *Cpu) Init() *Cpu {
	cpu.Stack = 0xFD
	cpu.A = 0
	cpu.X = 0
	cpu.Y = 0
	cpu.Status.Zero     = false
	cpu.Status.Carry    = false
	cpu.Status.IntDis   = true
	cpu.Status.Decimal  = false
	cpu.Status.Overflow = false
	cpu.Status.Negative = false
	return cpu
}

// Hooks a range of addresses to a host system memory reader function
func (cpu *Cpu) HookMemoryReader(adrBegin uint16, adrEnd uint16, callback MemoryReader) {
	for address := adrBegin;; address++ {
		cpu.readMemory[address] = callback
		// without this, if adrEnd == 0xFFFF address wraps around to 0 and enters an infinite loop
		if address == adrEnd {
			break
		}
	}
}

// Hooks a range of addresses to a host system memory writer function
func (cpu *Cpu) HookMemoryWriter(adrBegin uint16, adrEnd uint16, callback MemoryWriter) {
	for address := adrBegin;; address++ {
		cpu.writeMemory[address] = callback
		if address == adrEnd {
			break
		}
	}
}

// Jump to the address where the reset vector points to
func (cpu *Cpu) Reset() {
	cpu.PC = cpu.getWord(0xFFFC)
}

// Trigger an external IRQ interrupt
func (cpu *Cpu) IRQ() {
	if ! cpu.Status.IntDis {
		cpu.irq(false)
	}
}

// Trigger an external NMI interrupt
func (cpu *Cpu) NMI() {
	cpu.cycles = 7
	cpu.push( byte( cpu.PC >>8)) // PC's high byte
	cpu.push( byte( cpu.PC & lowByte)) // PC's low byte
	cpu.push(cpu.packStatus())
	// Marat Fayzullin and others clear the decimal mode here
	cpu.Status.IntDis = true // TODO: Marat Fayzullin doesn't do this
	cpu.PC = cpu.getWord(0xFFFA) // Jump to NMI vector
}