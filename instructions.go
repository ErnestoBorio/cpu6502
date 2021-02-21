package cpu6502

const rightmostBit = 1
const leftmostBit = 1 << 7
const signBit = leftmostBit

func (cpu *CPU) calculateZeroNegative(value byte) {
	cpu.Status.Zero = value == 0
	cpu.Status.Negative = value&signBit != 0
}

func (cpu *CPU) lda(address uint16) {
	cpu.A = cpu.ReadMemory(address)
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) ldx(address uint16) {
	cpu.X = cpu.ReadMemory(address)
	cpu.calculateZeroNegative(cpu.X)
}

func (cpu *CPU) ldy(address uint16) {
	cpu.Y = cpu.ReadMemory(address)
	cpu.calculateZeroNegative(cpu.Y)
}

func (cpu *CPU) sta(address uint16) {
	cpu.WriteMemory(address, cpu.A)
}

func (cpu *CPU) stx(address uint16) {
	cpu.WriteMemory(address, cpu.X)
}

func (cpu *CPU) sty(address uint16) {
	cpu.WriteMemory(address, cpu.Y)
}

func (cpu *CPU) inx(uint16) {
	cpu.X++
	cpu.calculateZeroNegative(cpu.X)
}

func (cpu *CPU) dex(uint16) {
	cpu.X--
	cpu.calculateZeroNegative(cpu.X)
}

func (cpu *CPU) iny(uint16) {
	cpu.Y++
	cpu.calculateZeroNegative(cpu.Y)
}

func (cpu *CPU) dey(uint16) {
	cpu.Y--
	cpu.calculateZeroNegative(cpu.Y)
}

func (cpu *CPU) inc(address uint16) {
	value := cpu.ReadMemory(address) + 1
	cpu.WriteMemory(address, value)
	cpu.calculateZeroNegative(value)
}

func (cpu *CPU) dec(address uint16) {
	value := cpu.ReadMemory(address) - 1
	cpu.WriteMemory(address, value)
	cpu.calculateZeroNegative(value)
}

func (cpu *CPU) adc(address uint16) {
	value := cpu.ReadMemory(address)
	var sum byte
	if cpu.Status.Carry {
		sum = cpu.A + value + 1
		cpu.Status.Carry = sum <= cpu.A
	} else {
		sum = cpu.A + value
		cpu.Status.Carry = sum < cpu.A
	}
	cpu.Status.Overflow = ((cpu.A ^ byte(sum)) & (value ^ byte(sum)) & signBit) != 0 // @todo Test this
	cpu.A = sum
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) sbc(address uint16) {
	value := cpu.ReadMemory(address)
	var diff byte
	if cpu.Status.Carry {
		diff = cpu.A - value
		cpu.Status.Carry = diff < cpu.A
	} else {
		diff = cpu.A - value - 1
		cpu.Status.Carry = diff <= cpu.A
	}
	cpu.Status.Overflow = ((cpu.A ^ value) & (cpu.A ^ diff) & signBit) != 0
	cpu.A = diff
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) compare(register byte, address uint16) {
	value := cpu.ReadMemory(address)
	cpu.Status.Zero = register == value
	cpu.Status.Carry = register >= value
	cpu.Status.Negative = ((register - value) & signBit) != 0
}

func (cpu *CPU) cmp(address uint16) {
	cpu.compare(cpu.A, address)
}

func (cpu *CPU) cpx(address uint16) {
	cpu.compare(cpu.X, address)
}

func (cpu *CPU) cpy(address uint16) {
	cpu.compare(cpu.Y, address)
}

func (cpu *CPU) asla(uint16) {
	cpu.Status.Carry = (cpu.A & leftmostBit) == leftmostBit
	cpu.A = cpu.A << 1
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) asl(address uint16) {
	value := cpu.ReadMemory(address)
	cpu.Status.Carry = (value & leftmostBit) == leftmostBit
	value = value << 1
	cpu.calculateZeroNegative(value)
	cpu.WriteMemory(address, value)
}

func (cpu *CPU) lsra(uint16) {
	cpu.Status.Carry = (cpu.A & rightmostBit) == rightmostBit
	cpu.A = cpu.A >> 1
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) lsr(address uint16) {
	value := cpu.ReadMemory(address)
	cpu.Status.Carry = (value & rightmostBit) == rightmostBit
	value = value >> 1
	cpu.calculateZeroNegative(value)
	cpu.WriteMemory(address, value)
}

