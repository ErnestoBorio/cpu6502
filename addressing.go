package cpu6502

// Returns the byte the PC points to
func (cpu *CPU) immediate() byte {
	value := cpu.ReadMemory(cpu.PC)
	cpu.PC++
	return value
}

// Returns zero page address from PC's following byte
func (cpu *CPU) zeroPageAddress() word {
	address := word(cpu.ReadMemory(cpu.PC))
	cpu.PC++
	return address
}

// TODO: altering the PC before calling ReadMemory can have bad side effects?
func (cpu *CPU) zeroPage() byte {
	address := cpu.zeroPageAddress()
	return cpu.ReadMemory(address)
}

func (cpu *CPU) zeroPageIndexedAddress(index byte) word {
	address := word(cpu.ReadMemory(cpu.PC) + index)
	cpu.PC++
	return address
}

func (cpu *CPU) zeroPageIndexed(index byte) byte {
	address := cpu.zeroPageIndexedAddress(index)
	return cpu.ReadMemory(address)
}

// Returns absolute address from PC's following 2 bytes
func (cpu *CPU) absoluteAddress() word {
	var address word = word( cpu.ReadMemory(cpu.PC))
	cpu.PC++
	address |= ( word( cpu.ReadMemory(cpu.PC)) << 8 )
	cpu.PC++
	return address
}

func (cpu *CPU) absolute() byte {
	address := cpu.absoluteAddress()
	return cpu.ReadMemory(address)
}

func (cpu *CPU) absoluteIndexedAddress(index byte) word {
	address := word(cpu.ReadMemory(cpu.PC)) + word(index)
	if address > 0xFF { // if crossed page boundary
		cpu.cycles++
	}
	cpu.PC++
	address += (word(cpu.ReadMemory(cpu.PC)) << 8)
	cpu.PC++
	return address
}

func (cpu *CPU) absoluteIndexed(index byte) byte {
	address := cpu.absoluteIndexedAddress(index)
	return cpu.ReadMemory(address)
}

func (cpu *CPU) indexedIndirectXaddress() word {
	pointer := cpu.ReadMemory(cpu.PC) + cpu.X
	cpu.PC++
	var address word = word( cpu.ReadMemory( word(pointer)))
	pointer++
	address |= (word( cpu.ReadMemory(word(pointer))) << 8)
	return address
}

func (cpu *CPU) indexedIndirectX() byte {
	address := cpu.indexedIndirectXaddress()
	return cpu.ReadMemory(address)
}

func (cpu *CPU) indirectIndexedYaddress() word {
	base := cpu.ReadMemory(cpu.PC)
	cpu.PC++
	address := word(cpu.ReadMemory(word(base))) + word(cpu.Y)
	if address > 0xFF {
		cpu.cycles++
	}
	base++
	return address + (word(cpu.ReadMemory(word(base))) <<8)
}

func (cpu *CPU) indirectIndexedY() byte {
	address := cpu.indirectIndexedYaddress()
	return cpu.ReadMemory(address)
}