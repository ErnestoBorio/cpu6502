package Cpu6502

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

type word = uint16
const lowByte = 0xF  // Bitmask for low byte of word
const signBit = 1<<7 // Bitmask for leftmost (sign) bit

var opcodeCycles = [0x100]byte {
//  0 1 2 3 4 5 6 7 8 9 A B C D E F
    0,6,2,8,3,3,5,5,3,2,2,2,4,4,6,6, //00
    2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7, //10
    6,6,2,8,3,3,5,5,4,2,2,2,4,4,6,6, //20
    2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7, //30
    6,6,2,8,3,3,5,5,3,2,2,2,3,4,6,6, //40
    2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7, //50
    6,6,2,8,3,3,5,5,4,2,2,2,5,4,6,6, //60
    2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7, //70
    2,6,2,6,3,3,3,3,2,2,2,2,4,4,4,4, //80
    2,6,2,6,4,4,4,4,2,5,2,5,5,5,5,5, //90
    2,6,2,6,3,3,3,3,2,2,2,2,4,4,4,4, //A0
    2,5,2,5,4,4,4,4,2,4,2,4,4,4,4,4, //B0
    2,6,2,8,3,3,5,5,2,2,2,2,4,4,6,6, //C0
    2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7, //D0
    2,6,3,8,3,3,5,5,2,2,2,2,4,4,6,6, //E0
    2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7} //F0