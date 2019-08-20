package cpu6502

const rightmostBit = 1
const leftmostBit = 1<<7
const bit6 = 1<<6
const signBit = leftmostBit

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