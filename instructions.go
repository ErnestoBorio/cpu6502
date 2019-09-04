package cpu6502

const rightmostBit = 1
const leftmostBit  = 1<<7
const bit6         = 1<<6
const signBit      = leftmostBit
const highByte     = 0xFF00
const lowByte      = 0xFF

func (cpu *CPU) calculateZeroNegative(value byte) {
	cpu.Status.Zero = value == 0
	cpu.Status.Negative = value & signBit != 0
}

// LDA LDX LDY
func (cpu *CPU) loadReg(register *byte, value byte) {
	*register = value
	cpu.calculateZeroNegative(value)
}

// STA STX STY
func (cpu *CPU) storeReg(value byte, address word) {
	(*cpu.writeMemory[address])(address, value)
}

// INX DEX INY DEY
func (cpu *CPU) incDecReg(register *byte, delta int) {
	*register += byte(delta)
	cpu.calculateZeroNegative(*register)
}

// INC DEC
func (cpu *CPU) incDec(address word, delta int) {
	value := (*cpu.readMemory[address])(address) + byte(delta)
	(*cpu.writeMemory[address])(address, value)
	cpu.calculateZeroNegative(value)
}

// ADC
func (cpu *CPU) adc(value byte) {
	var sum byte
	if cpu.Status.Carry {
		sum = cpu.A + value + 1
		cpu.Status.Carry = sum <= cpu.A
	} else {
		sum = cpu.A + value
		cpu.Status.Carry = sum < cpu.A
	}
	cpu.Status.Overflow = ((cpu.A ^ byte(sum)) & (value ^ byte(sum)) & signBit) != 0 // TODO Test this
	cpu.A = sum
	cpu.calculateZeroNegative(cpu.A)
}

