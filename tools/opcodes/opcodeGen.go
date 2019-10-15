package main

import (
	"fmt"
)

/*
	Generates cpu6502/opcodeTable.go that holds the info needed to disassemble any cpu operation
*/

func main() {
	ops := [0x100]Opcode{}

	// Initialize all operations as undocumented
	for cod,_:= range ops {
		ops[cod] = Opcode {
			mnemonic: "",
			addressing: "",
			addressingFunc: "DbgNoAddressing", 
			documented: false,
		}
	}

	// Now overwrite documented operations with actual data
	for cod, op := range Opcodes {
		ops[cod] = op
	}

	fmt.Print(
`package cpu6502

type Operation struct {
	Opcode      byte
	Mnemonic    string
	Length      uint8
	Cycles      uint8
	Documented  bool
	Addressing  string
	AddressFunc func(*CPU, uint16) uint16
}

// Only BRK (00) operation has no cycles here, because they're accounted for in the IRQ interrupt handler
var Opcodes = [0x100] Operation {
`)

	for cod,op := range ops {
		mnemPad := ""
		if op.mnemonic == ""{
			mnemPad = "   "
		}
		docPad := ""
		if op.documented {
			docPad = " "
		}
		addPad := ""
		if op.addressing == "" {
			addPad = "   "
		}
		
		fmt.Printf("\t{ Opcode: 0x%02X, Mnemonic: \"%s\"%s, Length: %d, Cycles: %d, Documented: %v,%s Addressing: \"%s\"%s, AddressFunc: (*CPU).%s },\n", 
			cod, op.mnemonic, mnemPad, op.length, opcodeCycles[cod], op.documented, docPad, op.addressing, addPad, op.addressingFunc)
	}
	fmt.Println("}")
}

var opcodeCycles = [0x100]byte{
//  0  1  2  3  4  5  6  7  8  9  A  B  C  D  E  F
	0, 6, 2, 8, 3, 3, 5, 5, 3, 2, 2, 2, 4, 4, 6, 6, //00
	2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 7, 7, //10
	6, 6, 2, 8, 3, 3, 5, 5, 4, 2, 2, 2, 4, 4, 6, 6, //20
	2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 7, 7, //30
	6, 6, 2, 8, 3, 3, 5, 5, 3, 2, 2, 2, 3, 4, 6, 6, //40
	2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 7, 7, //50
	6, 6, 2, 8, 3, 3, 5, 5, 4, 2, 2, 2, 5, 4, 6, 6, //60
	2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 7, 7, //70
	2, 6, 2, 6, 3, 3, 3, 3, 2, 2, 2, 2, 4, 4, 4, 4, //80
	2, 6, 2, 6, 4, 4, 4, 4, 2, 5, 2, 5, 5, 5, 5, 5, //90
	2, 6, 2, 6, 3, 3, 3, 3, 2, 2, 2, 2, 4, 4, 4, 4, //A0
	2, 5, 2, 5, 4, 4, 4, 4, 2, 4, 2, 4, 4, 4, 4, 4, //B0
	2, 6, 2, 8, 3, 3, 5, 5, 2, 2, 2, 2, 4, 4, 6, 6, //C0
	2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 7, 7, //D0
	2, 6, 3, 8, 3, 3, 5, 5, 2, 2, 2, 2, 4, 4, 6, 6, //E0
	2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 7, 7, //F0
}

type Opcode struct {
	mnemonic string
	length int
	addressing string
	addressingFunc string
	documented bool
}

