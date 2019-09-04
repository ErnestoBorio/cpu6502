package cpu6502

// Returns the byte the PC points to
func (cpu *CPU) immediate() byte {
	value := cpu.readMemory[cpu.PC](cpu.PC)
	cpu.PC++
	return value
}

// Returns zero page address from PC's following byte
func (cpu *CPU) zeroPageAddress() word {
	address := word(cpu.readMemory[cpu.PC](cpu.PC))
	cpu.PC++
	return address
}

// TODO: altering the PC before calling ReadMemory can have bad side effects?
func (cpu *CPU) zeroPage() byte {
	address := cpu.zeroPageAddress()
	return cpu.readMemory[address](address)
}

func (cpu *CPU) zeroPageIndexedAddress(index byte) word {
	address := word(cpu.readMemory[cpu.PC](cpu.PC) + index)
	cpu.PC++
	return address
}

func (cpu *CPU) zeroPageIndexed(index byte) byte {
	address := cpu.zeroPageIndexedAddress(index)
	return cpu.readMemory[address](address)
}

// Returns absolute address from PC's following 2 bytes
func (cpu *CPU) absoluteAddress() word {
	var address word = word( cpu.readMemory[cpu.PC](cpu.PC))
	cpu.PC++
	address |= ( word( cpu.readMemory[cpu.PC](cpu.PC)) << 8 )
	cpu.PC++
	return address
}

func (cpu *CPU) absolute() byte {
	address := cpu.absoluteAddress()
	return cpu.readMemory[address](address)
}

func (cpu *CPU) absoluteIndexedAddress(index byte) word {
	address := word(cpu.readMemory[cpu.PC](cpu.PC)) + word(index)
	if address > 0xFF { // if crossed page boundary
		cpu.cycles++
	}
	cpu.PC++
	address += (word(cpu.readMemory[cpu.PC](cpu.PC)) << 8)
	cpu.PC++
	return address
}

func (cpu *CPU) absoluteIndexed(index byte) byte {
	address := cpu.absoluteIndexedAddress(index)
	return cpu.readMemory[address](address)
}

func (cpu *CPU) indexedIndirectXaddress() word {
	pointer := cpu.readMemory[cpu.PC](cpu.PC) + cpu.X
	cpu.PC++
	var address word = word( cpu.readMemory[pointer]( word(pointer)))
	pointer++
	address |= (word( cpu.readMemory[pointer](word(pointer))) << 8)
	return address
}

func (cpu *CPU) indexedIndirectX() byte {
	address := cpu.indexedIndirectXaddress()
	return cpu.readMemory[address](address)
}

func (cpu *CPU) indirectIndexedYaddress() word {
	base := cpu.readMemory[cpu.PC](cpu.PC)
	cpu.PC++
	address := word(cpu.readMemory[base](word(base))) + word(cpu.Y)
	if address > 0xFF {
		cpu.cycles++
	}
	base++
	return address + (word(cpu.readMemory[base](word(base))) <<8)
}

func (cpu *CPU) indirectIndexedY() byte {
	address := cpu.indirectIndexedYaddress()
	return cpu.readMemory[address](address)
}