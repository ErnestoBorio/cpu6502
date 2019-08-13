package Cpu6502

// LDA, LDX, LDY
func (cpu *Cpu) loadReg( register *byte, value byte ) {
	*register = value
	cpu.Status.Zero = ( value == 0 )
	cpu.Status.Negative = (( value & signBit ) != 0)
}