var Opcodes map[byte]Opcode = map[byte]Opcode {
	0xA9: Opcode{mnemonic: "LDA", length: 2, addressing: "IMM", addressingFunc: "DbgImmediate",    documented: true},
	0xA5: Opcode{mnemonic: "LDA", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0xB5: Opcode{mnemonic: "LDA", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0xAD: Opcode{mnemonic: "LDA", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0xBD: Opcode{mnemonic: "LDA", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0xB9: Opcode{mnemonic: "LDA", length: 3, addressing: "ABY", addressingFunc: "DbgAbsoluteY",    documented: true},
	0xA1: Opcode{mnemonic: "LDA", length: 2, addressing: "IIX", addressingFunc: "DbgIndexedIndirectX", documented: true},
	0xB1: Opcode{mnemonic: "LDA", length: 2, addressing: "IIY", addressingFunc: "DbgIndirectIndexedY", documented: true},
	0xA2: Opcode{mnemonic: "LDX", length: 2, addressing: "IMM", addressingFunc: "DbgImmediate",    documented: true},
	0xA6: Opcode{mnemonic: "LDX", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0xB6: Opcode{mnemonic: "LDX", length: 2, addressing: "ZPY", addressingFunc: "DbgZeroPageY",    documented: true},
	0xAE: Opcode{mnemonic: "LDX", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0xBE: Opcode{mnemonic: "LDX", length: 3, addressing: "ABY", addressingFunc: "DbgAbsoluteY",    documented: true},
	0xA0: Opcode{mnemonic: "LDY", length: 2, addressing: "IMM", addressingFunc: "DbgImmediate",    documented: true},
	0xA4: Opcode{mnemonic: "LDY", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0xB4: Opcode{mnemonic: "LDY", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0xAC: Opcode{mnemonic: "LDY", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0xBC: Opcode{mnemonic: "LDY", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0x85: Opcode{mnemonic: "STA", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0x95: Opcode{mnemonic: "STA", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0x8D: Opcode{mnemonic: "STA", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x9D: Opcode{mnemonic: "STA", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0x99: Opcode{mnemonic: "STA", length: 3, addressing: "ABY", addressingFunc: "DbgAbsoluteY",    documented: true},
	0x81: Opcode{mnemonic: "STA", length: 2, addressing: "IIX", addressingFunc: "DbgIndexedIndirectX", documented: true},
	0x91: Opcode{mnemonic: "STA", length: 2, addressing: "IIY", addressingFunc: "DbgIndirectIndexedY", documented: true},
	0x86: Opcode{mnemonic: "STX", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0x96: Opcode{mnemonic: "STX", length: 2, addressing: "ZPY", addressingFunc: "DbgZeroPageY",    documented: true},
	0x8E: Opcode{mnemonic: "STX", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x84: Opcode{mnemonic: "STY", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0x94: Opcode{mnemonic: "STY", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0x8C: Opcode{mnemonic: "STY", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0xAA: Opcode{mnemonic: "TAX", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0xA8: Opcode{mnemonic: "TAY", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x8A: Opcode{mnemonic: "TXA", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x98: Opcode{mnemonic: "TYA", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0xBA: Opcode{mnemonic: "TSX", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x9A: Opcode{mnemonic: "TXS", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x48: Opcode{mnemonic: "PHA", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x08: Opcode{mnemonic: "PHP", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x68: Opcode{mnemonic: "PLA", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x28: Opcode{mnemonic: "PLP", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x29: Opcode{mnemonic: "AND", length: 2, addressing: "IMM", addressingFunc: "DbgImmediate",    documented: true},
	0x25: Opcode{mnemonic: "AND", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0x35: Opcode{mnemonic: "AND", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0x2D: Opcode{mnemonic: "AND", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x3D: Opcode{mnemonic: "AND", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0x39: Opcode{mnemonic: "AND", length: 3, addressing: "ABY", addressingFunc: "DbgAbsoluteY",    documented: true},
	0x21: Opcode{mnemonic: "AND", length: 2, addressing: "IIX", addressingFunc: "DbgIndexedIndirectX", documented: true},
	0x31: Opcode{mnemonic: "AND", length: 2, addressing: "IIY", addressingFunc: "DbgIndirectIndexedY", documented: true},
	0x49: Opcode{mnemonic: "EOR", length: 2, addressing: "IMM", addressingFunc: "DbgImmediate",    documented: true},
	0x45: Opcode{mnemonic: "EOR", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0x55: Opcode{mnemonic: "EOR", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0x4D: Opcode{mnemonic: "EOR", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x5D: Opcode{mnemonic: "EOR", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0x59: Opcode{mnemonic: "EOR", length: 3, addressing: "ABY", addressingFunc: "DbgAbsoluteY",    documented: true},
	0x41: Opcode{mnemonic: "EOR", length: 2, addressing: "IIX", addressingFunc: "DbgIndexedIndirectX", documented: true},
	0x51: Opcode{mnemonic: "EOR", length: 2, addressing: "IIY", addressingFunc: "DbgIndirectIndexedY", documented: true},
	0x09: Opcode{mnemonic: "ORA", length: 2, addressing: "IMM", addressingFunc: "DbgImmediate",    documented: true},
	0x05: Opcode{mnemonic: "ORA", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0x15: Opcode{mnemonic: "ORA", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0x0D: Opcode{mnemonic: "ORA", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x1D: Opcode{mnemonic: "ORA", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0x19: Opcode{mnemonic: "ORA", length: 3, addressing: "ABY", addressingFunc: "DbgAbsoluteY",    documented: true},
	0x01: Opcode{mnemonic: "ORA", length: 2, addressing: "IIX", addressingFunc: "DbgIndexedIndirectX", documented: true},
	0x11: Opcode{mnemonic: "ORA", length: 2, addressing: "IIY", addressingFunc: "DbgIndirectIndexedY", documented: true},
	0x24: Opcode{mnemonic: "BIT", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0x2C: Opcode{mnemonic: "BIT", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x69: Opcode{mnemonic: "ADC", length: 2, addressing: "IMM", addressingFunc: "DbgImmediate",    documented: true},
	0x65: Opcode{mnemonic: "ADC", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0x75: Opcode{mnemonic: "ADC", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0x6D: Opcode{mnemonic: "ADC", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x7D: Opcode{mnemonic: "ADC", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0x79: Opcode{mnemonic: "ADC", length: 3, addressing: "ABY", addressingFunc: "DbgAbsoluteY",    documented: true},
	0x61: Opcode{mnemonic: "ADC", length: 2, addressing: "IIX", addressingFunc: "DbgIndexedIndirectX", documented: true},
	0x71: Opcode{mnemonic: "ADC", length: 2, addressing: "IIY", addressingFunc: "DbgIndirectIndexedY", documented: true},
	0xE9: Opcode{mnemonic: "SBC", length: 2, addressing: "IMM", addressingFunc: "DbgImmediate",    documented: true},
	0xE5: Opcode{mnemonic: "SBC", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0xF5: Opcode{mnemonic: "SBC", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0xED: Opcode{mnemonic: "SBC", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0xFD: Opcode{mnemonic: "SBC", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0xF9: Opcode{mnemonic: "SBC", length: 3, addressing: "ABY", addressingFunc: "DbgAbsoluteY",    documented: true},
	0xE1: Opcode{mnemonic: "SBC", length: 2, addressing: "IIX", addressingFunc: "DbgIndexedIndirectX", documented: true},
	0xF1: Opcode{mnemonic: "SBC", length: 2, addressing: "IIY", addressingFunc: "DbgIndirectIndexedY", documented: true},
	0xC9: Opcode{mnemonic: "CMP", length: 2, addressing: "IMM", addressingFunc: "DbgImmediate",    documented: true},
	0xC5: Opcode{mnemonic: "CMP", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0xD5: Opcode{mnemonic: "CMP", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0xCD: Opcode{mnemonic: "CMP", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0xDD: Opcode{mnemonic: "CMP", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0xD9: Opcode{mnemonic: "CMP", length: 3, addressing: "ABY", addressingFunc: "DbgAbsoluteY",    documented: true},
	0xC1: Opcode{mnemonic: "CMP", length: 2, addressing: "IIX", addressingFunc: "DbgIndexedIndirectX", documented: true},
	0xD1: Opcode{mnemonic: "CMP", length: 2, addressing: "IIY", addressingFunc: "DbgIndirectIndexedY", documented: true},
	0xE0: Opcode{mnemonic: "CPX", length: 2, addressing: "IMM", addressingFunc: "DbgImmediate",    documented: true},
	0xE4: Opcode{mnemonic: "CPX", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0xEC: Opcode{mnemonic: "CPX", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0xC0: Opcode{mnemonic: "CPY", length: 2, addressing: "IMM", addressingFunc: "DbgImmediate",    documented: true},
	0xC4: Opcode{mnemonic: "CPY", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0xCC: Opcode{mnemonic: "CPY", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0xE6: Opcode{mnemonic: "INC", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0xF6: Opcode{mnemonic: "INC", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0xEE: Opcode{mnemonic: "INC", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0xFE: Opcode{mnemonic: "INC", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0xE8: Opcode{mnemonic: "INX", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0xC8: Opcode{mnemonic: "INY", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0xC6: Opcode{mnemonic: "DEC", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0xD6: Opcode{mnemonic: "DEC", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0xCE: Opcode{mnemonic: "DEC", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0xDE: Opcode{mnemonic: "DEC", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0xCA: Opcode{mnemonic: "DEX", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x88: Opcode{mnemonic: "DEY", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x0A: Opcode{mnemonic: "ASL", length: 1, addressing: "ACU", addressingFunc: "DbgNoAddressing", documented: true},
	0x4A: Opcode{mnemonic: "LSR", length: 1, addressing: "ACU", addressingFunc: "DbgNoAddressing", documented: true},
	0x46: Opcode{mnemonic: "LSR", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0x56: Opcode{mnemonic: "LSR", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0x4E: Opcode{mnemonic: "LSR", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x5E: Opcode{mnemonic: "LSR", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0x2A: Opcode{mnemonic: "ROL", length: 1, addressing: "ACU", addressingFunc: "DbgNoAddressing", documented: true},
	0x26: Opcode{mnemonic: "ROL", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0x36: Opcode{mnemonic: "ROL", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0x2E: Opcode{mnemonic: "ROL", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x3E: Opcode{mnemonic: "ROL", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0x6A: Opcode{mnemonic: "ROR", length: 1, addressing: "ACU", addressingFunc: "DbgNoAddressing", documented: true},
	0x66: Opcode{mnemonic: "ROR", length: 2, addressing: "ZRP", addressingFunc: "DbgZeroPage",     documented: true},
	0x76: Opcode{mnemonic: "ROR", length: 2, addressing: "ZPX", addressingFunc: "DbgZeroPageX",    documented: true},
	0x6E: Opcode{mnemonic: "ROR", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x7E: Opcode{mnemonic: "ROR", length: 3, addressing: "ABX", addressingFunc: "DbgAbsoluteX",    documented: true},
	0x4C: Opcode{mnemonic: "JMP", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x6C: Opcode{mnemonic: "JMP", length: 3, addressing: "IND", addressingFunc: "DbgIndirect",     documented: true},
	0x20: Opcode{mnemonic: "JSR", length: 3, addressing: "ABS", addressingFunc: "DbgAbsolute",     documented: true},
	0x60: Opcode{mnemonic: "RTS", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x90: Opcode{mnemonic: "BCC", length: 2, addressing: "REL", addressingFunc: "DbgImmediate",    documented: true},
	0xB0: Opcode{mnemonic: "BCS", length: 2, addressing: "REL", addressingFunc: "DbgImmediate",    documented: true},
	0xF0: Opcode{mnemonic: "BEQ", length: 2, addressing: "REL", addressingFunc: "DbgImmediate",    documented: true},
	0x30: Opcode{mnemonic: "BMI", length: 2, addressing: "REL", addressingFunc: "DbgImmediate",    documented: true},
	0xD0: Opcode{mnemonic: "BNE", length: 2, addressing: "REL", addressingFunc: "DbgImmediate",    documented: true},
	0x10: Opcode{mnemonic: "BPL", length: 2, addressing: "REL", addressingFunc: "DbgImmediate",    documented: true},
	0x50: Opcode{mnemonic: "BVC", length: 2, addressing: "REL", addressingFunc: "DbgImmediate",    documented: true},
	0x70: Opcode{mnemonic: "BVS", length: 2, addressing: "REL", addressingFunc: "DbgImmediate",    documented: true},
	0x18: Opcode{mnemonic: "CLC", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0xD8: Opcode{mnemonic: "CLD", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x58: Opcode{mnemonic: "CLI", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0xB8: Opcode{mnemonic: "CLV", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x38: Opcode{mnemonic: "SEC", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0xF8: Opcode{mnemonic: "SED", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x78: Opcode{mnemonic: "SEI", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x00: Opcode{mnemonic: "BRK", length: 2, addressing:   "", addressingFunc: "DbgNoAddressing", documented: true},
	0xEA: Opcode{mnemonic: "NOP", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x40: Opcode{mnemonic: "RTI", length: 1, addressing: "IMP", addressingFunc: "DbgNoAddressing", documented: true},
	0x1A: Opcode{mnemonic: "NO1", length: 1, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x3A: Opcode{mnemonic: "NO1", length: 1, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x5A: Opcode{mnemonic: "NO1", length: 1, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x7A: Opcode{mnemonic: "NO1", length: 1, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0xDA: Opcode{mnemonic: "NO1", length: 1, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0xFA: Opcode{mnemonic: "NO1", length: 1, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x80: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x82: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x89: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0xC2: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0xE2: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x04: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x14: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x34: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x44: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x54: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x64: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x74: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0xD4: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0xF4: Opcode{mnemonic: "NO2", length: 2, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x0C: Opcode{mnemonic: "NO3", length: 3, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x1C: Opcode{mnemonic: "NO3", length: 3, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x3C: Opcode{mnemonic: "NO3", length: 3, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x5C: Opcode{mnemonic: "NO3", length: 3, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0x7C: Opcode{mnemonic: "NO3", length: 3, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0xDC: Opcode{mnemonic: "NO3", length: 3, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
	0xFC: Opcode{mnemonic: "NO3", length: 3, addressing:    "", addressingFunc: "DbgNoAddressing", documented: false},
}