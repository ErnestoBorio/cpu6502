package Cpu6502

func (cpu *Cpu) step() byte {
	cpu.pageCross = false
	var opcode = cpu.ReadMemory[cpu.PC](cpu.PC)

	// this can be incremented by special cases like branch page crossing
	cpu.cycles = opcode_cycles[opcode]

	// The 6502 reads the next byte in advance to gain time, this could have side effects, so it's not trivial
	var operand = cpu.ReadMemory[cpu.PC+1](cpu.PC+1)

	
	return cpu.cycles
}

func (cpu *Cpu) implied() {
	cpu.PC += 1
}

func (cpu *Cpu) immediate() {
	cpu.PC += 2
}
// TODO: altering the PC before calling ReadMemory can have bad side effects?
func (cpu *Cpu) zeroPage(address byte) byte {
	cpu.PC += 2
	return cpu.ReadMemory[address](word(address))
}

func (cpu *Cpu) zeroPageX(address byte) byte {
	cpu.PC += 2
	effectiveAddress := word(address + cpu.X)
	return cpu.ReadMemory[effectiveAddress](effectiveAddress)
}

func (cpu *Cpu) zeroPageY(address byte) byte {
	cpu.PC += 2
	effectiveAddress := word(address + cpu.Y)
	return cpu.ReadMemory[effectiveAddress](effectiveAddress)
}

