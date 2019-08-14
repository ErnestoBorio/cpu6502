package Cpu6502

// LDA LDX LDY
func (cpu *Cpu) loadReg( register *byte, value byte ) {
	*register = value
	cpu.Status.Zero = ( value == 0 )
	cpu.Status.Negative = (( value & signBit ) != 0)
}

// STA STX STY
func (cpu *Cpu) storeReg( value byte, address word ) {
	cpu.writeMemory[address](address, value)
}

// INX DEX INY DEY
func (cpu *Cpu) incDecReg(register *byte, delta int8) {
	*register += byte(delta)
	cpu.Status.Zero = *register == 0
	cpu.Status.Negative = (*register & signBit) != 0
}