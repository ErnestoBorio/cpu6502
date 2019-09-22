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
			addressingFunc: "debugNoAddressing", 
			instructionFunc: "debugUndocumented",
			documented: false,
			cycles: 0,
		}
	}

	// Now overwrite documented operations with actual data
	for cod, op := range Opcodes {
		ops[cod] = op
	}

	for cod,_ := range ops {
		ops[cod].cycles = opcodeCycles[cod] // integrate cycles from the other array
	}

	fmt.Print(
`package cpu6502

type Operation struct {
	Opcode      uint8
	Mnemonic    string
	Cycles      uint8
	Documented  bool
	Addressing  func(*CPU) uint16
	Instruction func(*CPU, uint16)
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
		fmt.Printf("\t{ Opcode: 0x%02X, Mnemonic: \"%s\"%s, Cycles: %d, Documented: %v,%s Addressing: (*CPU).%s, Instruction: (*CPU).%s },\n", 
			cod, op.mnemonic, mnemPad, op.cycles, op.documented, docPad, op.addressingFunc, op.instructionFunc)
	}
	fmt.Println("}")
}



type Opcode struct {
	code byte
	mnemonic string
	codeStr string
	addrLabel string
	addressing string
	addressingFunc string
	instructionFunc string
	documented bool
	cycles uint8
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
	0xA9: Opcode{0xA9, "LDA", "A9", "IMM", "Immediate", "immediate", "_LDA", true, 0},
	0xA5: Opcode{0xA5, "LDA", "A5", "ZRP", "Zero page", "zeroPage", "_LDA", true, 0},
	0xB5: Opcode{0xB5, "LDA", "B5", "ZPX", "Zero page X", "debugZeroPageX", "_LDA", true, 0},
	0xAD: Opcode{0xAD, "LDA", "AD", "ABS", "Absolute", "absolute", "_LDA", true, 0},
	0xBD: Opcode{0xBD, "LDA", "BD", "ABX", "Absolute X", "debugAbsoluteX", "_LDA", true, 0},
	0xB9: Opcode{0xB9, "LDA", "B9", "ABY", "Absolute Y", "debugAbsoluteY", "_LDA", true, 0},
	0xA1: Opcode{0xA1, "LDA", "A1", "IIX", "Indexed Indirect X", "indexedIndirectX", "_LDA", true, 0},
	0xB1: Opcode{0xB1, "LDA", "B1", "IIY", "Indirect Indexed Y", "indirectIndexedY", "_LDA", true, 0},
	0xA2: Opcode{0xA2, "LDX", "A2", "IMM", "Immediate", "immediate", "_LDX", true, 0},
	0xA6: Opcode{0xA6, "LDX", "A6", "ZRP", "Zero page", "zeroPage", "_LDX", true, 0},
	0xB6: Opcode{0xB6, "LDX", "B6", "ZPY", "Zero page Y", "debugZeroPageY", "_LDX", true, 0},
	0xAE: Opcode{0xAE, "LDX", "AE", "ABS", "Absolute", "absolute", "_LDX", true, 0},
	0xBE: Opcode{0xBE, "LDX", "BE", "ABY", "Absolute Y", "debugAbsoluteY", "_LDX", true, 0},
	0xA0: Opcode{0xA0, "LDY", "A0", "IMM", "Immediate", "immediate", "_LDY", true, 0},
	0xA4: Opcode{0xA4, "LDY", "A4", "ZRP", "Zero page", "zeroPage", "_LDY", true, 0},
	0xB4: Opcode{0xB4, "LDY", "B4", "ZPX", "Zero page X", "debugZeroPageX", "_LDY", true, 0},
	0xAC: Opcode{0xAC, "LDY", "AC", "ABS", "Absolute", "absolute", "_LDY", true, 0},
	0xBC: Opcode{0xBC, "LDY", "BC", "ABX", "Absolute X", "debugAbsoluteX", "_LDY", true, 0},
	0x85: Opcode{0x85, "STA", "85", "ZRP", "Zero page", "zeroPage", "_STA", true, 0},
	0x95: Opcode{0x95, "STA", "95", "ZPX", "Zero page X", "debugZeroPageX", "_STA", true, 0},
	0x8D: Opcode{0x8D, "STA", "8D", "ABS", "Absolute", "absolute", "_STA", true, 0},
	0x9D: Opcode{0x9D, "STA", "9D", "ABX", "Absolute X", "debugAbsoluteX", "_STA", true, 0},
	0x99: Opcode{0x99, "STA", "99", "ABY", "Absolute Y", "debugAbsoluteY", "_STA", true, 0},
	0x81: Opcode{0x81, "STA", "81", "IIX", "Indexed Indirect X", "indexedIndirectX", "_STA", true, 0},
	0x91: Opcode{0x91, "STA", "91", "IIY", "Indirect Indexed Y", "indirectIndexedY", "_STA", true, 0},
	0x86: Opcode{0x86, "STX", "86", "ZRP", "Zero page", "zeroPage", "_STX", true, 0},
	0x96: Opcode{0x96, "STX", "96", "ZPY", "Zero page Y", "debugZeroPageY", "_STX", true, 0},
	0x8E: Opcode{0x8E, "STX", "8E", "ABS", "Absolute", "absolute", "_STX", true, 0},
	0x84: Opcode{0x84, "STY", "84", "ZRP", "Zero page", "zeroPage", "_STY", true, 0},
	0x94: Opcode{0x94, "STY", "94", "ZPX", "Zero page X", "debugZeroPageX", "_STY", true, 0},
	0x8C: Opcode{0x8C, "STY", "8C", "ABS", "Absolute", "absolute", "_STY", true, 0},
	0xAA: Opcode{0xAA, "TAX", "AA", "IMM", "Immediate", "immediate", "_TAX", true, 0},
	0xA8: Opcode{0xA8, "TAY", "A8", "IMM", "Immediate", "immediate", "_TAY", true, 0},
	0x8A: Opcode{0x8A, "TXA", "8A", "IMM", "Immediate", "immediate", "_TXA", true, 0},
	0x98: Opcode{0x98, "TYA", "98", "IMM", "Immediate", "immediate", "_TYA", true, 0},
	0xBA: Opcode{0xBA, "TSX", "BA", "IMM", "Immediate", "immediate", "_TSX", true, 0},
	0x9A: Opcode{0x9A, "TXS", "9A", "IMM", "Immediate", "immediate", "_TXS", true, 0},
	0x48: Opcode{0x48, "PHA", "48", "IMP", "Implied", "debugNoAddressing", "_PHA", true, 0},
	0x08: Opcode{0x08, "PHP", "08", "IMP", "Implied", "debugNoAddressing", "_PHP", true, 0},
	0x68: Opcode{0x68, "PLA", "68", "IMP", "Implied", "debugNoAddressing", "_PLA", true, 0},
	0x28: Opcode{0x28, "PLP", "28", "IMP", "Implied", "debugNoAddressing", "_PLP", true, 0},
	0x29: Opcode{0x29, "AND", "29", "IMM", "Immediate", "immediate", "and", true, 0},
	0x25: Opcode{0x25, "AND", "25", "ZRP", "Zero page", "zeroPage", "and", true, 0},
	0x35: Opcode{0x35, "AND", "35", "ZPX", "Zero page X", "debugZeroPageX", "and", true, 0},
	0x2D: Opcode{0x2D, "AND", "2D", "ABS", "Absolute", "absolute", "and", true, 0},
	0x3D: Opcode{0x3D, "AND", "3D", "ABX", "Absolute X", "debugAbsoluteX", "and", true, 0},
	0x39: Opcode{0x39, "AND", "39", "ABY", "Absolute Y", "debugAbsoluteY", "and", true, 0},
	0x21: Opcode{0x21, "AND", "21", "IIX", "Indexed Indirect X", "indexedIndirectX", "and", true, 0},
	0x31: Opcode{0x31, "AND", "31", "IIY", "Indirect Indexed Y", "indirectIndexedY", "and", true, 0},
	0x49: Opcode{0x49, "EOR", "49", "IMM", "Immediate", "immediate", "eor", true, 0},
	0x45: Opcode{0x45, "EOR", "45", "ZRP", "Zero page", "zeroPage", "eor", true, 0},
	0x55: Opcode{0x55, "EOR", "55", "ZPX", "Zero page X", "debugZeroPageX", "eor", true, 0},
	0x4D: Opcode{0x4D, "EOR", "4D", "ABS", "Absolute", "absolute", "eor", true, 0},
	0x5D: Opcode{0x5D, "EOR", "5D", "ABX", "Absolute X", "debugAbsoluteX", "eor", true, 0},
	0x59: Opcode{0x59, "EOR", "59", "ABY", "Absolute Y", "debugAbsoluteY", "eor", true, 0},
	0x41: Opcode{0x41, "EOR", "41", "IIX", "Indexed Indirect X", "indexedIndirectX", "eor", true, 0},
	0x51: Opcode{0x51, "EOR", "51", "IIY", "Indirect Indexed Y", "indirectIndexedY", "eor", true, 0},
	0x09: Opcode{0x09, "ORA", "09", "IMM", "Immediate", "immediate", "ora", true, 0},
	0x05: Opcode{0x05, "ORA", "05", "ZRP", "Zero page", "zeroPage", "ora", true, 0},
	0x15: Opcode{0x15, "ORA", "15", "ZPX", "Zero page X", "debugZeroPageX", "ora", true, 0},
	0x0D: Opcode{0x0D, "ORA", "0D", "ABS", "Absolute", "absolute", "ora", true, 0},
	0x1D: Opcode{0x1D, "ORA", "1D", "ABX", "Absolute X", "debugAbsoluteX", "ora", true, 0},
	0x19: Opcode{0x19, "ORA", "19", "ABY", "Absolute Y", "debugAbsoluteY", "ora", true, 0},
	0x01: Opcode{0x01, "ORA", "01", "IIX", "Indexed Indirect X", "indexedIndirectX", "ora", true, 0},
	0x11: Opcode{0x11, "ORA", "11", "IIY", "Indirect Indexed Y", "indirectIndexedY", "ora", true, 0},
	0x24: Opcode{0x24, "BIT", "24", "ZRP", "Zero page", "zeroPage", "bit", true, 0},
	0x2C: Opcode{0x2C, "BIT", "2C", "ABS", "Absolute", "absolute", "bit", true, 0},
	0x69: Opcode{0x69, "ADC", "69", "IMM", "Immediate", "immediate", "adc", true, 0},
	0x65: Opcode{0x65, "ADC", "65", "ZRP", "Zero page", "zeroPage", "adc", true, 0},
	0x75: Opcode{0x75, "ADC", "75", "ZPX", "Zero page X", "debugZeroPageX", "adc", true, 0},
	0x6D: Opcode{0x6D, "ADC", "6D", "ABS", "Absolute", "absolute", "adc", true, 0},
	0x7D: Opcode{0x7D, "ADC", "7D", "ABX", "Absolute X", "debugAbsoluteX", "adc", true, 0},
	0x79: Opcode{0x79, "ADC", "79", "ABY", "Absolute Y", "debugAbsoluteY", "adc", true, 0},
	0x61: Opcode{0x61, "ADC", "61", "IIX", "Indexed Indirect X", "indexedIndirectX", "adc", true, 0},
	0x71: Opcode{0x71, "ADC", "71", "IIY", "Indirect Indexed Y", "indirectIndexedY", "adc", true, 0},
	0xE9: Opcode{0xE9, "SBC", "E9", "IMM", "Immediate", "immediate", "sbc", true, 0},
	0xE5: Opcode{0xE5, "SBC", "E5", "ZRP", "Zero page", "zeroPage", "sbc", true, 0},
	0xF5: Opcode{0xF5, "SBC", "F5", "ZPX", "Zero page X", "debugZeroPageX", "sbc", true, 0},
	0xED: Opcode{0xED, "SBC", "ED", "ABS", "Absolute", "absolute", "sbc", true, 0},
	0xFD: Opcode{0xFD, "SBC", "FD", "ABX", "Absolute X", "debugAbsoluteX", "sbc", true, 0},
	0xF9: Opcode{0xF9, "SBC", "F9", "ABY", "Absolute Y", "debugAbsoluteY", "sbc", true, 0},
	0xE1: Opcode{0xE1, "SBC", "E1", "IIX", "Indexed Indirect X", "indexedIndirectX", "sbc", true, 0},
	0xF1: Opcode{0xF1, "SBC", "F1", "IIY", "Indirect Indexed Y", "indirectIndexedY", "sbc", true, 0},
	0xC9: Opcode{0xC9, "CMP", "C9", "IMM", "Immediate", "immediate", "_CMP", true, 0},
	0xC5: Opcode{0xC5, "CMP", "C5", "ZRP", "Zero page", "zeroPage", "_CMP", true, 0},
	0xD5: Opcode{0xD5, "CMP", "D5", "ZPX", "Zero page X", "debugZeroPageX", "_CMP", true, 0},
	0xCD: Opcode{0xCD, "CMP", "CD", "ABS", "Absolute", "absolute", "_CMP", true, 0},
	0xDD: Opcode{0xDD, "CMP", "DD", "ABX", "Absolute X", "debugAbsoluteX", "_CMP", true, 0},
	0xD9: Opcode{0xD9, "CMP", "D9", "ABY", "Absolute Y", "debugAbsoluteY", "_CMP", true, 0},
	0xC1: Opcode{0xC1, "CMP", "C1", "IIX", "Indexed Indirect X", "indexedIndirectX", "_CMP", true, 0},
	0xD1: Opcode{0xD1, "CMP", "D1", "IIY", "Indirect Indexed Y", "indirectIndexedY", "_CMP", true, 0},
	0xE0: Opcode{0xE0, "CPX", "E0", "IMM", "Immediate", "immediate", "_CPX", true, 0},
	0xE4: Opcode{0xE4, "CPX", "E4", "ZRP", "Zero page", "zeroPage", "_CPX", true, 0},
	0xEC: Opcode{0xEC, "CPX", "EC", "ABS", "Absolute", "absolute", "_CPX", true, 0},
	0xC0: Opcode{0xC0, "CPY", "C0", "IMM", "Immediate", "immediate", "_CPY", true, 0},
	0xC4: Opcode{0xC4, "CPY", "C4", "ZRP", "Zero page", "zeroPage", "_CPY", true, 0},
	0xCC: Opcode{0xCC, "CPY", "CC", "ABS", "Absolute", "absolute", "_CPY", true, 0},
	0xE6: Opcode{0xE6, "INC", "E6", "ZRP", "Zero page", "zeroPage", "_INC", true, 0},
	0xF6: Opcode{0xF6, "INC", "F6", "ZPX", "Zero page X", "debugZeroPageX", "_INC", true, 0},
	0xEE: Opcode{0xEE, "INC", "EE", "ABS", "Absolute", "absolute", "_INC", true, 0},
	0xFE: Opcode{0xFE, "INC", "FE", "ABX", "Absolute X", "debugAbsoluteX", "_INC", true, 0},
	0xE8: Opcode{0xE8, "INX", "E8", "IMP", "Implied", "debugNoAddressing", "_INX", true, 0},
	0xC8: Opcode{0xC8, "INY", "C8", "IMP", "Implied", "debugNoAddressing", "_INY", true, 0},
	0xC6: Opcode{0xC6, "DEC", "C6", "ZRP", "Zero page", "zeroPage", "_DEC", true, 0},
	0xD6: Opcode{0xD6, "DEC", "D6", "ZPX", "Zero page X", "debugZeroPageX", "_DEC", true, 0},
	0xCE: Opcode{0xCE, "DEC", "CE", "ABS", "Absolute", "absolute", "_DEC", true, 0},
	0xDE: Opcode{0xDE, "DEC", "DE", "ABX", "Absolute X", "debugAbsoluteX", "_DEC", true, 0},
	0xCA: Opcode{0xCA, "DEX", "CA", "IMP", "Implied", "debugNoAddressing", "_DEX", true, 0},
	0x88: Opcode{0x88, "DEY", "88", "IMP", "Implied", "debugNoAddressing", "_DEY", true, 0},
	0x0A: Opcode{0x0A, "ASL", "0A", "ACU", "Accumulator", "debugNoAddressing", "_ASLa", true, 0},
	0x4A: Opcode{0x4A, "LSR", "4A", "ACU", "Accumulator", "debugNoAddressing", "_LSRa", true, 0},
	0x46: Opcode{0x46, "LSR", "46", "ZRP", "Zero page", "zeroPage", "lsr", true, 0},
	0x56: Opcode{0x56, "LSR", "56", "ZPX", "Zero page X", "debugZeroPageX", "lsr", true, 0},
	0x4E: Opcode{0x4E, "LSR", "4E", "ABS", "Absolute", "absolute", "lsr", true, 0},
	0x5E: Opcode{0x5E, "LSR", "5E", "ABX", "Absolute X", "debugAbsoluteX", "lsr", true, 0},
	0x2A: Opcode{0x2A, "ROL", "2A", "ACU", "Accumulator", "debugNoAddressing", "_ROLa", true, 0},
	0x26: Opcode{0x26, "ROL", "26", "ZRP", "Zero page", "zeroPage", "rol", true, 0},
	0x36: Opcode{0x36, "ROL", "36", "ZPX", "Zero page X", "debugZeroPageX", "rol", true, 0},
	0x2E: Opcode{0x2E, "ROL", "2E", "ABS", "Absolute", "absolute", "rol", true, 0},
	0x3E: Opcode{0x3E, "ROL", "3E", "ABX", "Absolute X", "debugAbsoluteX", "rol", true, 0},
	0x6A: Opcode{0x6A, "ROR", "6A", "ACU", "Accumulator", "debugNoAddressing", "_RORa", true, 0},
	0x66: Opcode{0x66, "ROR", "66", "ZRP", "Zero page", "zeroPage", "ror", true, 0},
	0x76: Opcode{0x76, "ROR", "76", "ZPX", "Zero page X", "debugZeroPageX", "ror", true, 0},
	0x6E: Opcode{0x6E, "ROR", "6E", "ABS", "Absolute", "absolute", "ror", true, 0},
	0x7E: Opcode{0x7E, "ROR", "7E", "ABX", "Absolute X", "debugAbsoluteX", "ror", true, 0},
	0x4C: Opcode{0x4C, "JMP", "4C", "ABS", "Absolute", "absolute", "_JMP", true, 0},
	0x6C: Opcode{0x6C, "JMP", "6C", "IND", "Indirect", "debugIndirect", "_JMP", true, 0},
	0x20: Opcode{0x20, "JSR", "20", "IMP", "Implied", "debugNoAddressing", "_JSR", true, 0},
	0x60: Opcode{0x60, "RTS", "60", "IMP", "Implied", "debugNoAddressing", "_RTS", true, 0},
	0x90: Opcode{0x90, "BCC", "90", "REL", "Relative", "debugNoAddressing", "_BCC", true, 0},
	0xB0: Opcode{0xB0, "BCS", "B0", "REL", "Relative", "debugNoAddressing", "_BCS", true, 0},
	0xF0: Opcode{0xF0, "BEQ", "F0", "REL", "Relative", "debugNoAddressing", "_BEQ", true, 0},
	0x30: Opcode{0x30, "BMI", "30", "REL", "Relative", "debugNoAddressing", "_BMI", true, 0},
	0xD0: Opcode{0xD0, "BNE", "D0", "REL", "Relative", "debugNoAddressing", "_BNE", true, 0},
	0x10: Opcode{0x10, "BPL", "10", "REL", "Relative", "debugNoAddressing", "_BPL", true, 0},
	0x50: Opcode{0x50, "BVC", "50", "REL", "Relative", "debugNoAddressing", "_BVC", true, 0},
	0x70: Opcode{0x70, "BVS", "70", "REL", "Relative", "debugNoAddressing", "_BVS", true, 0},
	0x18: Opcode{0x18, "CLC", "18", "IMP", "Implied", "debugNoAddressing", "_CLC", true, 0},
	0xD8: Opcode{0xD8, "CLD", "D8", "IMP", "Implied", "debugNoAddressing", "_CLD", true, 0},
	0x58: Opcode{0x58, "CLI", "58", "IMP", "Implied", "debugNoAddressing", "_CLI", true, 0},
	0xB8: Opcode{0xB8, "CLV", "B8", "IMP", "Implied", "debugNoAddressing", "_CLV", true, 0},
	0x38: Opcode{0x38, "SEC", "38", "IMP", "Implied", "debugNoAddressing", "_SEC", true, 0},
	0xF8: Opcode{0xF8, "SED", "F8", "IMP", "Implied", "debugNoAddressing", "_SED", true, 0},
	0x78: Opcode{0x78, "SEI", "78", "IMP", "Implied", "debugNoAddressing", "_SEI", true, 0},
	0x00: Opcode{0x00, "BRK", "00", "IMP", "Implied", "debugNoAddressing", "_BRK", true, 0},
	0xEA: Opcode{0xEA, "NOP", "EA", "IMP", "Implied", "debugNoAddressing", "_NOP", true, 0},
	0x40: Opcode{0x40, "RTI", "40", "IMP", "Implied", "debugNoAddressing", "_RTI", true, 0},
	0x1A: Opcode{mnemonic: "NO1", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x3A: Opcode{mnemonic: "NO1", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x5A: Opcode{mnemonic: "NO1", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x7A: Opcode{mnemonic: "NO1", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0xDA: Opcode{mnemonic: "NO1", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0xFA: Opcode{mnemonic: "NO1", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x80: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x82: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x89: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0xC2: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0xE2: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x04: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x14: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x34: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x44: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x54: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x64: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x74: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0xD4: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0xF4: Opcode{mnemonic: "NO2", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x0C: Opcode{mnemonic: "NO3", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x1C: Opcode{mnemonic: "NO3", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x3C: Opcode{mnemonic: "NO3", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x5C: Opcode{mnemonic: "NO3", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0x7C: Opcode{mnemonic: "NO3", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0xDC: Opcode{mnemonic: "NO3", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
	0xFC: Opcode{mnemonic: "NO3", addressingFunc: "debugNoAddressing", instructionFunc: "debugUndocNOP", documented: false},
}