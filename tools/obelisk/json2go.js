#!/usr/bin/env node
/** json2go.js translates opcodes.js to opcodes.go with proper declarations. */

import opcodes from "./opcodes.js";

const outfile = process.stdout;

const addressing = {
	null: "implied",
	"IMP": "implied",
	"IMM": "immediate",
	"REL": "immediate",
	"ZPG": "zeroPage",
	"ZPX": "zeroPageX",
	"ZPY": "zeroPageY",
	"ABS": "absolute",
	"ABX": "absoluteX",
	"ABY": "absoluteY",
	"IND": "indirect",
	"IDX": "indexedIndirectX",
	"IDY": "indirectIndexedY"
}

outfile.write(`package cpu6502

// Operation models a CPU instruction
type Operation struct {
	Opcode      byte
	Length      uint8 // in bytes
	Cycles      int
	Documented  bool
	Addressing  func(*CPU) uint16
	Instruction func(*CPU, uint16)
}

// Only BRK (00) operation has no cycles here, because they're accounted for in the IRQ interrupt handler
var Opcodes = [0x100] Operation {
`);

for (let i = 0; i < opcodes.length; i++) {
	// Build an empty op object for undocumented opcodes.
	let documented = true;
	if (!opcodes[i]) {
		documented = false;
		opcodes[i] = {
			opcodeNum: i,
			opcode: i.toString(16).padStart(2, "0"),
			mnemonic: null,
			addressing: null,
			addressingLong: null,
			bytes: 1,
			cycles: 0
		}
	}
	let op = opcodes[i];
	if (op.opcodeNum == 0 ) op.cycles = 0; // For BRK ($00) the 7 cycles are counted for in irq()
	let func = !documented ? "undoc" : op.mnemonic.toLowerCase().slice(0, 3);
	// Bit shifting instructions on the accumulator need a separate function that doesn't use addressing
	if(['asl','lsr','ror','rol'].includes(func) && op.addressing == 'IMP') func += 'a';
	let padding = " ".repeat(6 - func.length);

	outfile.write(`\t{ Opcode: 0x${op.opcode}, Length: ${op.bytes}, Cycles: ${op.cycles}, Documented: ${documented}`);
	if(documented) outfile.write(" ");
	outfile.write(`, Instruction: (*CPU).${func},${padding}Addressing: (*CPU).${addressing[op.addressing]}},\n`);
}

outfile.write("}\n\n");
outfile.end();