func (cpu *CPU) rola(uint16) {
	oldCarry := byte(0)
	if cpu.Status.Carry {
		oldCarry = 1
	}
	cpu.Status.Carry = (cpu.A & leftmostBit) > 0
	cpu.A = (cpu.A << 1) | oldCarry
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) rol(address uint16) {
	value := cpu.ReadMemory(address)
	oldCarry := byte(0)
	if cpu.Status.Carry {
		oldCarry = 1
	}
	cpu.Status.Carry = (value & leftmostBit) > 0
	value = (value << 1) | oldCarry
	cpu.calculateZeroNegative(value)
	cpu.WriteMemory(address, value)
}

func (cpu *CPU) rora(uint16) {
	oldCarry := byte(0)
	cpu.Status.Carry = (cpu.A & rightmostBit) > 0
	cpu.A = (cpu.A >> 1) | (oldCarry << 7)
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) ror(address uint16) {
	value := cpu.ReadMemory(address)
	oldCarry := byte(0)
	if cpu.Status.Carry {
		oldCarry = 1
	}
	cpu.Status.Carry = (value & rightmostBit) > 0
	value = (value >> 1) | (oldCarry << 7)
	cpu.calculateZeroNegative(value)
	cpu.WriteMemory(address, value)
}

func (cpu *CPU) tax(uint16) {
	cpu.X = cpu.A
	cpu.calculateZeroNegative(cpu.X)
}

func (cpu *CPU) txa(uint16) {
	cpu.A = cpu.X
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) tay(uint16) {
	cpu.Y = cpu.A
	cpu.calculateZeroNegative(cpu.Y)
}

func (cpu *CPU) tya(uint16) {
	cpu.A = cpu.Y
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) tsx(uint16) {
	cpu.X = cpu.Stack
	cpu.calculateZeroNegative(cpu.X)
}

func (cpu *CPU) txs(uint16) {
	cpu.Stack = cpu.X
}

func (cpu *CPU) and(address uint16) {
	cpu.A &= cpu.ReadMemory(address)
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) eor(address uint16) {
	cpu.A ^= cpu.ReadMemory(address)
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) ora(address uint16) {
	cpu.A |= cpu.ReadMemory(address)
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) bit(address uint16) {
	value := cpu.ReadMemory(address)
	cpu.Status.Zero = value&cpu.A == 0
	cpu.Status.Overflow = value&(1<<6) != 0 // bit #6
	cpu.Status.Negative = value&signBit != 0
}

func (cpu *CPU) branch(flag bool, condition bool) {
	jump := cpu.ReadMemory(cpu.PC)
	cpu.PC++
	if flag == condition {
		cpu.tmpCycles++
		oldPage := cpu.PC & 0xFF00 // high byte
		if jump&signBit > 0 {      // relative jump is negative
			cpu.PC += uint16(jump) - 0x100 // subtract jump's 2's complement
		} else {
			cpu.PC += uint16(jump)
		}
		// Branching crosses page boundary?
		if oldPage != cpu.PC&0xFF00 { // high byte
			cpu.tmpCycles++
		}
	}
}

func (cpu *CPU) beq(uint16) {
	cpu.branch(cpu.Status.Zero, true)
}

func (cpu *CPU) bne(uint16) {
	cpu.branch(cpu.Status.Zero, false)
}

func (cpu *CPU) bpl(uint16) {
	cpu.branch(cpu.Status.Negative, false)
}

func (cpu *CPU) bmi(uint16) {
	cpu.branch(cpu.Status.Negative, true)
}

func (cpu *CPU) bvs(uint16) {
	cpu.branch(cpu.Status.Overflow, true)
}

func (cpu *CPU) bvc(uint16) {
	cpu.branch(cpu.Status.Overflow, false)
}

func (cpu *CPU) bcs(uint16) {
	cpu.branch(cpu.Status.Carry, true)
}

func (cpu *CPU) bcc(uint16) {
	cpu.branch(cpu.Status.Carry, false)
}

func (cpu *CPU) push(value byte) {
	stack := 0x100 + uint16(cpu.Stack)
	cpu.WriteMemory(stack, value)
	cpu.Stack--
}

func (cpu *CPU) pull() byte {
	cpu.Stack++
	stack := 0x100 + uint16(cpu.Stack)
	return cpu.ReadMemory(stack)
}

