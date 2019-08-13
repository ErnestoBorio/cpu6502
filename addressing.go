package Cpu6502

func (cpu *Cpu) immediate() byte {
	value := cpu.readMemory[cpu.PC](cpu.PC)
	cpu.PC++
	return value
}

// Returns zero page address from PC's following byte
func (cpu *Cpu) zeroPageAddress() byte {
	address := cpu.readMemory[cpu.PC](cpu.PC)
	cpu.PC++
	return address
}

// TODO: altering the PC before calling ReadMemory can have bad side effects?
func (cpu *Cpu) zeroPage() byte {
	address := cpu.zeroPageAddress()
	cpu.PC++
	return cpu.readMemory[address](word(address))
}

func (cpu *Cpu) zeroPageX() byte {
	address := cpu.zeroPageAddress() + cpu.X
	cpu.PC++
	return cpu.readMemory[address](word(address))
}

func (cpu *Cpu) zeroPageY() byte {
	address := cpu.zeroPageAddress() + cpu.Y
	cpu.PC++
	return cpu.readMemory[address](word(address))
}

// Returns absolute address from PC's following 2 bytes
func (cpu *Cpu) absoluteAddress() word {
	var address word = word( cpu.readMemory[cpu.PC](cpu.PC))
	cpu.PC++
	address |= ( word( cpu.readMemory[cpu.PC](cpu.PC)) << 8 )
	cpu.PC++
	return address
}

func (cpu *Cpu) absolute() byte {
	address := cpu.absoluteAddress()
	return cpu.readMemory[address](address)
}

func (cpu *Cpu) absoluteIndexed(index byte) word {
	lowByte := cpu.readMemory[cpu.PC](cpu.PC)
	indexedLowByte := lowByte + index
	cpu.pageCross = (indexedLowByte < lowByte)
	cpu.PC++
	address := (word(cpu.readMemory[cpu.PC](cpu.PC)) << 8) | word(indexedLowByte)
	cpu.PC++
	return address
}