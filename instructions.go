package cpu6502

const rightmostBit = 1
const leftmostBit  = 1<<7
const bit6         = 1<<6
const signBit      = leftmostBit
const highByte     = 0xFF00
const lowByte      = 0xFF

func (cpu *Cpu) calculateZeroNegative(value byte) {
	cpu.Status.Zero = value == 0
	cpu.Status.Negative = value & signBit != 0
}

// LDA LDX LDY
func (cpu *Cpu) loadReg(register *byte, value byte) {
	*register = value
	cpu.calculateZeroNegative(value)
}

// STA STX STY
func (cpu *Cpu) storeReg(value byte, address word) {
	cpu.writeMemory[address](address, value)
}

// INX DEX INY DEY
func (cpu *Cpu) incDecReg(register *byte, delta int) {
	*register += byte(delta)
	cpu.calculateZeroNegative(*register)
	cpu.PC++
}

// INC DEC
func (cpu *Cpu) incDec(address word, delta int) {
	value := cpu.readMemory[address](address) + byte(delta)
	cpu.writeMemory[address](address, value)
	cpu.calculateZeroNegative(value)
}

// ADC
func (cpu *Cpu) adc(value byte) {
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
func (cpu *Cpu) sbc(value byte) {
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
func (cpu *Cpu) compare(register, value byte) {
	cpu.Status.Zero  = register == value
	cpu.Status.Carry = register >= value
	cpu.Status.Negative = ((register - value) & signBit) != 0
}

// ASL A
func (cpu *Cpu) asla() {
	cpu.Status.Carry = (cpu.A & leftmostBit) == leftmostBit
	cpu.A = cpu.A <<1
	cpu.calculateZeroNegative(cpu.A)
	cpu.PC++
}

// ASL
func (cpu *Cpu) asl(address word) {
	value := cpu.readMemory[address](address)
	cpu.Status.Carry = (value & leftmostBit) == leftmostBit
	value = value <<1
	cpu.calculateZeroNegative(value)
	cpu.writeMemory[address](address, value)
}

// LSR A
func (cpu *Cpu) lsra() {
	cpu.Status.Carry = (cpu.A & rightmostBit) == rightmostBit
	cpu.A = cpu.A >>1
	cpu.calculateZeroNegative(cpu.A)
	cpu.PC++
}

// LSR
func (cpu *Cpu) lsr(address word) {
	value := cpu.readMemory[address](address)
	cpu.Status.Carry = (value & rightmostBit) == rightmostBit
	value = value >>1
	cpu.calculateZeroNegative(value)
	cpu.writeMemory[address](address, value)
}

// ROL A
func (cpu *Cpu) rola() {
	oldCarry := byte(0)
	if cpu.Status.Carry {
		oldCarry = 1
	}
	cpu.Status.Carry = (cpu.A & leftmostBit) > 0
	cpu.A = ( cpu.A <<1 ) | oldCarry
	cpu.calculateZeroNegative(cpu.A)
	cpu.PC++
}

// ROL
func (cpu *Cpu) rol(address word) {
	value := cpu.readMemory[address](address)
	oldCarry := byte(0)
	if cpu.Status.Carry {
		oldCarry = 1
	}
	cpu.Status.Carry = (value & leftmostBit) > 0
	value = ( value <<1 ) | oldCarry
	cpu.calculateZeroNegative(value)
	cpu.writeMemory[address](address, value)
}

// ROR A
func (cpu *Cpu) rora() {
	oldCarry := byte(0)
	cpu.Status.Carry = (cpu.A & rightmostBit) > 0
	cpu.A = (cpu.A >>1) | (oldCarry <<7)
	cpu.calculateZeroNegative(cpu.A)
	cpu.PC++
}

// ROR
func (cpu *Cpu) ror(address word) {
	value := cpu.readMemory[address](address)
	oldCarry := byte(0)
	cpu.Status.Carry = (value & rightmostBit) > 0
	value = (value >>1) | (oldCarry <<7)
	cpu.calculateZeroNegative(value)
	cpu.writeMemory[address](address, value)
}

// TAX TXA TAY TYA TSX
func (cpu *Cpu) transferRegister(from byte, to *byte) {
	*to = from
	cpu.calculateZeroNegative(from)
}

// AND
func (cpu *Cpu) and(value byte) {
	cpu.A &= value
	cpu.calculateZeroNegative(cpu.A)
}

// EOR
func (cpu *Cpu) eor(value byte) {
	cpu.A ^= value
	cpu.calculateZeroNegative(cpu.A)
}

// ORA
func (cpu *Cpu) ora(value byte) {
	cpu.A |= value
	cpu.calculateZeroNegative(cpu.A)
}

// BIT
func (cpu *Cpu) bit(value byte) {
	cpu.Status.Zero =     value & cpu.A == 0
	cpu.Status.Overflow = value & bit6 != 0
	cpu.Status.Negative = value & signBit != 0
}

// BEQ BNE BPL BMI BVS BVC BCS BCC
func (cpu *Cpu) branch(flag bool, condition bool, jump byte) {
	cpu.PC += 2 // The branch is relative to next instruction's address
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

func (cpu *Cpu) push(value byte) {
	stack := 0x100 + word(cpu.Stack)
	cpu.writeMemory[stack](stack, value)
	cpu.Stack--
}

func (cpu *Cpu) pull() byte {
	cpu.Stack++
	stack := 0x100 + word(cpu.Stack)
	return cpu.readMemory[stack](stack)
}

// PLA
func (cpu *Cpu) pla() {
	cpu.A = cpu.pull()
	cpu.calculateZeroNegative(cpu.A)
}

func (cpu *Cpu) packStatus() byte {
	flags := byte(1<<5) // unused bit 5 always set
	if cpu.Status.Carry {
		flags |= 1
	}
	if cpu.Status.Zero {
		flags |= 1<<1
	}
	if cpu.Status.IntDis {
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
func (cpu *Cpu) php() {
	// http://nesdev.com/the 'B' flag & BRK instruction.txt
	// According to Brad Taylor PHP pushes the break flag as 1
	// Also the unused flag bit 5 is always 1
	flags := cpu.packStatus() | (1<<4) | (1<<5)
	cpu.push(flags)
}

// PLP
func (cpu *Cpu) plp() {
	flags := cpu.pull()
	cpu.Status.Carry    = flags & (1<<0) != 0
	cpu.Status.Zero     = flags & (1<<1) != 0
	cpu.Status.IntDis   = flags & (1<<2) != 0
	cpu.Status.Decimal  = flags & (1<<3) != 0
	cpu.Status.Overflow = flags & (1<<6) != 0
	cpu.Status.Negative = flags & (1<<7) != 0
}

func (cpu *Cpu) getWord(address word) word {
	value := word( cpu.readMemory[address](address))
	value |= word( cpu.readMemory[address+1](address+1)) <<8
	return value
}

func (cpu *Cpu) jumpAbsolute() {
	cpu.PC = cpu.getWord(cpu.PC)
}

func (cpu *Cpu) jumpIndirect() {
	pointer := cpu.getWord(cpu.PC)
	cpu.PC = cpu.getWord(pointer)
}

// JSR
func (cpu *Cpu) jsr() {
	// return address is off by -1, pointing to JSR's last byte.
	// Will be fixed on RTS
	returnAddress := cpu.PC + 1
	cpu.push( byte( returnAddress >>8)) // address' high byte
	cpu.push( byte( returnAddress & lowByte)) // address' low byte
	cpu.PC = cpu.getWord(cpu.PC) // Jump
}

// RTS
func (cpu *Cpu) rts() {
	cpu.PC = word(cpu.pull())
	cpu.PC |= word(cpu.pull()) <<8
	cpu.PC++ // Fix JSR's off by -1 return address
}

// RTI
func (cpu *Cpu) rti() {
	cpu.plp()
	cpu.PC = word(cpu.pull())
	cpu.PC |= word(cpu.pull()) <<8
}

// BRK
func (cpu *Cpu) irq(brk bool) {
	cpu.cycles = 7
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
	cpu.Status.IntDis = true
	cpu.PC = cpu.getWord(0xFFFE) // Jump to IRQ/BRK vector
}

func (cpu *Cpu) IRQ() {
	if ! cpu.Status.IntDis {
		cpu.irq(false)
	}
}

func (cpu *Cpu) NMI() {
	cpu.cycles = 7
	cpu.push( byte( cpu.PC >>8)) // PC's high byte
	cpu.push( byte( cpu.PC & lowByte)) // PC's low byte
	cpu.push(cpu.packStatus())
	// Marat Fayzullin and others clear the decimal mode here
	cpu.Status.IntDis = true // TODO: Marat Fayzullin doesn't do this
	cpu.PC = cpu.getWord(0xFFFA) // Jump to NMI vector
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
	2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7} //F0
	
// 0 = illegal, 1 = illegal NOP (semi-legal), 2 = legal
var opcodeLegality = [0x100] byte {
//  0 1 2 3 4 5 6 7 8 9 A B C D E F
	2,2,0,0,1,2,2,0,2,2,2,0,1,2,2,0, //00
	2,2,0,0,1,2,2,0,2,2,1,0,1,2,2,0, //10
	2,2,0,0,2,2,2,0,2,2,2,0,2,2,2,0, //20
	2,2,0,0,1,2,2,0,2,2,1,0,1,2,2,0, //30
	2,2,0,0,1,2,2,0,2,2,2,0,2,2,2,0, //40
	2,2,0,0,1,2,2,0,2,2,1,0,1,2,2,0, //50
	2,2,0,0,1,2,2,0,2,2,2,0,2,2,2,0, //60
	2,2,0,0,1,2,2,0,2,2,1,0,1,2,2,0, //70
	1,2,1,0,2,2,2,0,2,1,2,0,2,2,2,0, //80
	2,2,0,0,2,2,2,0,2,2,2,0,0,2,0,0, //90
	2,2,2,0,2,2,2,0,2,2,2,0,2,2,2,0, //A0
	2,2,0,0,2,2,2,0,2,2,2,0,2,2,2,0, //B0
	2,2,1,0,2,2,2,0,2,2,2,0,2,2,2,0, //C0
	2,2,0,0,1,2,2,0,2,2,1,0,1,2,2,0, //D0
	2,2,1,0,2,2,2,0,2,2,2,0,2,2,2,0, //E0
	2,2,0,0,1,2,2,0,2,2,1,0,1,2,2,0} //F0