func (cpu *CPU) packStatus() byte {
	flags := byte(1 << 5) // unused bit 5 always set
	if cpu.Status.Carry {
		flags |= 1
	}
	if cpu.Status.Zero {
		flags |= 1 << 1
	}
	if cpu.Status.NoInterrupt {
		flags |= 1 << 2
	}
	if cpu.Status.Decimal {
		flags |= 1 << 3
	}
	if cpu.Status.Overflow {
		flags |= 1 << 6
	}
	if cpu.Status.Negative {
		flags |= 1 << 7
	}
	return flags
}

func (cpu *CPU) php(uint16) {
	// http://nesdev.com/the 'B' flag & BRK instruction.txt
	// According to Brad Taylor PHP pushes the break flag as 1
	// Also the unused flag bit 5 is always 1
	flags := cpu.packStatus() | (1 << 4) | (1 << 5)
	cpu.push(flags)
}

func (cpu *CPU) pha(uint16) {
	cpu.push(cpu.A)
}

func (cpu *CPU) plp(uint16) {
	flags := cpu.pull()
	cpu.Status.Carry = flags&(1<<0) != 0
	cpu.Status.Zero = flags&(1<<1) != 0
	cpu.Status.NoInterrupt = flags&(1<<2) != 0
	cpu.Status.Decimal = flags&(1<<3) != 0
	cpu.Status.Overflow = flags&(1<<6) != 0
	cpu.Status.Negative = flags&(1<<7) != 0
}

func (cpu *CPU) pla(uint16) {
	cpu.A = cpu.pull()
	cpu.calculateZeroNegative(cpu.A)
}

// JMP
func (cpu *CPU) jumpAbsolute(address uint16) {
	cpu.PC = address
}

// JMP
func (cpu *CPU) jumpIndirect(address uint16) {
	cpu.PC = uint16(cpu.ReadMemory(address)) // low byte
	if (address & 0xFF) == 0xFF {            // address wraps around page
		address -= 0x100
	}
	address++
	cpu.PC |= uint16(cpu.ReadMemory(address)) << 8 // high byte
}

func (cpu *CPU) jsr(uint16) {
	// Return address is off by -1, pointing to JSR's last byte. Will be fixed on RTS
	returnAddress := cpu.PC + 1
	cpu.push(byte(returnAddress >> 8))   // address' high byte
	cpu.push(byte(returnAddress & 0xFF)) // address' low byte
	cpu.PC = cpu.getUint16(cpu.PC)       // Jump
}

func (cpu *CPU) rts(uint16) {
	cpu.PC = uint16(cpu.pull())
	cpu.PC |= uint16(cpu.pull()) << 8
	cpu.PC++ // Fix JSR's off by -1 return address
}

func (cpu *CPU) rti(uint16) {
	cpu.plp(0)
	cpu.PC = uint16(cpu.pull())
	cpu.PC |= uint16(cpu.pull()) << 8
}

// BRK (brk = true) and IRQ interrupt (brk = false)
func (cpu *CPU) irq(brk bool) {
	cpu.push(byte(cpu.PC >> 8))   // PC's high byte
	cpu.push(byte(cpu.PC & 0xFF)) // PC's low byte

	flags := cpu.packStatus()
	if brk { // set the break virtual flag
		flags |= 1 << 4
	}
	cpu.push(flags)
	// @todo: DarcNES unsets decimal flag here, other sources don't
	// "NMOS 6502 do not clear the decimal mode flag when an interrupt occurs"
	// Marat Fayzullin and others clear the decimal mode here
	cpu.Status.NoInterrupt = true
	cpu.PC = cpu.getUint16(0xFFFE) // Jump to IRQ/BRK vector
	cpu.tmpCycles = 7
}

func (cpu *CPU) brk(uint16) {
	cpu.irq(true)
}

func (cpu *CPU) nop(uint16) {}

func (cpu *CPU) nop2(uint16) {
	cpu.PC += 1
}

func (cpu *CPU) nop3(uint16) {
	cpu.PC += 2
}

// Undocumented opcodes
func (cpu *CPU) undoc(uint16) {}



/** @todo 
CLC SEC CLI SEI CLD SED CLV 
NOP_EA, 0x1A, 0x3A, 0x5A, 0x7A, 0xDA, 0xFA: // implied undocumented NOPs (1 byte wide)
0x0C, 0x1C, 0x3C, 0x5C, 0x7C, 0xDC, 0xFC: cpu.PC += 2 // absolute undoc NOPs (3 bytes wide)
0x80, 0x82, 0x89, 0xC2, 0xE2, 0x04, 0x14, 0x34, 0x44, 0x54, 0x64, 0x74, 0xD4, 0xF4: cpu.PC++ // immediate and zeropage undoc NOPs (2 bytes wide)
*/