package Cpu6502

func (cpu *Cpu) immediate() byte {
	value := cpu.readMemory[cpu.PC](cpu.PC)
	cpu.PC++
	return value
}

// TODO: altering the PC before calling ReadMemory can have bad side effects?
func (cpu *Cpu) zeroPage() byte {
	address := cpu.readMemory[cpu.PC](cpu.PC)
	cpu.PC++
	return cpu.readMemory[address](word(address))
}

func (cpu *Cpu) zeroPageX() byte {
	address := cpu.readMemory[cpu.PC](cpu.PC) + cpu.X
	cpu.PC++
	return cpu.readMemory[address](word(address))
}

func (cpu *Cpu) zeroPageY() byte {
	address := cpu.readMemory[cpu.PC](cpu.PC) + cpu.Y
	cpu.PC++
	return cpu.readMemory[address](word(address))
}

func (cpu *Cpu) absolute() byte {
	var address word = word( cpu.readMemory[cpu.PC](cpu.PC))
	cpu.PC++
	address |= ( word( cpu.readMemory[cpu.PC](cpu.PC)) << 8 )
	return cpu.readMemory[address](address)
}