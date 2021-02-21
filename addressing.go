package cpu6502

// noAddressing is a dummy function for opcodes with no addressing, E.G. INX, NOP
func (*CPU) noAddressing() uint16 {
	return 0
}

func (cpu *CPU) immediate() uint16 {
	address := cpu.PC
	cpu.PC++
	return address
}

func (cpu *CPU) zeroPage() uint16 {
	address := uint16(cpu.ReadMemory(cpu.PC))
	cpu.PC++
	return address
}

func (cpu *CPU) zeroPageX() uint16 {
	address := uint16(cpu.ReadMemory(cpu.PC) + cpu.X)
	cpu.PC++
	return address
}

func (cpu *CPU) zeroPageY() uint16 {
	address := uint16(cpu.ReadMemory(cpu.PC) + cpu.Y)
	cpu.PC++
	return address
}

func (cpu *CPU) absolute() uint16 {
	address := cpu.getUint16(cpu.PC)
	cpu.PC += 2
	return address
}

func (cpu *CPU) absoluteIndexed(index byte) uint16 {
	address := uint16(cpu.ReadMemory(cpu.PC)) + uint16(index)
	if address > 0xFF { // if crossed page boundary
		cpu.tmpCycles++
	}
	cpu.PC++
	address += (uint16(cpu.ReadMemory(cpu.PC)) << 8)
	cpu.PC++
	return address
}

func (cpu *CPU) absoluteIndexedX() uint16 {
	return cpu.absoluteIndexed(cpu.X)
}

func (cpu *CPU) absoluteIndexedY() uint16 {
	return cpu.absoluteIndexed(cpu.Y)
}

func (cpu *CPU) indexedIndirectX() uint16 {
	pointer := cpu.ReadMemory(cpu.PC) + cpu.X
	cpu.PC++
	var address uint16 = uint16(cpu.ReadMemory(uint16(pointer)))
	pointer++
	address |= (uint16(cpu.ReadMemory(uint16(pointer))) << 8)
	return address
}

func (cpu *CPU) indirectIndexedY() uint16 {
	base := cpu.ReadMemory(cpu.PC)
	cpu.PC++
	address := uint16(cpu.ReadMemory(uint16(base))) + uint16(cpu.Y)
	if address > 0xFF {
		cpu.tmpCycles++
	}
	base++
	return address + (uint16(cpu.ReadMemory(uint16(base))) << 8)
}

func (cpu *CPU) absoluteX() uint16 {
	address := uint16(cpu.ReadMemory(cpu.PC)) + uint16(cpu.X)
	cpu.PC++
	if address > 0xFF { // if crossed page boundary
		cpu.tmpCycles++
	}
	address += (uint16(cpu.ReadMemory(cpu.PC+1)) << 8)
	return address
}

func (cpu *CPU) absoluteY() uint16 {
	address := uint16(cpu.ReadMemory(cpu.PC)) + uint16(cpu.Y)
	cpu.PC++
	if address > 0xFF { // if crossed page boundary
		cpu.tmpCycles++
	}
	address += (uint16(cpu.ReadMemory(cpu.PC+1)) << 8)
	return address
}

func (cpu *CPU) indirect() uint16 {
	pointer := cpu.getUint16(cpu.PC)
	cpu.PC++
	address := uint16(cpu.ReadMemory(pointer)) // low byte
	if (pointer & 0xFF) == 0xFF { // address wraps around page
		pointer -= 0x100
	}
	pointer++
	return address | (uint16(cpu.ReadMemory(pointer)) << 8)
}