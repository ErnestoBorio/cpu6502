package cpu6502

const rightmostBit = 1
const leftmostBit  = 1<<7
const bit6         = 1<<6
const signBit      = leftmostBit
const highByte     = 0xFF00
const lowByte      = 0x00FF

func (cpu *CPU) calculateZeroNegative(value byte) {
	cpu.Status.Zero = value == 0
	cpu.Status.Negative = value & signBit != 0
}

// LDA LDX LDY
func (cpu *CPU) loadReg(register *byte, address uint16) {
	*register = cpu.ReadMemory(address)
	cpu.calculateZeroNegative(*register)
}

// STA STX STY
func (cpu *CPU) storeReg(value byte, address uint16) {
	cpu.WriteMemory(address, value)
}

// INX DEX INY DEY
func (cpu *CPU) incDecReg(register *byte, delta int) {
	*register += byte(delta)
	cpu.calculateZeroNegative(*register)
}

// INC DEC
func (cpu *CPU) incDec(address uint16, delta int) {
	value := cpu.ReadMemory(address) + byte(delta)
	cpu.WriteMemory(address, value)
	cpu.calculateZeroNegative(value)
}

// ADC
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
	cpu.Status.Overflow = ((cpu.A ^ byte(sum)) & (value ^ byte(sum)) & signBit) != 0 // TODO Test this
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

// CMP CPX CPY
func (cpu *CPU) compare(register byte, address uint16) {
	value := cpu.ReadMemory(address)
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
func (cpu *CPU) asl(address uint16) {
	value := cpu.ReadMemory(address)
	cpu.Status.Carry = (value & leftmostBit) == leftmostBit
	value = value <<1
	cpu.calculateZeroNegative(value)
	cpu.WriteMemory(address, value)
}

// LSR A
func (cpu *CPU) lsra() {
	cpu.Status.Carry = (cpu.A & rightmostBit) == rightmostBit
	cpu.A = cpu.A >>1
	cpu.calculateZeroNegative(cpu.A)
}

// LSR
func (cpu *CPU) lsr(address uint16) {
	value := cpu.ReadMemory(address)
	cpu.Status.Carry = (value & rightmostBit) == rightmostBit
	value = value >>1
	cpu.calculateZeroNegative(value)
	cpu.WriteMemory(address, value)
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
func (cpu *CPU) rol(address uint16) {
	value := cpu.ReadMemory(address)
	oldCarry := byte(0)
	if cpu.Status.Carry {
		oldCarry = 1
	}
	cpu.Status.Carry = (value & leftmostBit) > 0
	value = ( value <<1 ) | oldCarry
	cpu.calculateZeroNegative(value)
	cpu.WriteMemory(address, value)
}

// ROR A
func (cpu *CPU) rora() {
	oldCarry := byte(0)
	cpu.Status.Carry = (cpu.A & rightmostBit) > 0
	cpu.A = (cpu.A >>1) | (oldCarry <<7)
	cpu.calculateZeroNegative(cpu.A)
}

// ROR
func (cpu *CPU) ror(address uint16) {
	value := cpu.ReadMemory(address)
	oldCarry := byte(0)
	if cpu.Status.Carry {
		oldCarry = 1
	}
	cpu.Status.Carry = (value & rightmostBit) > 0
	value = (value >>1) | (oldCarry <<7)
	cpu.calculateZeroNegative(value)
	cpu.WriteMemory(address, value)
}

// TAX TXA TAY TYA TSX
func (cpu *CPU) transferRegister(from byte, to *byte) {
	*to = from
	cpu.calculateZeroNegative(from)
}

// AND
func (cpu *CPU) and(address uint16) {
	cpu.A &= cpu.ReadMemory(address)
	cpu.calculateZeroNegative(cpu.A)
}

// EOR
func (cpu *CPU) eor(address uint16) {
	cpu.A ^= cpu.ReadMemory(address)
	cpu.calculateZeroNegative(cpu.A)
}

// ORA
func (cpu *CPU) ora(address uint16) {
	cpu.A |= cpu.ReadMemory(address)
	cpu.calculateZeroNegative(cpu.A)
}

// BIT
func (cpu *CPU) bit(address uint16) {
	value := cpu.ReadMemory(address)
	cpu.Status.Zero =     value & cpu.A == 0
	cpu.Status.Overflow = value & bit6 != 0
	cpu.Status.Negative = value & signBit != 0
}

// BEQ BNE BPL BMI BVS BVC BCS BCC
func (cpu *CPU) branch(flag bool, condition bool) {
	jump := cpu.ReadMemory(cpu.PC)
	cpu.PC++
	if flag == condition {
		cpu.cycles++
		oldPage := cpu.PC & highByte
		if jump & signBit > 0 { // relative jump is negative
			cpu.PC += uint16(jump) - 0x100 // subtract jump's 2's complement
		} else {
			cpu.PC += uint16(jump)
		}
		// Branching crosses page boundary?
		if oldPage != cpu.PC & highByte {
			cpu.cycles++
		}
	}
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

func (cpu *CPU) getUint16(address uint16) uint16 {
	value := uint16( cpu.ReadMemory(address))
	value |= uint16( cpu.ReadMemory(address+1)) <<8
	return value
}

// JMP
func (cpu *CPU) jumpAbsolute() {
	cpu.PC = cpu.getUint16(cpu.PC)
}

// JMP
func (cpu *CPU) jumpIndirect() {
	pointer := cpu.getUint16(cpu.PC)
	cpu.PC = uint16(cpu.ReadMemory(pointer)) // low byte
	if (pointer & lowByte) == 0xFF { // address wraps around page
		pointer -= 0x100
	}
	pointer++
	cpu.PC |= uint16(cpu.ReadMemory(pointer)) <<8 // high byte
}

// JSR
func (cpu *CPU) jsr() {
	// return address is off by -1, pointing to JSR's last byte.
	// Will be fixed on RTS
	returnAddress := cpu.PC + 1
	cpu.push( byte( returnAddress >>8)) // address' high byte
	cpu.push( byte( returnAddress & lowByte)) // address' low byte
	cpu.PC = cpu.getUint16(cpu.PC) // Jump
}

// RTS
func (cpu *CPU) rts() {
	cpu.PC = uint16(cpu.pull())
	cpu.PC |= uint16(cpu.pull()) <<8
	cpu.PC++ // Fix JSR's off by -1 return address
}

// RTI
func (cpu *CPU) rti() {
	cpu.plp()
	cpu.PC = uint16(cpu.pull())
	cpu.PC |= uint16(cpu.pull()) <<8
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
	cpu.PC = cpu.getUint16(0xFFFE) // Jump to IRQ/BRK vector
	cpu.cycles = 7
}