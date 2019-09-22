package cpu6502

// Returns the address following the opcode (PC+1)
func (cpu *CPU) immediate() uint16 {
	address := cpu.PC
	cpu.PC++
	return address
}

// Returns zero page address from PC's following byte
func (cpu *CPU) zeroPage() uint16 {
	address := uint16(cpu.ReadMemory(cpu.PC))
	cpu.PC++
	return address
}

func (cpu *CPU) zeroPageIndexed(index byte) uint16 {
	address := uint16(cpu.ReadMemory(cpu.PC) + index)
	cpu.PC++
	return address
}

// Returns absolute address from PC's following 2 bytes
func (cpu *CPU) absolute() uint16 {
	address := cpu.getUint16(cpu.PC)
	cpu.PC += 2
	return address
}

func (cpu *CPU) absoluteIndexed(index byte) uint16 {
	address := uint16(cpu.ReadMemory(cpu.PC)) + uint16(index)
	if address > 0xFF { // if crossed page boundary
		cpu.cycles++
	}
	cpu.PC++
	address += (uint16(cpu.ReadMemory(cpu.PC)) << 8)
	cpu.PC++
	return address
}

func (cpu *CPU) indexedIndirectX() uint16 {
	pointer := cpu.ReadMemory(cpu.PC) + cpu.X
	cpu.PC++
	var address uint16 = uint16( cpu.ReadMemory( uint16(pointer)))
	pointer++
	address |= (uint16( cpu.ReadMemory(uint16(pointer))) << 8)
	return address
}

func (cpu *CPU) indirectIndexedY() uint16 {
	base := cpu.ReadMemory(cpu.PC)
	cpu.PC++
	address := uint16(cpu.ReadMemory(uint16(base))) + uint16(cpu.Y)
	if address > 0xFF {
		cpu.cycles++
	}
	base++
	return address + (uint16(cpu.ReadMemory(uint16(base))) <<8)
}