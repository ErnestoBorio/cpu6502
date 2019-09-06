package cpu6502

// Returns the address following the opcode (PC+1)
func (cpu *CPU) immediate() word {
	address := cpu.PC
	cpu.PC++
	return address
}

// Returns zero page address from PC's following byte
func (cpu *CPU) zeroPage() word {
	address := word(cpu.ReadMemory(cpu.PC))
	cpu.PC++
	return address
}

func (cpu *CPU) zeroPageIndexed(index byte) word {
	address := word(cpu.ReadMemory(cpu.PC) + index)
	cpu.PC++
	return address
}

// Returns absolute address from PC's following 2 bytes
func (cpu *CPU) absolute() word {
	var address word = word( cpu.ReadMemory(cpu.PC))
	cpu.PC++
	address |= ( word( cpu.ReadMemory(cpu.PC)) << 8 )
	cpu.PC++
	return address
}

func (cpu *CPU) absoluteIndexed(index byte) word {
	address := word(cpu.ReadMemory(cpu.PC)) + word(index)
	if address > 0xFF { // if crossed page boundary
		cpu.cycles++
	}
	cpu.PC++
	address += (word(cpu.ReadMemory(cpu.PC)) << 8)
	cpu.PC++
	return address
}

func (cpu *CPU) indexedIndirectX() word {
	pointer := cpu.ReadMemory(cpu.PC) + cpu.X
	cpu.PC++
	var address word = word( cpu.ReadMemory( word(pointer)))
	pointer++
	address |= (word( cpu.ReadMemory(word(pointer))) << 8)
	return address
}

func (cpu *CPU) indirectIndexedY() word {
	base := cpu.ReadMemory(cpu.PC)
	cpu.PC++
	address := word(cpu.ReadMemory(word(base))) + word(cpu.Y)
	if address > 0xFF {
		cpu.cycles++
	}
	base++
	return address + (word(cpu.ReadMemory(word(base))) <<8)
}