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

	fn calculate_zero_negative(&mut self, value: u8) {
		self.status.zero = value == 0;
		self.status.negative = value & 0b1000000 != 0;
	}

	fn lda(&mut self, address: u16) {
		self.a = (self.read_memory)(address);
		self.calculate_zero_negative(self.a);
	}

	fn ldx(&mut self, address: u16) {
		self.x = (self.read_memory)(address);
		self.calculate_zero_negative(self.x);
	}

	fn ldy(&mut self, address: u16) {
		self.y = (self.read_memory)(address);
		self.calculate_zero_negative(self.y);
	}

	fn sta(&mut self, address: u16) {
		(self.write_memory)(address, self.a);
	}

	fn stx(&mut self, address: u16) {
		(self.write_memory)(address, self.x);
	}

	fn sty(&mut self, address: u16) {
		(self.write_memory)(address, self.y);
	}

	fn inx(&mut self) {
		self.x = self.x.wrapping_add(1);
		self.calculate_zero_negative(self.x);
	}

	fn dex(&mut self) {
		self.x = self.x.wrapping_sub(1);
		self.calculate_zero_negative(self.x);
	}

	fn iny(&mut self) {
		self.y = self.y.wrapping_add(1);
		self.calculate_zero_negative(self.y);
	}

	fn dey(&mut self) {
		self.y = self.y.wrapping_sub(1);
		self.calculate_zero_negative(self.y);
	}

	fn inc(&mut self, address: u16) {
		let value = (self.read_memory)(address).wrapping_add(1);
		(self.write_memory)(address, value);
		self.calculate_zero_negative(value);
	}

	fn dec(&mut self, address: u16) {
		let value = (self.read_memory)(address).wrapping_sub(1);
		(self.write_memory)(address, value);
		self.calculate_zero_negative(value);
	}

	fn adc(&mut self, address: u16) {
		let value = (self.read_memory)(address);
		let sum = if self.status.carry {
			let (sum, _) = self.a.overflowing_add(value);
			let (sum, carry) = sum.overflowing_add(1);
			self.status.carry = carry || sum <= self.a;
			sum
		} else {
			let (sum, carry) = self.a.overflowing_add(value);
			self.status.carry = carry;
			sum
		};
		self.status.overflow = ((self.a ^ sum) & (value ^ sum) & 0x80) != 0;
		self.a = sum;
		self.calculate_zero_negative(self.a);
	}

	fn sbc(&mut self, address: u16) {
		let value = (self.read_memory)(address);
		let diff = if self.status.carry {
			let (diff, carry) = self.a.overflowing_sub(value);
			self.status.carry = !carry;
			diff
		} else {
			let (diff, _) = self.a.overflowing_sub(value);
			let (diff, carry) = diff.overflowing_sub(1);
			self.status.carry = !carry;
			diff
		};
		self.status.overflow = ((self.a ^ value) & (self.a ^ diff) & 0x80) != 0;
		self.a = diff;
		self.calculate_zero_negative(self.a);
	}

	fn compare(&mut self, register: u8, address: u16) {
		let value = (self.read_memory)(address);
		self.status.zero = register == value;
		self.status.carry = register >= value;
		self.status.negative = (register.wrapping_sub(value)) & 0x80 != 0;
	}

	fn cmp(&mut self, address: u16) {
		self.compare(self.a, address);
	}

	fn cpx(&mut self, address: u16) {
		self.compare(self.x, address);
	}

	fn cpy(&mut self, address: u16) {
		self.compare(self.y, address);
	}

	fn asla(&mut self) {
		self.status.carry = (self.a & 0x80) != 0;
		self.a = self.a << 1;
		self.calculate_zero_negative(self.a);
	}

	fn asl(&mut self, address: u16) {
		let mut value = (self.read_memory)(address);
		self.status.carry = (value & 0x80) != 0;
		value = value << 1;
		self.calculate_zero_negative(value);
		(self.write_memory)(address, value);
	}

	fn lsra(&mut self) {
		self.status.carry = (self.a & 0x01) != 0;
		self.a = self.a >> 1;
		self.calculate_zero_negative(self.a);
	}

	fn lsr(&mut self, address: u16) {
		let mut value = (self.read_memory)(address);
		self.status.carry = (value & 0x01) != 0;
		value = value >> 1;
		self.calculate_zero_negative(value);
		(self.write_memory)(address, value);
	}

	fn rola(&mut self) {
		let old_carry = if self.status.carry { 1 } else { 0 };
		self.status.carry = (self.a & 0x80) != 0;
		self.a = (self.a << 1) | old_carry;
		self.calculate_zero_negative(self.a);
	}

	fn rol(&mut self, address: u16) {
		let mut value = (self.read_memory)(address);
		let old_carry = if self.status.carry { 1 } else { 0 };
		self.status.carry = (value & 0x80) != 0;
		value = (value << 1) | old_carry;
		self.calculate_zero_negative(value);
		(self.write_memory)(address, value);
	}

	fn rora(&mut self) {
		let old_carry = if self.status.carry { 0x80 } else { 0 };
		self.status.carry = (self.a & 0x01) != 0;
		self.a = (self.a >> 1) | old_carry;
		self.calculate_zero_negative(self.a);
	}

	fn ror(&mut self, address: u16) {
		let mut value = (self.read_memory)(address);
		let old_carry = if self.status.carry { 0x80 } else { 0 };
		self.status.carry = (value & 0x01) != 0;
		value = (value >> 1) | old_carry;
		self.calculate_zero_negative(value);
		(self.write_memory)(address, value);
	}

	fn tax(&mut self) {
		self.x = self.a;
		self.calculate_zero_negative(self.x);
	}

	fn txa(&mut self) {
		self.a = self.x;
		self.calculate_zero_negative(self.a);
	}

	fn tay(&mut self) {
		self.y = self.a;
		self.calculate_zero_negative(self.y);
	}

	fn tya(&mut self) {
		self.a = self.y;
		self.calculate_zero_negative(self.a);
	}

	fn tsx(&mut self) {
		self.x = self.stack;
		self.calculate_zero_negative(self.x);
	}

	fn txs(&mut self) {
		self.stack = self.x;
	}

	fn and(&mut self, address: u16) {
		self.a &= (self.read_memory)(address);
		self.calculate_zero_negative(self.a);
	}

	fn eor(&mut self, address: u16) {
		self.a ^= (self.read_memory)(address);
		self.calculate_zero_negative(self.a);
	}

	fn ora(&mut self, address: u16) {
		self.a |= (self.read_memory)(address);
		self.calculate_zero_negative(self.a);
	}

	fn bit(&mut self, address: u16) {
		let value = (self.read_memory)(address);
		self.status.zero = (value & self.a) == 0;
		self.status.overflow = (value & 0x40) != 0;
		self.status.negative = (value & 0x80) != 0;
	}

	fn php(&mut self) {
		let flags = self.pack_status() | (1 << 4) | (1 << 5);
		self.push(flags);
	}

	fn pha(&mut self) {
		self.push(self.a);
	}

	fn plp(&mut self) {
		let flags = self.pull();
		self.status.carry = (flags & (1 << 0)) != 0;
		self.status.zero = (flags & (1 << 1)) != 0;
		self.status.no_interrupt = (flags & (1 << 2)) != 0;
		self.status.decimal = (flags & (1 << 3)) != 0;
		self.status.overflow = (flags & (1 << 6)) != 0;
		self.status.negative = (flags & (1 << 7)) != 0;
	}

	fn pla(&mut self) {
		self.a = self.pull();
		self.calculate_zero_negative(self.a);
	}

	fn jmp(&mut self, address: u16) {
		self.pc = address;
	}

	fn jsr(&mut self, address: u16) {
		let return_address = self.pc.wrapping_sub(1);
		self.push((return_address >> 8) as u8);
		self.push((return_address & 0xFF) as u8);
		self.pc = address;
	}

	fn rts(&mut self) {
		self.pc = self.pull() as u16;
		self.pc |= (self.pull() as u16) << 8;
		self.pc = self.pc.wrapping_add(1);
	}

	fn rti(&mut self) {
		self.plp();
		self.pc = self.pull() as u16;
		self.pc |= (self.pull() as u16) << 8;
	}

	fn clc(&mut self) {
		self.status.carry = false;
	}

	fn sec(&mut self) {
		self.status.carry = true;
	}

	fn cli(&mut self) {
		self.status.no_interrupt = false;
	}

	fn sei(&mut self) {
		self.status.no_interrupt = true;
	}

	fn cld(&mut self) {
		self.status.decimal = false;
	}

	fn sed(&mut self) {
		self.status.decimal = true;
	}

	fn clv(&mut self) {
		self.status.overflow = false;
	}

	fn nop(&mut self) {}

	fn nop2(&mut self) {
		self.pc = self.pc.wrapping_add(1);
	}

	fn nop3(&mut self) {
		self.pc = self.pc.wrapping_add(2);
	}

	fn undoc(&mut self) {}
}
