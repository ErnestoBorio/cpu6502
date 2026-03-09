pub struct Status {
	zero: bool,
	carry: bool,
	decimal: bool,
	overflow: bool,
	negative: bool,
	no_interrupt: bool,
}

type MemoryReader = fn(u16) -> u8;
type MemoryWriter = fn(u16, u8);

pub struct CPU {
	pub pc: u16,
	pub stack: u8,
	pub a: u8,
	pub x: u8,
	pub y: u8,
	pub status: Status,

	pub read_memory: MemoryReader,
	pub write_memory: MemoryWriter,
}

impl CPU {
	pub fn new(read_memory: MemoryReader, write_memory: MemoryWriter) -> Self {
		Self {
			pc: 0,
			stack: 0,
			a: 0,
			x: 0,
			y: 0,
			status: Status {
				zero: false,
				carry: false,
				decimal: false,
				overflow: false,
				negative: false,
				no_interrupt: false,
			},
			read_memory,
			write_memory,
		}
	}

	fn get_u16(&mut self, address: u16) -> u16 {
		let low = (self.read_memory)(address) as u16;
		let high = (self.read_memory)(address + 1) as u16;
		(high << 8) | low
	}

	// Jump to the address where the reset vector points to
	pub fn reset(&mut self) {
		self.stack = 0xFD;
		self.pc = self.get_u16(0xFFFC);
	}

	// Trigger an external IRQ interrupt
	pub fn interrupt(&mut self) {
		if !self.status.no_interrupt {
			self.irq(false);
		}
	}

	pub fn nmi(&mut self) -> u32 {
		self.push((self.pc >> 8) as u8);
		self.push((self.pc & 0xFF) as u8);
		self.push(self.pack_status());

		// @todo: Marat Fayzullin and others clear the decimal mode here (NES specific?)
		self.status.no_interrupt = true; // @todo: Marat Fayzullin doesn't do this (NES specific?)

		self.pc = self.get_u16(0xFFFA);

		return 7;
	}

	pub fn step(&mut self) -> u32 {
		// Fetch operation code from current PC address
		let opcode = (self.read_memory)(self.pc);
		// Advance PC for either first argument byte or next instruction
		self.pc += 1;
		// Count basic instruction cycles, they can be incremented afterwards in certain conditions

		// Call the appropriate instruction and addressing mode for the fetched opcode
		// @todo: byte following opcode is read in advance, possibly causing side effects. Test it.
	}
}

mod instructions;
mod opcodes;
