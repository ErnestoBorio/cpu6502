package cpu6502

/* Addressing modes */

func (cpu *CPU) debugNoAddressing() uint16 {
	return 0
}

func (cpu *CPU) debugZeroPageX() uint16 {
	return cpu.zeroPageIndexed(cpu.X)
}

func (cpu *CPU) debugZeroPageY() uint16 {
	return cpu.zeroPageIndexed(cpu.Y)
}

func (cpu *CPU) debugAbsoluteX() uint16 {
	return cpu.absoluteIndexed(cpu.X)
}

func (cpu *CPU) debugAbsoluteY() uint16 {
	return cpu.absoluteIndexed(cpu.Y)
}

func (cpu *CPU) debugIndirect() uint16 {
	pointer := cpu.getUint16(cpu.PC)
	address := uint16(cpu.ReadMemory(pointer)) // low byte
	if (pointer & lowByte) == 0xFF { // address wraps around page
		pointer -= 0x100
	}
	pointer++
	return address | (uint16(cpu.ReadMemory(pointer)) <<8)
}

/* Operations */

func (cpu *CPU) debugUndocumented(uint16) {
	panic("Undocumented opcode")
}

func (cpu *CPU) debugUndocNOP(uint16) {
	panic("Undocumented NOP")
}

func (cpu *CPU) _LDA(address uint16) {
	cpu.loadReg(&cpu.A, address)
}

func (cpu *CPU) _LDX(address uint16) {
	cpu.loadReg(&cpu.X, address)
}

func (cpu *CPU) _LDY(address uint16) {
	cpu.loadReg(&cpu.Y, address)
}

func (cpu *CPU) _STA(address uint16) {
	cpu.storeReg(cpu.A, address)
}

func (cpu *CPU) _STX(address uint16) {
	cpu.storeReg(cpu.X, address)
}

func (cpu *CPU) _STY(address uint16) {
	cpu.storeReg(cpu.Y, address)
}

func (cpu *CPU) _INX(uint16) {
	cpu.incDecReg(&cpu.X, +1)
}

func (cpu *CPU) _INY(uint16) {
	cpu.incDecReg(&cpu.Y, +1)
}

func (cpu *CPU) _DEX(uint16) {
	cpu.incDecReg(&cpu.X, -1)
}

func (cpu *CPU) _DEY(uint16) {
	cpu.incDecReg(&cpu.Y, -1)
}

func (cpu *CPU) _INC(address uint16) {
	cpu.incDec(address, +1)
}

func (cpu *CPU) _DEC(address uint16) {
	cpu.incDec(address, -1)
}

func (cpu *CPU) _CMP(address uint16) {
	cpu.compare(cpu.A, address)
}

func (cpu *CPU) _CPX(address uint16) {
	cpu.compare(cpu.X, address)
}

func (cpu *CPU) _CPY(address uint16) {
	cpu.compare(cpu.Y, address)
}

func (cpu *CPU) _ASLa(uint16) {
	cpu.asla()
}

func (cpu *CPU) _LSRa(uint16) {
	cpu.lsra()
}

func (cpu *CPU) _ROLa(uint16) {
	cpu.rola()
}

func (cpu *CPU) _RORa(uint16) {
	cpu.rora()
}

func (cpu *CPU) _TAX(uint16) {
	cpu.transferRegister(cpu.A, &cpu.X)
}

func (cpu *CPU) _TAY(uint16) {
	cpu.transferRegister(cpu.A, &cpu.Y)
}

func (cpu *CPU) _TXA(uint16) {
	cpu.transferRegister(cpu.X, &cpu.A)
}

func (cpu *CPU) _TYA(uint16) {
	cpu.transferRegister(cpu.Y, &cpu.A)
}

func (cpu *CPU) _TSX(uint16) {
	cpu.transferRegister(cpu.Stack, &cpu.X)
}

func (cpu *CPU) _TXS(uint16) {
	cpu.Stack = cpu.X
}

func (cpu *CPU) _BPL(uint16) {
	cpu.branch(cpu.Status.Negative, false)
}

func (cpu *CPU) _BMI(uint16) {
	cpu.branch(cpu.Status.Negative, true)
}

func (cpu *CPU) _BVC(uint16) {
	cpu.branch(cpu.Status.Overflow, false)
}

func (cpu *CPU) _BVS(uint16) {
	cpu.branch(cpu.Status.Overflow, true)
}

func (cpu *CPU) _BCC(uint16) {
	cpu.branch(cpu.Status.Carry, false)
}

func (cpu *CPU) _BCS(uint16) {
	cpu.branch(cpu.Status.Carry, true)
}

func (cpu *CPU) _BNE(uint16) {
	cpu.branch(cpu.Status.Zero, false)
}

func (cpu *CPU) _BEQ(uint16) {
	cpu.branch(cpu.Status.Zero, true)
}

func (cpu *CPU) _PHP(uint16) {
	cpu.php()
}

func (cpu *CPU) _PLP(uint16) {
	cpu.plp()
}

func (cpu *CPU) _PHA(uint16) {
	cpu.push(cpu.A)
}

func (cpu *CPU) _PLA(uint16) {
	cpu.pla()
}

func (cpu *CPU) _JMP(address uint16) {
	cpu.PC = address
}

func (cpu *CPU) _JSR(uint16) {
	cpu.jsr()
}

func (cpu *CPU) _RTI(uint16) {
	cpu.rti()
}

func (cpu *CPU) _RTS(uint16) {
	cpu.rts()
}

func (cpu *CPU) _BRK(uint16) {
	cpu.PC++
	cpu.irq(true)
}

func (cpu *CPU) _CLC(uint16) {
	cpu.Status.Carry = false
}

func (cpu *CPU) _SEC(uint16) {
	cpu.Status.Carry = true
}

func (cpu *CPU) _CLI(uint16) {
	cpu.Status.NoInterrupt = false
}

func (cpu *CPU) _SEI(uint16) {
	cpu.Status.NoInterrupt = true
}

func (cpu *CPU) _CLD(uint16) {
	cpu.Status.Decimal = false
}

func (cpu *CPU) _SED(uint16) {
	cpu.Status.Decimal = true
}

func (cpu *CPU) _CLV(uint16) {
	cpu.Status.Overflow = false
}

func (cpu *CPU) _NOP(uint16) {
}
