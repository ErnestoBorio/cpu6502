export default [
	{
		"opcodeNum": 0,
		"opcode": "00",
		"mnemonic": "BRK",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 7
	},
	{
		"opcodeNum": 1,
		"opcode": "01",
		"mnemonic": "ORA",
		"addressing": "IDX",
		"addressingLong": "(Indirect,X)",
		"bytes": 2,
		"cycles": 6
	},
	null,
	null,
	null,
	{
		"opcodeNum": 5,
		"opcode": "05",
		"mnemonic": "ORA",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 6,
		"opcode": "06",
		"mnemonic": "ASL",
		"addressing": "ZPG",
		"addressingLong": "Zero\n      Page",
		"bytes": 2,
		"cycles": 5
	},
	null,
	{
		"opcodeNum": 8,
		"opcode": "08",
		"mnemonic": "PHP",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 3
	},
	{
		"opcodeNum": 9,
		"opcode": "09",
		"mnemonic": "ORA",
		"addressing": "IMM",
		"addressingLong": "Immediate",
		"bytes": 2,
		"cycles": 2
	},
	{
		"opcodeNum": 10,
		"opcode": "0A",
		"mnemonic": "ASL",
		"addressing": "IMP",
		"addressingLong": "Accumulator",
		"bytes": 1,
		"cycles": 2
	},
	null,
	null,
	{
		"opcodeNum": 13,
		"opcode": "0D",
		"mnemonic": "ORA",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 14,
		"opcode": "0E",
		"mnemonic": "ASL",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 16,
		"opcode": "10",
		"mnemonic": "BPL",
		"addressing": "REL",
		"addressingLong": "Relative",
		"bytes": 2,
		"cycles": 2,
		"extra": "(+1 if branch succeeds\n      +2 if to a new page)"
	},
	{
		"opcodeNum": 17,
		"opcode": "11",
		"mnemonic": "ORA",
		"addressing": "IDY",
		"addressingLong": "(Indirect),Y",
		"bytes": 2,
		"cycles": 5,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 21,
		"opcode": "15",
		"mnemonic": "ORA",
		"addressing": "ZPX",
		"addressingLong": "Zero Page,X",
		"bytes": 2,
		"cycles": 4
	},
	{
		"opcodeNum": 22,
		"opcode": "16",
		"mnemonic": "ASL",
		"addressing": "ZPX",
		"addressingLong": "Zero\n      Page,X",
		"bytes": 2,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 24,
		"opcode": "18",
		"mnemonic": "CLC",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	{
		"opcodeNum": 25,
		"opcode": "19",
		"mnemonic": "ORA",
		"addressing": "ABY",
		"addressingLong": "Absolute,Y",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 29,
		"opcode": "1D",
		"mnemonic": "ORA",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	{
		"opcodeNum": 30,
		"opcode": "1E",
		"mnemonic": "ASL",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 7
	},
	null,
	{
		"opcodeNum": 32,
		"opcode": "20",
		"mnemonic": "JSR",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 6
	},
	{
		"opcodeNum": 33,
		"opcode": "21",
		"mnemonic": "AND",
		"addressing": "IDX",
		"addressingLong": "(Indirect,X)",
		"bytes": 2,
		"cycles": 6
	},
	null,
	null,
	{
		"opcodeNum": 36,
		"opcode": "24",
		"mnemonic": "BIT",
		"addressing": "ZPG",
		"addressingLong": "Zero\n      Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 37,
		"opcode": "25",
		"mnemonic": "AND",
		"addressing": "ZPG",
		"addressingLong": "Zero\n      Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 38,
		"opcode": "26",
		"mnemonic": "ROL",
		"addressing": "ZPG",
		"addressingLong": "Zero\n      Page",
		"bytes": 2,
		"cycles": 5
	},
	null,
	{
		"opcodeNum": 40,
		"opcode": "28",
		"mnemonic": "PLP",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 4
	},
	{
		"opcodeNum": 41,
		"opcode": "29",
		"mnemonic": "AND",
		"addressing": "IMM",
		"addressingLong": "Immediate",
		"bytes": 2,
		"cycles": 2
	},
	{
		"opcodeNum": 42,
		"opcode": "2A",
		"mnemonic": "ROL",
		"addressing": "IMP",
		"addressingLong": "Accumulator",
		"bytes": 1,
		"cycles": 2
	},
	null,
	{
		"opcodeNum": 44,
		"opcode": "2C",
		"mnemonic": "BIT",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 45,
		"opcode": "2D",
		"mnemonic": "AND",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 46,
		"opcode": "2E",
		"mnemonic": "ROL",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 48,
		"opcode": "30",
		"mnemonic": "BMI",
		"addressing": "REL",
		"addressingLong": "Relative",
		"bytes": 2,
		"cycles": 2,
		"extra": "(+1 if branch succeeds\n      +2 if to a new page)"
	},
	{
		"opcodeNum": 49,
		"opcode": "31",
		"mnemonic": "AND",
		"addressing": "IDY",
		"addressingLong": "(Indirect),Y",
		"bytes": 2,
		"cycles": 5,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 53,
		"opcode": "35",
		"mnemonic": "AND",
		"addressing": "ZPX",
		"addressingLong": "Zero\n      Page,X",
		"bytes": 2,
		"cycles": 4
	},
	{
		"opcodeNum": 54,
		"opcode": "36",
		"mnemonic": "ROL",
		"addressing": "ZPX",
		"addressingLong": "Zero\n      Page,X",
		"bytes": 2,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 56,
		"opcode": "38",
		"mnemonic": "SEC",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	{
		"opcodeNum": 57,
		"opcode": "39",
		"mnemonic": "AND",
		"addressing": "ABY",
		"addressingLong": "Absolute,Y",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 61,
		"opcode": "3D",
		"mnemonic": "AND",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	{
		"opcodeNum": 62,
		"opcode": "3E",
		"mnemonic": "ROL",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 7
	},
	null,
	{
		"opcodeNum": 64,
		"opcode": "40",
		"mnemonic": "RTI",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 6
	},
	{
		"opcodeNum": 65,
		"opcode": "41",
		"mnemonic": "EOR",
		"addressing": "IDX",
		"addressingLong": "(Indirect,X)",
		"bytes": 2,
		"cycles": 6
	},
	null,
	null,
	null,
	{
		"opcodeNum": 69,
		"opcode": "45",
		"mnemonic": "EOR",
		"addressing": "ZPG",
		"addressingLong": "Zero\n      Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 70,
		"opcode": "46",
		"mnemonic": "LSR",
		"addressing": "ZPG",
		"addressingLong": "Zero\n      Page",
		"bytes": 2,
		"cycles": 5
	},
	null,
	{
		"opcodeNum": 72,
		"opcode": "48",
		"mnemonic": "PHA",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 3
	},
	{
		"opcodeNum": 73,
		"opcode": "49",
		"mnemonic": "EOR",
		"addressing": "IMM",
		"addressingLong": "Immediate",
		"bytes": 2,
		"cycles": 2
	},
	{
		"opcodeNum": 74,
		"opcode": "4A",
		"mnemonic": "LSR",
		"addressing": "IMP",
		"addressingLong": "Accumulator",
		"bytes": 1,
		"cycles": 2
	},
	null,
	{
		"opcodeNum": 76,
		"opcode": "4C",
		"mnemonic": "JMP",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 3
	},
	{
		"opcodeNum": 77,
		"opcode": "4D",
		"mnemonic": "EOR",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 78,
		"opcode": "4E",
		"mnemonic": "LSR",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 80,
		"opcode": "50",
		"mnemonic": "BVC",
		"addressing": "REL",
		"addressingLong": "Relative",
		"bytes": 2,
		"cycles": 2,
		"extra": "(+1 if branch succeeds\n      +2 if to a new page)"
	},
	{
		"opcodeNum": 81,
		"opcode": "51",
		"mnemonic": "EOR",
		"addressing": "IDY",
		"addressingLong": "(Indirect),Y",
		"bytes": 2,
		"cycles": 5,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 85,
		"opcode": "55",
		"mnemonic": "EOR",
		"addressing": "ZPX",
		"addressingLong": "Zero\n      Page,X",
		"bytes": 2,
		"cycles": 4
	},
	{
		"opcodeNum": 86,
		"opcode": "56",
		"mnemonic": "LSR",
		"addressing": "ZPX",
		"addressingLong": "Zero\n      Page,X",
		"bytes": 2,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 88,
		"opcode": "58",
		"mnemonic": "CLI",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	{
		"opcodeNum": 89,
		"opcode": "59",
		"mnemonic": "EOR",
		"addressing": "ABY",
		"addressingLong": "Absolute,Y",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 93,
		"opcode": "5D",
		"mnemonic": "EOR",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	{
		"opcodeNum": 94,
		"opcode": "5E",
		"mnemonic": "LSR",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 7
	},
	null,
	{
		"opcodeNum": 96,
		"opcode": "60",
		"mnemonic": "RTS",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 6
	},
	{
		"opcodeNum": 97,
		"opcode": "61",
		"mnemonic": "ADC",
		"addressing": "IDX",
		"addressingLong": "(Indirect,X)",
		"bytes": 2,
		"cycles": 6
	},
	null,
	null,
	null,
	{
		"opcodeNum": 101,
		"opcode": "65",
		"mnemonic": "ADC",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 102,
		"opcode": "66",
		"mnemonic": "ROR",
		"addressing": "ZPG",
		"addressingLong": "Zero\n      Page",
		"bytes": 2,
		"cycles": 5
	},
	null,
	{
		"opcodeNum": 104,
		"opcode": "68",
		"mnemonic": "PLA",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 4
	},
	{
		"opcodeNum": 105,
		"opcode": "69",
		"mnemonic": "ADC",
		"addressing": "IMM",
		"addressingLong": "Immediate",
		"bytes": 2,
		"cycles": 2
	},
	{
		"opcodeNum": 106,
		"opcode": "6A",
		"mnemonic": "ROR",
		"addressing": "IMP",
		"addressingLong": "Accumulator",
		"bytes": 1,
		"cycles": 2
	},
	null,
	{
		"opcodeNum": 108,
		"opcode": "6C",
		"mnemonic": "JMP",
		"addressing": "IND",
		"addressingLong": "Indirect",
		"bytes": 3,
		"cycles": 5
	},
	{
		"opcodeNum": 109,
		"opcode": "6D",
		"mnemonic": "ADC",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 110,
		"opcode": "6E",
		"mnemonic": "ROR",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 112,
		"opcode": "70",
		"mnemonic": "BVS",
		"addressing": "REL",
		"addressingLong": "Relative",
		"bytes": 2,
		"cycles": 2,
		"extra": "(+1 if branch succeeds\n      +2 if to a new page)"
	},
	{
		"opcodeNum": 113,
		"opcode": "71",
		"mnemonic": "ADC",
		"addressing": "IDY",
		"addressingLong": "(Indirect),Y",
		"bytes": 2,
		"cycles": 5,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 117,
		"opcode": "75",
		"mnemonic": "ADC",
		"addressing": "ZPX",
		"addressingLong": "Zero Page,X",
		"bytes": 2,
		"cycles": 4
	},
	{
		"opcodeNum": 118,
		"opcode": "76",
		"mnemonic": "ROR",
		"addressing": "ZPX",
		"addressingLong": "Zero\n      Page,X",
		"bytes": 2,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 120,
		"opcode": "78",
		"mnemonic": "SEI",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	{
		"opcodeNum": 121,
		"opcode": "79",
		"mnemonic": "ADC",
		"addressing": "ABY",
		"addressingLong": "Absolute,Y",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 125,
		"opcode": "7D",
		"mnemonic": "ADC",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	{
		"opcodeNum": 126,
		"opcode": "7E",
		"mnemonic": "ROR",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 7
	},
	null,
	null,
	{
		"opcodeNum": 129,
		"opcode": "81",
		"mnemonic": "STA",
		"addressing": "IDX",
		"addressingLong": "(Indirect,X)",
		"bytes": 2,
		"cycles": 6
	},
	null,
	null,
	{
		"opcodeNum": 132,
		"opcode": "84",
		"mnemonic": "STY",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 133,
		"opcode": "85",
		"mnemonic": "STA",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 134,
		"opcode": "86",
		"mnemonic": "STX",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	null,
	{
		"opcodeNum": 136,
		"opcode": "88",
		"mnemonic": "DEY",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	null,
	{
		"opcodeNum": 138,
		"opcode": "8A",
		"mnemonic": "TXA",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	null,
	{
		"opcodeNum": 140,
		"opcode": "8C",
		"mnemonic": "STY",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 141,
		"opcode": "8D",
		"mnemonic": "STA",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 142,
		"opcode": "8E",
		"mnemonic": "STX",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	null,
	{
		"opcodeNum": 144,
		"opcode": "90",
		"mnemonic": "BCC",
		"addressing": "REL",
		"addressingLong": "Relative",
		"bytes": 2,
		"cycles": 2,
		"extra": "(+1 if branch succeeds\n      +2 if to a new page)"
	},
	{
		"opcodeNum": 145,
		"opcode": "91",
		"mnemonic": "STA",
		"addressing": "IDY",
		"addressingLong": "(Indirect),Y",
		"bytes": 2,
		"cycles": 6
	},
	null,
	null,
	{
		"opcodeNum": 148,
		"opcode": "94",
		"mnemonic": "STY",
		"addressing": "ZPX",
		"addressingLong": "Zero Page,X",
		"bytes": 2,
		"cycles": 4
	},
	{
		"opcodeNum": 149,
		"opcode": "95",
		"mnemonic": "STA",
		"addressing": "ZPX",
		"addressingLong": "Zero Page,X",
		"bytes": 2,
		"cycles": 4
	},
	{
		"opcodeNum": 150,
		"opcode": "96",
		"mnemonic": "STX",
		"addressing": "ZPY",
		"addressingLong": "Zero Page,Y",
		"bytes": 2,
		"cycles": 4
	},
	null,
	{
		"opcodeNum": 152,
		"opcode": "98",
		"mnemonic": "TYA",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	{
		"opcodeNum": 153,
		"opcode": "99",
		"mnemonic": "STA",
		"addressing": "ABY",
		"addressingLong": "Absolute,Y",
		"bytes": 3,
		"cycles": 5
	},
	{
		"opcodeNum": 154,
		"opcode": "9A",
		"mnemonic": "TXS",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	null,
	null,
	{
		"opcodeNum": 157,
		"opcode": "9D",
		"mnemonic": "STA",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 5
	},
	null,
	null,
	{
		"opcodeNum": 160,
		"opcode": "A0",
		"mnemonic": "LDY",
		"addressing": "IMM",
		"addressingLong": "Immediate",
		"bytes": 2,
		"cycles": 2
	},
	{
		"opcodeNum": 161,
		"opcode": "A1",
		"mnemonic": "LDA",
		"addressing": "IDX",
		"addressingLong": "(Indirect,X)",
		"bytes": 2,
		"cycles": 6
	},
	{
		"opcodeNum": 162,
		"opcode": "A2",
		"mnemonic": "LDX",
		"addressing": "IMM",
		"addressingLong": "Immediate",
		"bytes": 2,
		"cycles": 2
	},
	null,
	{
		"opcodeNum": 164,
		"opcode": "A4",
		"mnemonic": "LDY",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 165,
		"opcode": "A5",
		"mnemonic": "LDA",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 166,
		"opcode": "A6",
		"mnemonic": "LDX",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	null,
	{
		"opcodeNum": 168,
		"opcode": "A8",
		"mnemonic": "TAY",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	{
		"opcodeNum": 169,
		"opcode": "A9",
		"mnemonic": "LDA",
		"addressing": "IMM",
		"addressingLong": "Immediate",
		"bytes": 2,
		"cycles": 2
	},
	{
		"opcodeNum": 170,
		"opcode": "AA",
		"mnemonic": "TAX",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	null,
	{
		"opcodeNum": 172,
		"opcode": "AC",
		"mnemonic": "LDY",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 173,
		"opcode": "AD",
		"mnemonic": "LDA",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 174,
		"opcode": "AE",
		"mnemonic": "LDX",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	null,
	{
		"opcodeNum": 176,
		"opcode": "B0",
		"mnemonic": "BCS",
		"addressing": "REL",
		"addressingLong": "Relative",
		"bytes": 2,
		"cycles": 2,
		"extra": "(+1 if branch succeeds\n      +2 if to a new page)"
	},
	{
		"opcodeNum": 177,
		"opcode": "B1",
		"mnemonic": "LDA",
		"addressing": "IDY",
		"addressingLong": "(Indirect),Y",
		"bytes": 2,
		"cycles": 5,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	{
		"opcodeNum": 180,
		"opcode": "B4",
		"mnemonic": "LDY",
		"addressing": "ZPX",
		"addressingLong": "Zero Page,X",
		"bytes": 2,
		"cycles": 4
	},
	{
		"opcodeNum": 181,
		"opcode": "B5",
		"mnemonic": "LDA",
		"addressing": "ZPX",
		"addressingLong": "Zero Page,X",
		"bytes": 2,
		"cycles": 4
	},
	{
		"opcodeNum": 182,
		"opcode": "B6",
		"mnemonic": "LDX",
		"addressing": "ZPY",
		"addressingLong": "Zero Page,Y",
		"bytes": 2,
		"cycles": 4
	},
	null,
	{
		"opcodeNum": 184,
		"opcode": "B8",
		"mnemonic": "CLV",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	{
		"opcodeNum": 185,
		"opcode": "B9",
		"mnemonic": "LDA",
		"addressing": "ABY",
		"addressingLong": "Absolute,Y",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	{
		"opcodeNum": 186,
		"opcode": "BA",
		"mnemonic": "TSX",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	null,
	{
		"opcodeNum": 188,
		"opcode": "BC",
		"mnemonic": "LDY",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	{
		"opcodeNum": 189,
		"opcode": "BD",
		"mnemonic": "LDA",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	{
		"opcodeNum": 190,
		"opcode": "BE",
		"mnemonic": "LDX",
		"addressing": "ABY",
		"addressingLong": "Absolute,Y",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	null,
	{
		"opcodeNum": 192,
		"opcode": "C0",
		"mnemonic": "CPY",
		"addressing": "IMM",
		"addressingLong": "Immediate",
		"bytes": 2,
		"cycles": 2
	},
	{
		"opcodeNum": 193,
		"opcode": "C1",
		"mnemonic": "CMP",
		"addressing": "IDX",
		"addressingLong": "(Indirect,X)",
		"bytes": 2,
		"cycles": 6
	},
	null,
	null,
	{
		"opcodeNum": 196,
		"opcode": "C4",
		"mnemonic": "CPY",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 197,
		"opcode": "C5",
		"mnemonic": "CMP",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 198,
		"opcode": "C6",
		"mnemonic": "DEC",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 5
	},
	null,
	{
		"opcodeNum": 200,
		"opcode": "C8",
		"mnemonic": "INY",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	{
		"opcodeNum": 201,
		"opcode": "C9",
		"mnemonic": "CMP",
		"addressing": "IMM",
		"addressingLong": "Immediate",
		"bytes": 2,
		"cycles": 2
	},
	{
		"opcodeNum": 202,
		"opcode": "CA",
		"mnemonic": "DEX",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	null,
	{
		"opcodeNum": 204,
		"opcode": "CC",
		"mnemonic": "CPY",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 205,
		"opcode": "CD",
		"mnemonic": "CMP",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 206,
		"opcode": "CE",
		"mnemonic": "DEC",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 208,
		"opcode": "D0",
		"mnemonic": "BNE",
		"addressing": "REL",
		"addressingLong": "Relative",
		"bytes": 2,
		"cycles": 2,
		"extra": "(+1 if branch succeeds\n      +2 if to a new page)"
	},
	{
		"opcodeNum": 209,
		"opcode": "D1",
		"mnemonic": "CMP",
		"addressing": "IDY",
		"addressingLong": "(Indirect),Y",
		"bytes": 2,
		"cycles": 5,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 213,
		"opcode": "D5",
		"mnemonic": "CMP",
		"addressing": "ZPX",
		"addressingLong": "Zero Page,X",
		"bytes": 2,
		"cycles": 4
	},
	{
		"opcodeNum": 214,
		"opcode": "D6",
		"mnemonic": "DEC",
		"addressing": "ZPX",
		"addressingLong": "Zero Page,X",
		"bytes": 2,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 216,
		"opcode": "D8",
		"mnemonic": "CLD",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	{
		"opcodeNum": 217,
		"opcode": "D9",
		"mnemonic": "CMP",
		"addressing": "ABY",
		"addressingLong": "Absolute,Y",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 221,
		"opcode": "DD",
		"mnemonic": "CMP",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	{
		"opcodeNum": 222,
		"opcode": "DE",
		"mnemonic": "DEC",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 7
	},
	null,
	{
		"opcodeNum": 224,
		"opcode": "E0",
		"mnemonic": "CPX",
		"addressing": "IMM",
		"addressingLong": "Immediate",
		"bytes": 2,
		"cycles": 2
	},
	{
		"opcodeNum": 225,
		"opcode": "E1",
		"mnemonic": "SBC",
		"addressing": "IDX",
		"addressingLong": "(Indirect,X)",
		"bytes": 2,
		"cycles": 6
	},
	null,
	null,
	{
		"opcodeNum": 228,
		"opcode": "E4",
		"mnemonic": "CPX",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 229,
		"opcode": "E5",
		"mnemonic": "SBC",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 3
	},
	{
		"opcodeNum": 230,
		"opcode": "E6",
		"mnemonic": "INC",
		"addressing": "ZPG",
		"addressingLong": "Zero Page",
		"bytes": 2,
		"cycles": 5
	},
	null,
	{
		"opcodeNum": 232,
		"opcode": "E8",
		"mnemonic": "INX",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	{
		"opcodeNum": 233,
		"opcode": "E9",
		"mnemonic": "SBC",
		"addressing": "IMM",
		"addressingLong": "Immediate",
		"bytes": 2,
		"cycles": 2
	},
	{
		"opcodeNum": 234,
		"opcode": "EA",
		"mnemonic": "NOP",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	null,
	{
		"opcodeNum": 236,
		"opcode": "EC",
		"mnemonic": "CPX",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 237,
		"opcode": "ED",
		"mnemonic": "SBC",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 4
	},
	{
		"opcodeNum": 238,
		"opcode": "EE",
		"mnemonic": "INC",
		"addressing": "ABS",
		"addressingLong": "Absolute",
		"bytes": 3,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 240,
		"opcode": "F0",
		"mnemonic": "BEQ",
		"addressing": "REL",
		"addressingLong": "Relative",
		"bytes": 2,
		"cycles": 2,
		"extra": "(+1 if branch succeeds\n      +2 if to a new page)"
	},
	{
		"opcodeNum": 241,
		"opcode": "F1",
		"mnemonic": "SBC",
		"addressing": "IDY",
		"addressingLong": "(Indirect),Y",
		"bytes": 2,
		"cycles": 5,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 245,
		"opcode": "F5",
		"mnemonic": "SBC",
		"addressing": "ZPX",
		"addressingLong": "Zero Page,X",
		"bytes": 2,
		"cycles": 4
	},
	{
		"opcodeNum": 246,
		"opcode": "F6",
		"mnemonic": "INC",
		"addressing": "ZPX",
		"addressingLong": "Zero Page,X",
		"bytes": 2,
		"cycles": 6
	},
	null,
	{
		"opcodeNum": 248,
		"opcode": "F8",
		"mnemonic": "SED",
		"addressing": "IMP",
		"addressingLong": "Implied",
		"bytes": 1,
		"cycles": 2
	},
	{
		"opcodeNum": 249,
		"opcode": "F9",
		"mnemonic": "SBC",
		"addressing": "ABY",
		"addressingLong": "Absolute,Y",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	null,
	null,
	null,
	{
		"opcodeNum": 253,
		"opcode": "FD",
		"mnemonic": "SBC",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 4,
		"extra": "(+1 if page crossed)"
	},
	{
		"opcodeNum": 254,
		"opcode": "FE",
		"mnemonic": "INC",
		"addressing": "ABX",
		"addressingLong": "Absolute,X",
		"bytes": 3,
		"cycles": 7
	},
	null
];
