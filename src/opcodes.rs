use crate::opcodes;

use super::CPU;

struct Operation {
	opcode: u8,
	length: u8,
	cycles: u8,
	documented: bool,
	addressing: fn(&mut CPU) -> u16,
	instruction: fn(&mut CPU, u16),
	mnemonic: String,
	addressing_label: String,
}

static opcodes: [Operation; 0x1] = [Operation {
	opcode: 0x00,
	length: 0,
	cycles: 7,
	documented: true,
	addressing: CPU::implied,
	instruction: CPU::brk,
	mnemonic: "BRK".to_string(),
	addressing_label: "Implied".to_string(),
}];