//SBC TODO TEST!
func (cpu *CPU) sbc(value byte) {
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

// CMP CPX CPY
func (cpu *CPU) compare(register, value byte) {
	cpu.Status.Zero  = register == value
	cpu.Status.Carry = register >= value
	cpu.Status.Negative = ((register - value) & signBit) != 0
}

// ASL A
func (cpu *CPU) asla() {
	cpu.Status.Carry = (cpu.A & leftmostBit) == leftmostBit
	cpu.A = cpu.A <<1
	cpu.calculateZeroNegative(cpu.A)
}

// ASL
func (cpu *CPU) asl(address word) {
	value := (*cpu.readMemory[address])(address)
	cpu.Status.Carry = (value & leftmostBit) == leftmostBit
	value = value <<1
	cpu.calculateZeroNegative(value)
	(*cpu.writeMemory[address])(address, value)
}

// LSR A
func (cpu *CPU) lsra() {
	cpu.Status.Carry = (cpu.A & rightmostBit) == rightmostBit
	cpu.A = cpu.A >>1
	cpu.calculateZeroNegative(cpu.A)
}

// LSR
func (cpu *CPU) lsr(address word) {
	value := (*cpu.readMemory[address])(address)
	cpu.Status.Carry = (value & rightmostBit) == rightmostBit
	value = value >>1
	cpu.calculateZeroNegative(value)
	(*cpu.writeMemory[address])(address, value)
}

// ROL A
func (cpu *CPU) rola() {
	oldCarry := byte(0)
	if cpu.Status.Carry {
		oldCarry = 1
	}
	cpu.Status.Carry = (cpu.A & leftmostBit) > 0
	cpu.A = ( cpu.A <<1 ) | oldCarry
	cpu.calculateZeroNegative(cpu.A)
}

// ROL
func (cpu *CPU) rol(address word) {
	value := (*cpu.readMemory[address])(address)
	oldCarry := byte(0)
	if cpu.Status.Carry {
		oldCarry = 1
	}
	cpu.Status.Carry = (value & leftmostBit) > 0
	value = ( value <<1 ) | oldCarry
	cpu.calculateZeroNegative(value)
	(*cpu.writeMemory[address])(address, value)
}

// ROR A
func (cpu *CPU) rora() {
	oldCarry := byte(0)
	cpu.Status.Carry = (cpu.A & rightmostBit) > 0
	cpu.A = (cpu.A >>1) | (oldCarry <<7)
	cpu.calculateZeroNegative(cpu.A)
}

// ROR
func (cpu *CPU) ror(address word) {
	value := (*cpu.readMemory[address])(address)
	oldCarry := byte(0)
	if cpu.Status.Carry {
		oldCarry = 1
	}
	cpu.Status.Carry = (value & rightmostBit) > 0
	value = (value >>1) | (oldCarry <<7)
	cpu.calculateZeroNegative(value)
	(*cpu.writeMemory[address])(address, value)
}

// TAX TXA TAY TYA TSX
func (cpu *CPU) transferRegister(from byte, to *byte) {
	*to = from
	cpu.calculateZeroNegative(from)
}

// AND
func (cpu *CPU) and(value byte) {
	cpu.A &= value
	cpu.calculateZeroNegative(cpu.A)
}

// EOR
func (cpu *CPU) eor(value byte) {
	cpu.A ^= value
	cpu.calculateZeroNegative(cpu.A)
}

// ORA
func (cpu *CPU) ora(value byte) {
	cpu.A |= value
	cpu.calculateZeroNegative(cpu.A)
}

// BIT
func (cpu *CPU) bit(value byte) {
	cpu.Status.Zero =     value & cpu.A == 0
	cpu.Status.Overflow = value & bit6 != 0
	cpu.Status.Negative = value & signBit != 0
}

// BEQ BNE BPL BMI BVS BVC BCS BCC
func (cpu *CPU) branch(flag bool, condition bool, jump byte) {
	if flag == condition {
		cpu.cycles++
		oldPage := cpu.PC & highByte
		if jump & signBit > 0 { // relative jump is negative
			cpu.PC += word(jump) - 0x100 // subtract jump's 2's complement
		} else {
			cpu.PC += word(jump)
		}
		// Branching crosses page boundary?
		if oldPage != cpu.PC & highByte {
			cpu.cycles++
		}
	}
}

func (cpu *CPU) push(value byte) {
	stack := 0x100 + word(cpu.Stack)
	(*cpu.writeMemory[stack])(stack, value)
	cpu.Stack--
}

func (cpu *CPU) pull() byte {
	cpu.Stack++
	stack := 0x100 + word(cpu.Stack)
	return (*cpu.readMemory[stack])(stack)
}

// PLA
func (cpu *CPU) pla() {
	cpu.A = cpu.pull()
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *CPU) packStatus() byte {
	flags := byte(1<<5) // unused bit 5 always set
	if cpu.Status.Carry {
		flags |= 1
	}
	if cpu.Status.Zero {
		flags |= 1<<1
	}
	if cpu.Status.NoInterrupt {
		flags |= 1<<2
	}
	if cpu.Status.Decimal {
		flags |= 1<<3
	}
	if cpu.Status.Overflow {
		flags |= 1<<6
	}
	if cpu.Status.Negative {
		flags |= 1<<7
	}
	return flags
}

// PHP
func (cpu *CPU) php() {
	// http://nesdev.com/the 'B' flag & BRK instruction.txt
	// According to Brad Taylor PHP pushes the break flag as 1
	// Also the unused flag bit 5 is always 1
	flags := cpu.packStatus() | (1<<4) | (1<<5)
	cpu.push(flags)
}

// PLP
func (cpu *CPU) plp() {
	flags := cpu.pull()
	cpu.Status.Carry       = flags & (1<<0) != 0
	cpu.Status.Zero        = flags & (1<<1) != 0
	cpu.Status.NoInterrupt = flags & (1<<2) != 0
	cpu.Status.Decimal     = flags & (1<<3) != 0
	cpu.Status.Overflow    = flags & (1<<6) != 0
	cpu.Status.Negative    = flags & (1<<7) != 0
}

func (cpu *CPU) getWord(address word) word {
	value := word( (*cpu.readMemory[address])(address))
	value |= word( (*cpu.readMemory[address+1])(address+1)) <<8
	return value
}

// JMP
func (cpu *CPU) jumpAbsolute() {
	cpu.PC = cpu.getWord(cpu.PC)
}

// JMP
func (cpu *CPU) jumpIndirect() {
	pointer := cpu.getWord(cpu.PC)
	cpu.PC = word((*cpu.readMemory[pointer])(pointer)) // low byte
	if (pointer & lowByte) == 0xFF { // address wraps around page
		pointer -= 0x100
	}
	pointer++
	cpu.PC |= word((*cpu.readMemory[pointer])(pointer)) <<8 // high byte
}

// JSR
func (cpu *CPU) jsr() {
	// return address is off by -1, pointing to JSR's last byte.
	// Will be fixed on RTS
	returnAddress := cpu.PC + 1
	cpu.push( byte( returnAddress >>8)) // address' high byte
	cpu.push( byte( returnAddress & lowByte)) // address' low byte
	cpu.PC = cpu.getWord(cpu.PC) // Jump
}

// RTS
func (cpu *CPU) rts() {
	cpu.PC = word(cpu.pull())
	cpu.PC |= word(cpu.pull()) <<8
	cpu.PC++ // Fix JSR's off by -1 return address
}

// RTI
func (cpu *CPU) rti() {
	cpu.plp()
	cpu.PC = word(cpu.pull())
	cpu.PC |= word(cpu.pull()) <<8
}

// BRK (brk = true) and IRQ interrupt (brk = false)
func (cpu *CPU) irq(brk bool) {
	cpu.push( byte( cpu.PC >>8)) // PC's high byte
	cpu.push( byte( cpu.PC & lowByte)) // PC's low byte
	
	flags := cpu.packStatus()
	if brk { // set the break virtual flag
		flags |= 1<<4
	}
	cpu.push(flags)
	// TODO: DarcNES unsets decimal flag here, other sources don't
	// "NMOS 6502 do not clear the decimal mode flag when an interrupt occurs"
	// Marat Fayzullin and others clear the decimal mode here
	cpu.Status.NoInterrupt = true
	cpu.PC = cpu.getWord(0xFFFE) // Jump to IRQ/BRK vector
	cpu.cycles = 7
}