package Cpu6502

func (cpu *Cpu) LDreg( register *byte, value byte ) { // LDA, LDX, LDY
	*register = value
	cpu.Status.Zero = ( value == 0 )
	cpu.Status.Negative = (( value & signBit ) != 0)
}