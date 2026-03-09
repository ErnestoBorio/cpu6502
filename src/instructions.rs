use super::CPU;

impl CPU {
	pub(super) fn push(&mut self, value: u8) {
		let stack_addr = 0x100 + self.stack as u16;
		(self.write_memory)(stack_addr, value);
		self.stack -= 1;
	}

	fn pull(&mut self) -> u8 {
		self.stack += 1;
		let stack_addr = 0x100 + self.stack as u16;
		(self.read_memory)(stack_addr)
	}

	pub(super) fn pack_status(&mut self) -> u8 {
		(self.status.negative as u8) << 7
			| (self.status.overflow as u8) << 6
			| (1 << 5)
			| (self.status.no_interrupt as u8) << 4
			| (self.status.decimal as u8) << 3
			| (self.status.carry as u8) << 0
	}

	// If executing a BRK instruction, the brk parameter should be true, otherwise false for an IRQ.
	pub(super) fn irq(&mut self, brk: bool) -> u32 {
		self.push((self.pc >> 8) as u8);
		self.push((self.pc & 0xFF) as u8);
		let flags = self.pack_status() | ((brk as u8) << 4);
		self.push(flags);
		self.status.no_interrupt = true;
		self.pc = self.get_u16(0xFFFE);
		return 7; // cycles consumed by IRQ or BRK
	}

	fn brk(&mut self) {
		self.irq(true);
	}
}
