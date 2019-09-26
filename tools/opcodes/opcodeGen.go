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
			code: byte(cod),
			mnemonic: "",
			codeStr:  fmt.Sprintf("%2X", cod),
			addrLabel: "",
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
	Opcode      uint8
	Mnemonic    string
	Cycles      uint8
	Documented  bool
	Addressing  func(*CPU) uint16
}

// Only BRK (00) operation has no cycles here, because they're accounted for in the IRQ interrupt handler
var Opcodes [0x100] Operation = [0x100] Operation {
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
		fmt.Printf("\t{ Opcode: 0x%02X, Mnemonic: \"%s\"%s, Cycles: %d, Documented: %v,%s Addressing: (*CPU).%s },\n", 
			cod, op.mnemonic, mnemPad, opcodeCycles[cod], op.documented, docPad, op.addressingFunc)
	}
	fmt.Println("}")
}



type Opcode struct {
	code byte
	mnemonic string
	codeStr string
	addrLabel string
	addressingFunc string
	addressing string
	documented bool
}

var opcodeCycles = [0x100] byte {
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
	2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7, //F0
}

var Opcodes map[byte]Opcode = map[byte]Opcode {
	0xA9: Opcode{0xA9, "LDA", "A9", "IMM", "DbgImmediate", "Immediate", true},
	0xA5: Opcode{0xA5, "LDA", "A5", "ZRP", "DbgZeroPage", "Zero page", true},
	0xB5: Opcode{0xB5, "LDA", "B5", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0xAD: Opcode{0xAD, "LDA", "AD", "ABS", "DbgAbsolute", "Absolute", true},
	0xBD: Opcode{0xBD, "LDA", "BD", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0xB9: Opcode{0xB9, "LDA", "B9", "ABY", "DbgAbsoluteY", "Absolute Y", true},
	0xA1: Opcode{0xA1, "LDA", "A1", "IIX", "DbgIndexedIndirectX", "Indexed Indirect X", true},
	0xB1: Opcode{0xB1, "LDA", "B1", "IIY", "DbgIndirectIndexedY", "Indirect Indexed Y", true},
	0xA2: Opcode{0xA2, "LDX", "A2", "IMM", "DbgImmediate", "Immediate", true},
	0xA6: Opcode{0xA6, "LDX", "A6", "ZRP", "DbgZeroPage", "Zero page", true},
	0xB6: Opcode{0xB6, "LDX", "B6", "ZPY", "DbgZeroPageY", "Zero page Y", true},
	0xAE: Opcode{0xAE, "LDX", "AE", "ABS", "DbgAbsolute", "Absolute", true},
	0xBE: Opcode{0xBE, "LDX", "BE", "ABY", "DbgAbsoluteY", "Absolute Y", true},
	0xA0: Opcode{0xA0, "LDY", "A0", "IMM", "DbgImmediate", "Immediate", true},
	0xA4: Opcode{0xA4, "LDY", "A4", "ZRP", "DbgZeroPage", "Zero page", true},
	0xB4: Opcode{0xB4, "LDY", "B4", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0xAC: Opcode{0xAC, "LDY", "AC", "ABS", "DbgAbsolute", "Absolute", true},
	0xBC: Opcode{0xBC, "LDY", "BC", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0x85: Opcode{0x85, "STA", "85", "ZRP", "DbgZeroPage", "Zero page", true},
	0x95: Opcode{0x95, "STA", "95", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0x8D: Opcode{0x8D, "STA", "8D", "ABS", "DbgAbsolute", "Absolute", true},
	0x9D: Opcode{0x9D, "STA", "9D", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0x99: Opcode{0x99, "STA", "99", "ABY", "DbgAbsoluteY", "Absolute Y", true},
	0x81: Opcode{0x81, "STA", "81", "IIX", "DbgIndexedIndirectX", "Indexed Indirect X", true},
	0x91: Opcode{0x91, "STA", "91", "IIY", "DbgIndirectIndexedY", "Indirect Indexed Y", true},
	0x86: Opcode{0x86, "STX", "86", "ZRP", "DbgZeroPage", "Zero page", true},
	0x96: Opcode{0x96, "STX", "96", "ZPY", "DbgZeroPageY", "Zero page Y", true},
	0x8E: Opcode{0x8E, "STX", "8E", "ABS", "DbgAbsolute", "Absolute", true},
	0x84: Opcode{0x84, "STY", "84", "ZRP", "DbgZeroPage", "Zero page", true},
	0x94: Opcode{0x94, "STY", "94", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0x8C: Opcode{0x8C, "STY", "8C", "ABS", "DbgAbsolute", "Absolute", true},
	0xAA: Opcode{0xAA, "TAX", "AA", "IMM", "DbgImmediate", "Immediate", true},
	0xA8: Opcode{0xA8, "TAY", "A8", "IMM", "DbgImmediate", "Immediate", true},
	0x8A: Opcode{0x8A, "TXA", "8A", "IMM", "DbgImmediate", "Immediate", true},
	0x98: Opcode{0x98, "TYA", "98", "IMM", "DbgImmediate", "Immediate", true},
	0xBA: Opcode{0xBA, "TSX", "BA", "IMM", "DbgImmediate", "Immediate", true},
	0x9A: Opcode{0x9A, "TXS", "9A", "IMM", "DbgImmediate", "Immediate", true},
	0x48: Opcode{0x48, "PHA", "48", "IMP", "DbgNoAddressing", "Implied", true},
	0x08: Opcode{0x08, "PHP", "08", "IMP", "DbgNoAddressing", "Implied", true},
	0x68: Opcode{0x68, "PLA", "68", "IMP", "DbgNoAddressing", "Implied", true},
	0x28: Opcode{0x28, "PLP", "28", "IMP", "DbgNoAddressing", "Implied", true},
	0x29: Opcode{0x29, "AND", "29", "IMM", "DbgImmediate", "Immediate", true},
	0x25: Opcode{0x25, "AND", "25", "ZRP", "DbgZeroPage", "Zero page", true},
	0x35: Opcode{0x35, "AND", "35", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0x2D: Opcode{0x2D, "AND", "2D", "ABS", "DbgAbsolute", "Absolute", true},
	0x3D: Opcode{0x3D, "AND", "3D", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0x39: Opcode{0x39, "AND", "39", "ABY", "DbgAbsoluteY", "Absolute Y", true},
	0x21: Opcode{0x21, "AND", "21", "IIX", "DbgIndexedIndirectX", "Indexed Indirect X", true},
	0x31: Opcode{0x31, "AND", "31", "IIY", "DbgIndirectIndexedY", "Indirect Indexed Y", true},
	0x49: Opcode{0x49, "EOR", "49", "IMM", "DbgImmediate", "Immediate", true},
	0x45: Opcode{0x45, "EOR", "45", "ZRP", "DbgZeroPage", "Zero page", true},
	0x55: Opcode{0x55, "EOR", "55", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0x4D: Opcode{0x4D, "EOR", "4D", "ABS", "DbgAbsolute", "Absolute", true},
	0x5D: Opcode{0x5D, "EOR", "5D", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0x59: Opcode{0x59, "EOR", "59", "ABY", "DbgAbsoluteY", "Absolute Y", true},
	0x41: Opcode{0x41, "EOR", "41", "IIX", "DbgIndexedIndirectX", "Indexed Indirect X", true},
	0x51: Opcode{0x51, "EOR", "51", "IIY", "DbgIndirectIndexedY", "Indirect Indexed Y", true},
	0x09: Opcode{0x09, "ORA", "09", "IMM", "DbgImmediate", "Immediate", true},
	0x05: Opcode{0x05, "ORA", "05", "ZRP", "DbgZeroPage", "Zero page", true},
	0x15: Opcode{0x15, "ORA", "15", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0x0D: Opcode{0x0D, "ORA", "0D", "ABS", "DbgAbsolute", "Absolute", true},
	0x1D: Opcode{0x1D, "ORA", "1D", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0x19: Opcode{0x19, "ORA", "19", "ABY", "DbgAbsoluteY", "Absolute Y", true},
	0x01: Opcode{0x01, "ORA", "01", "IIX", "DbgIndexedIndirectX", "Indexed Indirect X", true},
	0x11: Opcode{0x11, "ORA", "11", "IIY", "DbgIndirectIndexedY", "Indirect Indexed Y", true},
	0x24: Opcode{0x24, "BIT", "24", "ZRP", "DbgZeroPage", "Zero page", true},
	0x2C: Opcode{0x2C, "BIT", "2C", "ABS", "DbgAbsolute", "Absolute", true},
	0x69: Opcode{0x69, "ADC", "69", "IMM", "DbgImmediate", "Immediate", true},
	0x65: Opcode{0x65, "ADC", "65", "ZRP", "DbgZeroPage", "Zero page", true},
	0x75: Opcode{0x75, "ADC", "75", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0x6D: Opcode{0x6D, "ADC", "6D", "ABS", "DbgAbsolute", "Absolute", true},
	0x7D: Opcode{0x7D, "ADC", "7D", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0x79: Opcode{0x79, "ADC", "79", "ABY", "DbgAbsoluteY", "Absolute Y", true},
	0x61: Opcode{0x61, "ADC", "61", "IIX", "DbgIndexedIndirectX", "Indexed Indirect X", true},
	0x71: Opcode{0x71, "ADC", "71", "IIY", "DbgIndirectIndexedY", "Indirect Indexed Y", true},
	0xE9: Opcode{0xE9, "SBC", "E9", "IMM", "DbgImmediate", "Immediate", true},
	0xE5: Opcode{0xE5, "SBC", "E5", "ZRP", "DbgZeroPage", "Zero page", true},
	0xF5: Opcode{0xF5, "SBC", "F5", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0xED: Opcode{0xED, "SBC", "ED", "ABS", "DbgAbsolute", "Absolute", true},
	0xFD: Opcode{0xFD, "SBC", "FD", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0xF9: Opcode{0xF9, "SBC", "F9", "ABY", "DbgAbsoluteY", "Absolute Y", true},
	0xE1: Opcode{0xE1, "SBC", "E1", "IIX", "DbgIndexedIndirectX", "Indexed Indirect X", true},
	0xF1: Opcode{0xF1, "SBC", "F1", "IIY", "DbgIndirectIndexedY", "Indirect Indexed Y", true},
	0xC9: Opcode{0xC9, "CMP", "C9", "IMM", "DbgImmediate", "Immediate", true},
	0xC5: Opcode{0xC5, "CMP", "C5", "ZRP", "DbgZeroPage", "Zero page", true},
	0xD5: Opcode{0xD5, "CMP", "D5", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0xCD: Opcode{0xCD, "CMP", "CD", "ABS", "DbgAbsolute", "Absolute", true},
	0xDD: Opcode{0xDD, "CMP", "DD", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0xD9: Opcode{0xD9, "CMP", "D9", "ABY", "DbgAbsoluteY", "Absolute Y", true},
	0xC1: Opcode{0xC1, "CMP", "C1", "IIX", "DbgIndexedIndirectX", "Indexed Indirect X", true},
	0xD1: Opcode{0xD1, "CMP", "D1", "IIY", "DbgIndirectIndexedY", "Indirect Indexed Y", true},
	0xE0: Opcode{0xE0, "CPX", "E0", "IMM", "DbgImmediate", "Immediate", true},
	0xE4: Opcode{0xE4, "CPX", "E4", "ZRP", "DbgZeroPage", "Zero page", true},
	0xEC: Opcode{0xEC, "CPX", "EC", "ABS", "DbgAbsolute", "Absolute", true},
	0xC0: Opcode{0xC0, "CPY", "C0", "IMM", "DbgImmediate", "Immediate", true},
	0xC4: Opcode{0xC4, "CPY", "C4", "ZRP", "DbgZeroPage", "Zero page", true},
	0xCC: Opcode{0xCC, "CPY", "CC", "ABS", "DbgAbsolute", "Absolute", true},
	0xE6: Opcode{0xE6, "INC", "E6", "ZRP", "DbgZeroPage", "Zero page", true},
	0xF6: Opcode{0xF6, "INC", "F6", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0xEE: Opcode{0xEE, "INC", "EE", "ABS", "DbgAbsolute", "Absolute", true},
	0xFE: Opcode{0xFE, "INC", "FE", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0xE8: Opcode{0xE8, "INX", "E8", "IMP", "DbgNoAddressing", "Implied", true},
	0xC8: Opcode{0xC8, "INY", "C8", "IMP", "DbgNoAddressing", "Implied", true},
	0xC6: Opcode{0xC6, "DEC", "C6", "ZRP", "DbgZeroPage", "Zero page", true},
	0xD6: Opcode{0xD6, "DEC", "D6", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0xCE: Opcode{0xCE, "DEC", "CE", "ABS", "DbgAbsolute", "Absolute", true},
	0xDE: Opcode{0xDE, "DEC", "DE", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0xCA: Opcode{0xCA, "DEX", "CA", "IMP", "DbgNoAddressing", "Implied", true},
	0x88: Opcode{0x88, "DEY", "88", "IMP", "DbgNoAddressing", "Implied", true},
	0x0A: Opcode{0x0A, "ASL", "0A", "ACU", "DbgNoAddressing", "Accumulator", true},
	0x4A: Opcode{0x4A, "LSR", "4A", "ACU", "DbgNoAddressing", "Accumulator", true},
	0x46: Opcode{0x46, "LSR", "46", "ZRP", "DbgZeroPage", "Zero page", true},
	0x56: Opcode{0x56, "LSR", "56", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0x4E: Opcode{0x4E, "LSR", "4E", "ABS", "DbgAbsolute", "Absolute", true},
	0x5E: Opcode{0x5E, "LSR", "5E", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0x2A: Opcode{0x2A, "ROL", "2A", "ACU", "DbgNoAddressing", "Accumulator", true},
	0x26: Opcode{0x26, "ROL", "26", "ZRP", "DbgZeroPage", "Zero page", true},
	0x36: Opcode{0x36, "ROL", "36", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0x2E: Opcode{0x2E, "ROL", "2E", "ABS", "DbgAbsolute", "Absolute", true},
	0x3E: Opcode{0x3E, "ROL", "3E", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0x6A: Opcode{0x6A, "ROR", "6A", "ACU", "DbgNoAddressing", "Accumulator", true},
	0x66: Opcode{0x66, "ROR", "66", "ZRP", "DbgZeroPage", "Zero page", true},
	0x76: Opcode{0x76, "ROR", "76", "ZPX", "DbgZeroPageX", "Zero page X", true},
	0x6E: Opcode{0x6E, "ROR", "6E", "ABS", "DbgAbsolute", "Absolute", true},
	0x7E: Opcode{0x7E, "ROR", "7E", "ABX", "DbgAbsoluteX", "Absolute X", true},
	0x4C: Opcode{0x4C, "JMP", "4C", "ABS", "DbgAbsolute", "Absolute", true},
	0x6C: Opcode{0x6C, "JMP", "6C", "IND", "DbgIndirect", "Indirect", true},
	0x20: Opcode{0x20, "JSR", "20", "IMP", "DbgNoAddressing", "Implied", true},
	0x60: Opcode{0x60, "RTS", "60", "IMP", "DbgNoAddressing", "Implied", true},
	0x90: Opcode{0x90, "BCC", "90", "REL", "DbgNoAddressing", "Relative", true},
	0xB0: Opcode{0xB0, "BCS", "B0", "REL", "DbgNoAddressing", "Relative", true},
	0xF0: Opcode{0xF0, "BEQ", "F0", "REL", "DbgNoAddressing", "Relative", true},
	0x30: Opcode{0x30, "BMI", "30", "REL", "DbgNoAddressing", "Relative", true},
	0xD0: Opcode{0xD0, "BNE", "D0", "REL", "DbgNoAddressing", "Relative", true},
	0x10: Opcode{0x10, "BPL", "10", "REL", "DbgNoAddressing", "Relative", true},
	0x50: Opcode{0x50, "BVC", "50", "REL", "DbgNoAddressing", "Relative", true},
	0x70: Opcode{0x70, "BVS", "70", "REL", "DbgNoAddressing", "Relative", true},
	0x18: Opcode{0x18, "CLC", "18", "IMP", "DbgNoAddressing", "Implied", true},
	0xD8: Opcode{0xD8, "CLD", "D8", "IMP", "DbgNoAddressing", "Implied", true},
	0x58: Opcode{0x58, "CLI", "58", "IMP", "DbgNoAddressing", "Implied", true},
	0xB8: Opcode{0xB8, "CLV", "B8", "IMP", "DbgNoAddressing", "Implied", true},
	0x38: Opcode{0x38, "SEC", "38", "IMP", "DbgNoAddressing", "Implied", true},
	0xF8: Opcode{0xF8, "SED", "F8", "IMP", "DbgNoAddressing", "Implied", true},
	0x78: Opcode{0x78, "SEI", "78", "IMP", "DbgNoAddressing", "Implied", true},
	0x00: Opcode{0x00, "BRK", "00", "IMP", "DbgNoAddressing", "Implied", true},
	0xEA: Opcode{0xEA, "NOP", "EA", "IMP", "DbgNoAddressing", "Implied", true},
	0x40: Opcode{0x40, "RTI", "40", "IMP", "DbgNoAddressing", "Implied", true},
	0x1A: Opcode{mnemonic: "NO1", addressingFunc: "DbgNoAddressing", documented: false},
	0x3A: Opcode{mnemonic: "NO1", addressingFunc: "DbgNoAddressing", documented: false},
	0x5A: Opcode{mnemonic: "NO1", addressingFunc: "DbgNoAddressing", documented: false},
	0x7A: Opcode{mnemonic: "NO1", addressingFunc: "DbgNoAddressing", documented: false},
	0xDA: Opcode{mnemonic: "NO1", addressingFunc: "DbgNoAddressing", documented: false},
	0xFA: Opcode{mnemonic: "NO1", addressingFunc: "DbgNoAddressing", documented: false},
	0x80: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0x82: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0x89: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0xC2: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0xE2: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0x04: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0x14: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0x34: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0x44: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0x54: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0x64: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0x74: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0xD4: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0xF4: Opcode{mnemonic: "NO2", addressingFunc: "DbgNoAddressing", documented: false},
	0x0C: Opcode{mnemonic: "NO3", addressingFunc: "DbgNoAddressing", documented: false},
	0x1C: Opcode{mnemonic: "NO3", addressingFunc: "DbgNoAddressing", documented: false},
	0x3C: Opcode{mnemonic: "NO3", addressingFunc: "DbgNoAddressing", documented: false},
	0x5C: Opcode{mnemonic: "NO3", addressingFunc: "DbgNoAddressing", documented: false},
	0x7C: Opcode{mnemonic: "NO3", addressingFunc: "DbgNoAddressing", documented: false},
	0xDC: Opcode{mnemonic: "NO3", addressingFunc: "DbgNoAddressing", documented: false},
	0xFC: Opcode{mnemonic: "NO3", addressingFunc: "DbgNoAddressing", documented: false},
}