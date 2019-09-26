package cpu6502

// Get a little endian 16 bits value from 2 consecutive memory addresses
func (cpu *CPU) DbgGetUint16(address uint16) uint16 {
	value := uint16( cpu.DbgReadMemory(address))
	value |= uint16( cpu.DbgReadMemory(address+1)) <<8
	return value
}

/* Addressing modes */

func (cpu *CPU) DbgNoAddressing(uint16) uint16 {
	return 0
}

func (cpu *CPU) DbgImmediate(pc uint16) uint16 {
	return pc
}

// Returns zero page address from PC's following byte
func (cpu *CPU) DbgZeroPage(pc uint16) uint16 {
	return uint16(cpu.DbgReadMemory(pc))
}

func (cpu *CPU) DbgZeroPageX(pc uint16) uint16 {
	return uint16(cpu.DbgReadMemory(pc) + cpu.X)
}

func (cpu *CPU) DbgZeroPageY(pc uint16) uint16 {
	return uint16(cpu.DbgReadMemory(pc) + cpu.Y)
}

// Returns absolute address from PC's following 2 bytes
func (cpu *CPU) DbgAbsolute(pc uint16) uint16 {
	return cpu.DbgGetUint16(pc)
}

func (cpu *CPU) DbgAbsoluteX(pc uint16) uint16 {
	address := uint16(cpu.DbgReadMemory(pc)) + uint16(cpu.X)
	if address > 0xFF { // if crossed page boundary
		cpu.cycles++
	}
	address += (uint16(cpu.DbgReadMemory(pc+1)) << 8)
	return address
}

func (cpu *CPU) DbgAbsoluteY(pc uint16) uint16 {
	address := uint16(cpu.DbgReadMemory(pc)) + uint16(cpu.Y)
	if address > 0xFF { // if crossed page boundary
		cpu.cycles++
	}
	address += (uint16(cpu.DbgReadMemory(pc+1)) << 8)
	return address
}

func (cpu *CPU) DbgIndexedIndirectX(pc uint16) uint16 {
	pointer := cpu.DbgReadMemory(pc) + cpu.X
	var address uint16 = uint16( cpu.DbgReadMemory( uint16(pointer)))
	pointer++
	address |= (uint16( cpu.DbgReadMemory(uint16(pointer))) << 8)
	return address
}

func (cpu *CPU) DbgIndirectIndexedY(pc uint16) uint16 {
	base := cpu.DbgReadMemory(pc)
	address := uint16(cpu.DbgReadMemory(uint16(base))) + uint16(cpu.Y)
	if address > 0xFF {
		cpu.cycles++
	}
	base++
	return address + (uint16(cpu.DbgReadMemory(uint16(base))) <<8)
}

func (cpu *CPU) DbgIndirect(pc uint16) uint16 {
	pointer := cpu.DbgGetUint16(pc)
	address := uint16(cpu.DbgReadMemory(pointer)) // low byte
	if (pointer & lowByte) == 0xFF { // address wraps around page
		pointer -= 0x100
	}
	pointer++
	return address | (uint16(cpu.DbgReadMemory(pointer)) <<8)
}
